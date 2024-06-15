package strategy

import "log"

type TransferOperator struct{}

func (r *TransferOperator) UpdateTask(req *UploadOperationRequest) (bool, error) {
	log.Println("这里是TransferOperator")
	return true, nil
}

func newTransferOperator() *TransferOperator {
	return &TransferOperator{}
}
