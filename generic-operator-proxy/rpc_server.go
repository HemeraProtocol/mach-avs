package genericproxy

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"github.com/sourcegraph/jsonrpc2"

	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"

	"github.com/alt-research/avs-generic-aggregator/core/config"
	"github.com/alt-research/avs-generic-aggregator/core/types"
	proxyUtils "github.com/alt-research/avs-generic-aggregator/proxy/utils"
	"github.com/alt-research/avs/core/alert"
	"github.com/alt-research/avs/operator"
)

// Hash32HeaderParam is the 1st parameter for the Mach AVS contract 's confirmAlert
type AlertHeaderParam struct {
	MessageHash                [32]byte `abi:"messageHash"`
	QuorumNumbers              []byte   `abi:"quorumNumbers"`
	QuorumThresholdPercentages []byte   `abi:"quorumThresholdPercentages"`
	ReferenceBlockNumber       uint32   `abi:"referenceBlockNumber"`
}

type CreateSigTaskForByte32Hash struct {
	Hash types.Bytes32 `json:"hash"`
}

type ProxyHashRpcServer struct {
	baseServers    map[string]*proxyUtils.ProxyRpcServerBase
	avsCfgs        map[string]config.GenericAVSConfig
	defaultAVSName string
	logger         logging.Logger
	// We use this rpc server as the handler for jsonrpc
	// to make sure same as legacy operator 's api
	rpcServer operator.RpcServer
	// receive new tasks in this chan (typically from mach service)
	newTaskCreatedChan chan alert.AlertRequest
}

func NewAlertProxyRpcServer(
	logger logging.Logger,
	ethClient eth.Client,
	avsCfgs []config.GenericAVSConfig,
	defaultAVSName string,
	method string,
	genericOperatorAddr string,
	jsonrpcCfg config.JsonRpcServerConfig,
) *ProxyHashRpcServer {
	bases := make(map[string]*proxyUtils.ProxyRpcServerBase, len(avsCfgs))
	avsCfgsMap := make(map[string]config.GenericAVSConfig, len(avsCfgs))
	for _, avsCfg := range avsCfgs {
		logger.Infof("load service for %s", avsCfg.AVSName)
		base := proxyUtils.NewProxyRpcServerBase(
			logger,
			ethClient,
			avsCfg,
			method,
			genericOperatorAddr,
			jsonrpcCfg,
		)
		bases[avsCfg.AVSName] = base
		avsCfgsMap[avsCfg.AVSName] = avsCfg
	}

	newTaskCreatedChan := make(chan alert.AlertRequest, 32)
	rpcServer := operator.NewRpcServer(logger, jsonrpcCfg.Addr, newTaskCreatedChan)

	server := &ProxyHashRpcServer{
		baseServers:        bases,
		logger:             logger,
		defaultAVSName:     defaultAVSName,
		newTaskCreatedChan: newTaskCreatedChan,
		rpcServer:          rpcServer,
		avsCfgs:            avsCfgsMap,
	}

	// not need do this because we not need use this server impl
	// server.SetHandler(server)

	return server
}

func (s *ProxyHashRpcServer) HttpRPCHandler(w http.ResponseWriter, r *http.Request) {
	rpcRequest := jsonrpc2.Request{}
	err := json.NewDecoder(r.Body).Decode(&rpcRequest)
	if err != nil {
		operator.WriteErrorJSON(s.logger, w, rpcRequest.ID, http.StatusBadRequest, 1, err)
		return
	}

	if rpcRequest.Params == nil {
		err := errors.New("failed to unmarshal request.Params for mevBundle from mev-builder, error: EOF")
		operator.WriteErrorJSON(s.logger, w, rpcRequest.ID, http.StatusBadRequest, 1, err)
		return
	}

	if rpcRequest.Method == "operator_getConfig" {
		var req []string
		if err = json.Unmarshal(*rpcRequest.Params, &req); err != nil {
			s.logger.Error("the unmarshal", "err", err)
			operator.WriteErrorJSON(s.logger, w, rpcRequest.ID, http.StatusBadRequest, 3, fmt.Errorf("failed to unmarshal req bundle params: %s", err.Error()))
			return
		}

		if len(req) != 1 {
			operator.WriteErrorJSON(s.logger, w, rpcRequest.ID, http.StatusBadRequest, 3, fmt.Errorf("failed to unmarshal req bundle params"))
			return
		}

		baseService, ok := s.baseServers[req[0]]
		if !ok || baseService == nil {
			operator.WriteErrorJSON(s.logger, w, rpcRequest.ID, http.StatusBadRequest, 3, fmt.Errorf("failed to found the avs config"))
			return
		}

		avsCfg, ok := s.avsCfgs[req[0]]
		if !ok {
			operator.WriteErrorJSON(s.logger, w, rpcRequest.ID, http.StatusBadRequest, 3, fmt.Errorf("failed to found the avs config"))
			return
		}

		cfg := avsCfg.OperatorConfigs
		if cfg == "" {
			cfg = "{}"
		}

		operator.WriteJSON(s.logger, w, rpcRequest.ID, http.StatusOK, json.RawMessage(cfg))
	} else {
		s.rpcServer.HttpRPCHandlerRequest(w, rpcRequest)
	}

}

func (s *ProxyHashRpcServer) Start(ctx context.Context) error {
	s.logger.Info("start rpc server for got alert")

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.HttpRPCHandler)
	s.rpcServer.SetHandler(mux)

	if err := s.rpcServer.StartServer(ctx); err != nil {
		s.logger.Error("Error start Rpc server", "err", err)
		return err
	}

	defer func() {
		err := s.rpcServer.Stop()
		if err != nil {
			s.logger.Error("Stop Rpc server failed", "err", err)
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return nil
		case newTaskCreatedLog := <-s.newTaskCreatedChan:
			s.logger.Info("newTaskCreatedLog", "new", newTaskCreatedLog.Alert)
			resp, err := s.createSigTaskH32(ctx, newTaskCreatedLog.Alert.GetAVSName(), newTaskCreatedLog.Alert.MessageHash())
			if err != nil {
				newTaskCreatedLog.ResChan <- alert.AlertResponse{
					Err: err,
					Msg: "createSigTaskH32 failed",
				}
				continue
			}

			newTaskCreatedLog.ResChan <- alert.AlertResponse{
				TxHash:    resp.TxHash,
				TaskIndex: uint32(resp.TaskIndex),
			}
		}
	}
}

func (s *ProxyHashRpcServer) createSigTaskH32(ctx context.Context, avsName string, hash [32]byte) (proxyUtils.CreateSigTaskResp, error) {
	avsNameToProxy := avsName
	if avsNameToProxy == "" {
		avsNameToProxy = s.defaultAVSName
	}
	baseServer, ok := s.baseServers[avsNameToProxy]
	if !ok || baseServer == nil {
		return proxyUtils.CreateSigTaskResp{}, errors.Errorf("not found avs name %v", avsNameToProxy)
	}

	res, err := baseServer.CreateGenericSigTaskWithSigHash(
		ctx,
		func(
			referenceBlockNumber uint64,
			quorumNumbers sdktypes.QuorumNums,
			quorumThresholdPercentages sdktypes.QuorumThresholdPercentages,
		) ([]interface{}, [32]byte) {
			sigHash, err := CalcSighHash(hash, uint32(referenceBlockNumber))
			if err != nil {
				panic(err)
			}

			return []interface{}{
				AlertHeaderParam{
					MessageHash:                hash,
					QuorumNumbers:              quorumNumbers.UnderlyingType(),
					QuorumThresholdPercentages: quorumThresholdPercentages.UnderlyingType(),
					ReferenceBlockNumber:       uint32(referenceBlockNumber),
				}}, sigHash
		})
	if err != nil {
		return proxyUtils.CreateSigTaskResp{}, errors.Wrap(err, "create sig task failed")
	}

	response := proxyUtils.CreateSigTaskResp{
		TaskIndex: uint64(res.TaskIndex),
		TxHash:    res.TxHash,
		SigHash:   res.SigHash,
	}

	return response, nil
}

var (
	sighHashAbiParams, _ = abi.NewType("tuple", "Hash32SighHashParam", []abi.ArgumentMarshaling{
		{Name: "messageHash", Type: "bytes32"},
		{Name: "referenceBlockNumber", Type: "uint32"},
	})

	sighHashAbiArgs = abi.Arguments{
		{Type: sighHashAbiParams, Name: "one"},
	}
)

func CalcSighHash(messageHash [32]byte, referenceBlockNumber uint32) ([32]byte, error) {
	record := struct {
		MessageHash          [32]byte `abi:"messageHash"`
		ReferenceBlockNumber uint32   `abi:"referenceBlockNumber"`
	}{
		MessageHash:          messageHash,
		ReferenceBlockNumber: referenceBlockNumber,
	}

	packed, err := sighHashAbiArgs.Pack(&record)
	if err != nil {
		return [32]byte{}, err
	}

	return crypto.Keccak256Hash(packed), nil

}
