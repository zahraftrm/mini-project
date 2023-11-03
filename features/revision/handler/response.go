package handler

import (
	"time"

	"github.com/zahraftrm/mini-project/features/revision"
)

type RevisionResponse struct {
	Id        uint      `json:"id"`
	TrainingID uint      `json:"training_id"`
	Message  string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

func CoreToResponse(dataCore revision.Core) RevisionResponse {
	return RevisionResponse{
		Id:        dataCore.Id,
		TrainingID: dataCore.TrainingID,
		Message:  dataCore.Message,
		CreatedAt: dataCore.CreatedAt,
	}
}

func CoreToResponseList(dataCore []revision.Core) []RevisionResponse {
	var result []RevisionResponse
	for _, v := range dataCore {
		result = append(result, CoreToResponse(v))
	}
	return result
}
