package strategy

import "log"

type ManualSubmit struct{}

func (r *ManualSubmit) UpdateTask(req *UploadOperationRequest) (bool, error) {
	log.Println("这里是ManualSubmit")
	return true, nil
}
