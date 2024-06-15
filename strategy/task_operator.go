package strategy

type UploadOperationRequest struct {
	HandleType int64
	Data       string
}

type TaskOperator interface {
	UpdateTask(req *UploadOperationRequest) (bool, error)
}
