package operator

import (
	"context"
	"fmt"
	"net/rpc"
	"time"

	aggRpc "github.com/HemeraProtocol/avs/legacy/aggregator/rpc"
	"github.com/HemeraProtocol/avs/legacy/api/grpc/aggregator"
	"github.com/HemeraProtocol/avs/legacy/core/alert"
	"github.com/HemeraProtocol/avs/legacy/core/config"
	"github.com/HemeraProtocol/avs/legacy/core/message"
	"github.com/HemeraProtocol/avs/legacy/metrics"
	"github.com/Layr-Labs/eigensdk-go/logging"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	gethrpc "github.com/ethereum/go-ethereum/rpc"
)

type AggregatorJsonRpcClient struct {
	rpcClient                   *rpc.Client
	metrics                     metrics.Metrics
	logger                      logging.Logger
	config                      config.NodeConfig
	operatorId                  sdktypes.OperatorId
	operatorAddr                common.Address
	OperatorStateRetrieverAddr  common.Address
	RegistryCoordinatorAddr     common.Address
	jsonRPCAggregatorIpPortAddr string
	timeout                     time.Duration
}

func NewAggregatorJsonRpcClient(config config.NodeConfig, operatorId sdktypes.OperatorId, operatorAddr common.Address, logger logging.Logger, metrics metrics.Metrics) (*AggregatorJsonRpcClient, error) {
	return &AggregatorJsonRpcClient{
		// set to nil so that we can create an rpc client even if the aggregator is not running
		rpcClient:                   nil,
		metrics:                     metrics,
		logger:                      logger,
		config:                      config,
		operatorId:                  operatorId,
		operatorAddr:                operatorAddr,
		OperatorStateRetrieverAddr:  common.HexToAddress(config.OperatorStateRetrieverAddress),
		RegistryCoordinatorAddr:     common.HexToAddress(config.AVSRegistryCoordinatorAddress),
		jsonRPCAggregatorIpPortAddr: config.AggregatorJSONRPCServerIpPortAddr,
		timeout:                     1 * time.Second,
	}, nil
}

// CreateAlertTaskToAggregator create a new alert task, if had existing, just return current alert task.
func (c *AggregatorJsonRpcClient) InitOperatorToAggregator() error {
	client, err := gethrpc.DialContext(context.Background(), c.jsonRPCAggregatorIpPortAddr)
	if err != nil {
		return fmt.Errorf("dial initOperatorToAggregator connection failed: %v", err.Error())
	}

	var res aggRpc.InitOperatorResponse

	err = client.CallContext(
		context.Background(), &res, "aggregator_initOperator",
		c.config.Layer1ChainId,
		c.config.Layer2ChainId,
		hexutil.Bytes(c.operatorId[:]),
		c.operatorAddr.Hex(),
		c.config.OperatorStateRetrieverAddress,
		c.config.AVSRegistryCoordinatorAddress,
	)
	if err != nil {
		return fmt.Errorf("call initOperatorToAggregator failed: %v", err.Error())
	}

	if !res.Ok {
		return fmt.Errorf("initOperatorToAggregator failed: %v", res.Reason)
	}

	return nil

}

// CreateAlertTaskToAggregator create a new alert task, if had existing, just return current alert task.
func (c *AggregatorJsonRpcClient) CreateAlertTaskToAggregator(alertHash [32]byte) (*message.AlertTaskInfo, error) {
	client, err := gethrpc.DialContext(context.Background(), c.jsonRPCAggregatorIpPortAddr)
	if err != nil {
		return nil, fmt.Errorf("dial CreateAlertTask connection failed: %v", err.Error())
	}

	var res aggRpc.AlertTaskInfo
	err = client.CallContext(
		context.Background(), &res, "aggregator_createTask",
		hexutil.Bytes(alertHash[:]),
	)

	if err != nil {
		return nil, fmt.Errorf("call CreateAlertTask failed: %v", err.Error())
	}

	info, err := message.NewAlertTaskInfo(&aggregator.AlertTaskInfo{
		AlertHash:                  res.AlertHash,
		QuorumNumbers:              res.QuorumNumbers,
		QuorumThresholdPercentages: res.QuorumThresholdPercentages,
		TaskIndex:                  res.TaskIndex,
		ReferenceBlockNumber:       res.ReferenceBlockNumber,
	})
	if err != nil {
		return nil, fmt.Errorf("call CreateAlertTask failed by decode alert info: %v", err.Error())
	}

	return info, nil
}

// SendSignedTaskResponseToAggregator sends a signed task response to the aggregator.
// it is meant to be ran inside a go thread, so doesn't return anything.
// this is because sending the signed task response to the aggregator is time sensitive,
// so there is no point in retrying if it fails for a few times.
// Currently hardcoded to retry sending the signed task response 5 times, waiting 2 seconds in between each attempt.
func (c *AggregatorJsonRpcClient) SendSignedTaskResponseToAggregator(signedTaskResponse *message.SignedTaskRespRequest, resChan chan alert.AlertResponse) {
	client, err := gethrpc.DialContext(context.Background(), c.jsonRPCAggregatorIpPortAddr)
	if err != nil {
		resChan <- alert.AlertResponse{
			Err: err,
			Msg: "dial initOperatorToAggregator connection failed",
		}
		return
	}

	alertData := signedTaskResponse.Alert.ToPbType()
	alertDataReq := aggRpc.AlertTaskInfo{
		AlertHash:                  alertData.AlertHash,
		QuorumNumbers:              alertData.QuorumNumbers,
		QuorumThresholdPercentages: alertData.QuorumThresholdPercentages,
		TaskIndex:                  alertData.TaskIndex,
		ReferenceBlockNumber:       alertData.ReferenceBlockNumber,
	}
	qperatorRequestSignature := signedTaskResponse.BlsSignature.Serialize()

	c.logger.Info("CreateAlertTask to aggregator", "alert", fmt.Sprintf("%#v", alertDataReq))

	var resp aggRpc.SignedTaskRespResponse
	err = client.CallContext(
		context.Background(), &resp, "aggregator_processSignedTaskResponse",
		alertDataReq, hexutil.Bytes(qperatorRequestSignature), hexutil.Bytes(signedTaskResponse.OperatorId[:]),
	)
	if err != nil {
		resChan <- alert.AlertResponse{
			Err: err,
			Msg: "call CreateAlertTask failed",
		}
		return
	}

	c.logger.Info("Signed task response header accepted by aggregator.", "response", fmt.Sprintf("%#v", resp))

	res := alert.AlertResponse{
		Code:      0,
		TaskIndex: signedTaskResponse.Alert.TaskIndex,
	}
	copy(res.TxHash[:], resp.TxHash[:32])

	c.logger.Info("Signed task resp", "response", res)
	c.metrics.IncNumTasksAcceptedByAggregator()

	resChan <- res
}
