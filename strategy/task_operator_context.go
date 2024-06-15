package strategy

import "errors"

var operatorMap = make(map[int64]TaskOperator)

func init() {
	operatorMap[1] = &TransferOperator{}
	operatorMap[2] = &ManualSubmit{}
}

type TaskOperatorContext struct {
	handler TaskOperator
	req     *UploadOperationRequest
}

func NewTaskContext(req *UploadOperationRequest) (*TaskOperatorContext, error) {
	handlerType := req.HandleType
	if operator, ok := operatorMap[handlerType]; ok {
		return &TaskOperatorContext{
			req:     req,
			handler: operator,
		}, nil
	}
	return nil, errors.New("类型不存在")
}

func (r *TaskOperatorContext) Update() (bool, error) {
	return r.handler.UpdateTask(r.req)
}
