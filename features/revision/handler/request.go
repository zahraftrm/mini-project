package handler

import (
	"github.com/zahraftrm/mini-project/features/revision"
)

type RevisionRequest struct {
	TrainingID uint   `json:"training_id" form:"training_id"`
	Message  string `json:"message" form:"message"`
}

func RequestToCore(dataRequest RevisionRequest) revision.Core {
	return revision.Core{
		TrainingID: dataRequest.TrainingID,
		Message:  dataRequest.Message,
	}
}
