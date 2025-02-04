package aggregator

import (
	"context"
	"encoding/hex"
	"fmt"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	sdkclients "github.com/Layr-Labs/eigensdk-go/chainio/clients"
	"github.com/Layr-Labs/eigensdk-go/chainio/clients/eth"
	"github.com/Layr-Labs/eigensdk-go/logging"
	"github.com/Layr-Labs/eigensdk-go/services/avsregistry"
	blsagg "github.com/Layr-Labs/eigensdk-go/services/bls_aggregation"
	"github.com/Layr-Labs/eigensdk-go/services/operatorsinfo"

	"github.com/HemeraProtocol/avs/legacy/aggregator/models"
	"github.com/HemeraProtocol/avs/legacy/aggregator/rpc"
	"github.com/HemeraProtocol/avs/legacy/aggregator/types"
	"github.com/HemeraProtocol/avs/legacy/core/chainio"
	"github.com/HemeraProtocol/avs/legacy/core/config"
	"github.com/HemeraProtocol/avs/legacy/core/message"
	sdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/ethereum/go-ethereum/common"
)

type AggregatorService struct {
	logger logging.Logger
	cfg    *config.Config

	avsReader chainio.AvsReaderer
	ethClient eth.Client

	blsAggregationService blsagg.BlsAggregationService
	tasks                 map[types.TaskIndex]*message.AlertTaskInfo
	tasksMu               sync.RWMutex
	finishedTasks         map[[32]byte]*FinishedTaskStatus
	finishedTasksMu       sync.RWMutex
	nextTaskIndex         types.TaskIndex
	nextTaskIndexMu       sync.RWMutex
	operatorStatus        map[common.Address]*OperatorStatus
	operatorStatusMu      sync.RWMutex
	model                 *models.Model
}

// NewAggregator creates a new Aggregator with the provided config.
func NewAggregatorService(c *config.Config) (*AggregatorService, error) {
	avsReader, err := chainio.BuildAvsReaderFromConfig(c)
	if err != nil {
		c.Logger.Error("Cannot create avsReader", "err", err)
		return nil, err
	}

	chainioConfig := sdkclients.BuildAllConfig{
		EthHttpUrl:                 c.EthHttpRpcUrl,
		EthWsUrl:                   c.EthWsRpcUrl,
		RegistryCoordinatorAddr:    c.RegistryCoordinatorAddr.String(),
		OperatorStateRetrieverAddr: c.OperatorStateRetrieverAddr.String(),
		AvsName:                    avsName,
		PromMetricsIpPortAddress:   ":9090",
	}
	clients, err := sdkclients.BuildAll(chainioConfig, c.PrivateKey, c.Logger)
	if err != nil {
		c.Logger.Errorf("Cannot create sdk clients", "err", err)
		return nil, err
	}

	operatorsinfoService := operatorsinfo.NewOperatorsInfoServiceInMemory(context.Background(), clients.AvsRegistryChainSubscriber, clients.AvsRegistryChainReader, c.Logger)
	avsRegistryService := avsregistry.NewAvsRegistryServiceChainCaller(avsReader, operatorsinfoService, c.Logger)
	blsAggregationService := blsagg.NewBlsAggregatorService(avsRegistryService, c.Logger)

	db, err := gorm.Open(postgres.Open(c.DBConfig), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &AggregatorService{
		logger:                c.Logger,
		avsReader:             avsReader,
		ethClient:             clients.EthHttpClient,
		blsAggregationService: blsAggregationService,
		tasks:                 make(map[types.TaskIndex]*message.AlertTaskInfo),
		finishedTasks:         make(map[[32]byte]*FinishedTaskStatus),
		operatorStatus:        make(map[common.Address]*OperatorStatus),
		cfg:                   c,
		model:                 models.NewModel(db),
	}, nil
}

var _ rpc.AggregatorRpcHandler = (*AggregatorService)(nil)

func (agg *AggregatorService) GetTaskByAlertHash(alertHash [32]byte) *message.AlertTaskInfo {
	agg.tasksMu.RLock()
	defer agg.tasksMu.RUnlock()

	for _, task := range agg.tasks {
		if task.AlertHash == alertHash {
			return task
		}
	}

	return nil
}

func (agg *AggregatorService) GetTaskByIndex(taskIndex types.TaskIndex) *message.AlertTaskInfo {
	agg.tasksMu.RLock()
	defer agg.tasksMu.RUnlock()

	res := agg.tasks[taskIndex]

	return res
}

func (agg *AggregatorService) newIndex() types.TaskIndex {
	agg.nextTaskIndexMu.Lock()
	defer agg.nextTaskIndexMu.Unlock()

	res := agg.nextTaskIndex
	agg.nextTaskIndex += 1

	return res
}

func (agg *AggregatorService) GetFinishedTaskByAlertHash(alertHash [32]byte) *FinishedTaskStatus {
	agg.finishedTasksMu.RLock()
	defer agg.finishedTasksMu.RUnlock()

	return agg.finishedTasks[alertHash]
}

func (agg *AggregatorService) SetFinishedTask(alertHash [32]byte, finished *FinishedTaskStatus) {
	agg.finishedTasksMu.Lock()
	defer agg.finishedTasksMu.Unlock()

	agg.finishedTasks[alertHash] = finished
}

// rpc endpoint which is called by operator
// will init operator, just for keep config valid
func (agg *AggregatorService) InitOperator(req *message.InitOperatorRequest) (*message.InitOperatorResponse, error) {
	agg.logger.Infof("Received InitOperator: %#v", req)

	reply := &message.InitOperatorResponse{
		Ok: false,
	}

	if agg.cfg.OperatorStateRetrieverAddr != req.OperatorStateRetrieverAddr {
		reply.Res = fmt.Sprintf("OperatorStateRetrieverAddr invaild, expect %s", agg.cfg.OperatorStateRetrieverAddr.Hex())
		return reply, nil
	}

	if agg.cfg.RegistryCoordinatorAddr != req.RegistryCoordinatorAddr {
		reply.Res = fmt.Sprintf("RegistryCoordinatorAddr invaild, expect %s", agg.cfg.RegistryCoordinatorAddr.Hex())
		return reply, nil
	}

	if agg.cfg.Layer1ChainId != req.Layer1ChainId {
		reply.Res = fmt.Sprintf("Layer1ChainId invaild, expect %d", agg.cfg.Layer1ChainId)
		return reply, nil
	}

	if agg.cfg.Layer2ChainId != req.ChainId {
		reply.Res = fmt.Sprintf("Layer2ChainId invaild, expect %d", agg.cfg.Layer2ChainId)
		return reply, nil
	}

	agg.operatorStatusMu.Lock()
	defer agg.operatorStatusMu.Unlock()

	agg.operatorStatus[req.OperatorAddress] = &OperatorStatus{
		LastTime:   time.Now().Unix(),
		OperatorId: req.OperatorId,
	}

	reply.Ok = true

	agg.logger.Infof("new operator status: %s", req.OperatorAddress.Hex())

	return reply, nil
}

// rpc endpoint which is called by operator
// will try to init the task, if currently had a same task for the alert,
// it will return the existing task.
func (agg *AggregatorService) CreateTask(req *message.CreateTaskRequest) (*message.CreateTaskResponse, error) {
	agg.logger.Info("Received CreateTask", "alertHash", req.AlertHash)

	finished := agg.GetFinishedTaskByAlertHash(req.AlertHash)
	if finished != nil {
		return nil, fmt.Errorf("the task 0x%x already finished: 0x%x", req.AlertHash, finished.TxHash)
	}

	task := agg.GetTaskByAlertHash(req.AlertHash)
	if task == nil {
		agg.logger.Info("create new task", "alert", req.AlertHash)
		taskIndex := agg.newIndex()

		var err error
		task, err = agg.sendNewTask(req.AlertHash, taskIndex)

		if err != nil {
			agg.logger.Error("send new task failed", "err", err)
			return nil, err
		}
	} else {
		agg.logger.Info("the task had created", "task", task)
	}

	modelTask := &models.AggregatorTask{
		AlertHash:            task.AlertHash[:],
		QuorumNumbers:        task.QuorumNumbers.UnderlyingType(),
		TaskIndex:            task.TaskIndex,
		ReferenceBlockNumber: task.ReferenceBlockNumber,
	}
	if err := agg.model.CreateTask(modelTask); err != nil {
		agg.logger.Error("create task failed", "err", err, "task", modelTask)
	}

	return &message.CreateTaskResponse{Info: *task}, nil
}

// rpc endpoint which is called by operator
// reply doesn't need to be checked. If there are no errors, the task response is accepted
// rpc framework forces a reply type to exist, so we put bool as a placeholder
func (agg *AggregatorService) ProcessSignedTaskResponse(signedTaskResponse *message.SignedTaskRespRequest) (*message.SignedTaskRespResponse, error) {
	agg.logger.Info(
		"Received signed task response",
		"alert", signedTaskResponse.Alert,
		"operatorId", hex.EncodeToString(signedTaskResponse.OperatorId[:]),
	)

	taskIndex := signedTaskResponse.Alert.TaskIndex
	taskResponseDigest, err := signedTaskResponse.Alert.SignHash()
	if err != nil {
		return nil, err
	}

	if task := agg.GetTaskByIndex(taskIndex); task == nil {
		agg.logger.Error("ProcessNewSignature error by no task exist", "taskIndex", taskIndex)
		return nil, fmt.Errorf("task not found")
	}

	modelTaskSignature := &models.AggregatorTaskSignature{
		AlertHash:  signedTaskResponse.Alert.AlertHash[:],
		OperatorId: signedTaskResponse.OperatorId[:],
		SignResult: true,
	}

	agg.logger.Infof("ProcessNewSignature: %#v", signedTaskResponse.Alert.TaskIndex)
	err = agg.blsAggregationService.ProcessNewSignature(
		context.Background(), taskIndex, taskResponseDigest,
		&signedTaskResponse.BlsSignature, signedTaskResponse.OperatorId,
	)
	if err != nil {
		agg.logger.Error("ProcessNewSignature error", "err", err)
		modelTaskSignature.SignResult = false
	}

	if err := agg.model.CreateTaskSignature(modelTaskSignature); err != nil {
		agg.logger.Error("create task signature failed", "err", err, "taskSignature", modelTaskSignature)
	}

	return &message.SignedTaskRespResponse{}, err
}

// GetResponseChannel returns the single channel that meant to be used as the response channel
// Any task that is completed (see the completion criterion in the comment above InitializeNewTask)
// will be sent on this channel along with all the necessary information to call BLSSignatureChecker onchain
func (agg *AggregatorService) GetResponseChannel() <-chan blsagg.BlsAggregationServiceResponse {
	return agg.blsAggregationService.GetResponseChannel()
}

// sendNewTask sends a new task to the task manager contract, and updates the Task dict struct
// with the information of operators opted into quorum 0 at the block of task creation.
func (agg *AggregatorService) sendNewTask(alertHash message.Bytes32, taskIndex types.TaskIndex) (*message.AlertTaskInfo, error) {
	agg.logger.Info("Aggregator sending new task", "alert", alertHash, "task", taskIndex)

	var err error

	var referenceBlockNumber uint64
	if referenceBlockNumber, err = agg.ethClient.BlockNumber(context.Background()); err != nil {
		return nil, err
	}

	// the reference block number must < the current block number.
	referenceBlockNumber -= 1

	agg.logger.Info("get from layer1", "referenceBlockNumber", referenceBlockNumber)

	quorumNumbers, err := agg.avsReader.GetQuorumsByBlockNumber(context.Background(), uint32(referenceBlockNumber))
	if err != nil {
		agg.logger.Error("GetQuorumCountByBlockNumber failed", "err", err)
		return nil, err
	}
	agg.logger.Info("get quorumNumbers from layer1", "quorumNumbers", fmt.Sprintf("%v", quorumNumbers))

	if len(quorumNumbers) < len(agg.cfg.QuorumNums) {
		agg.logger.Error("the cfg quorum numbers is larger to the layer1, it will commit failed")
		return nil, fmt.Errorf("the quorum numbers is larger to the layer1 %v, expected %v", agg.cfg.QuorumNums, quorumNumbers)
	}

	// just use config value
	quorumNumbers = agg.cfg.QuorumNums

	// hardcoded in the contract
	quorumThresholdPercentages := sdktypes.QuorumThresholdPercentages{66}

	agg.logger.Info(
		"quorum datas",
		"numbers", fmt.Sprintf("%v", quorumNumbers.UnderlyingType()),
		"thresholdPercentages", fmt.Sprintf("%v", quorumThresholdPercentages.UnderlyingType()),
	)

	newAlertTask := &message.AlertTaskInfo{
		AlertHash:                  alertHash,
		QuorumNumbers:              quorumNumbers,
		QuorumThresholdPercentages: quorumThresholdPercentages,
		TaskIndex:                  taskIndex,
		ReferenceBlockNumber:       referenceBlockNumber,
	}

	agg.tasksMu.Lock()
	agg.tasks[taskIndex] = newAlertTask
	agg.tasksMu.Unlock()

	// TODO(samlaf): we use seconds for now, but we should ideally pass a blocknumber to the blsAggregationService
	// and it should monitor the chain and only expire the task aggregation once the chain has reached that block number.
	taskTimeToExpiry := taskChallengeWindowBlock * blockTimeDuration

	agg.logger.Infof("InitializeNewTask %v %v", taskIndex, taskTimeToExpiry)
	err = agg.blsAggregationService.InitializeNewTask(
		taskIndex,
		uint32(newAlertTask.ReferenceBlockNumber),
		newAlertTask.QuorumNumbers,
		newAlertTask.QuorumThresholdPercentages,
		taskTimeToExpiry,
	)
	if err != nil {
		agg.logger.Error("InitializeNewTask failed", "err", err)
		return nil, err
	}
	return newAlertTask, nil
}
