package rpc

import (
	"github.com/HemeraProtocol/avs/legacy/core/message"
)

type AggregatorRpcHandler interface {
	InitOperator(req *message.InitOperatorRequest) (*message.InitOperatorResponse, error)
	CreateTask(req *message.CreateTaskRequest) (*message.CreateTaskResponse, error)
	ProcessSignedTaskResponse(signedTaskResponse *message.SignedTaskRespRequest) (*message.SignedTaskRespResponse, error)
}
