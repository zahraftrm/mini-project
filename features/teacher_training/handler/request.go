package handler

import (
	"time"

	training "github.com/zahraftrm/mini-project/features/teacher_training"
)

type TrainingRequest struct {
	Title       string `json:"title" form:"title"`
	Expertise	string `json:"expertise" form:"expertise"`
	Category	string `json:"category" form:"category"`
	Description string `json:"description" form:"description"`
	Cost 		string `json:"cost" form:"cost"`
	StartDate   string `json:"start_date" form:"start_date"`
	EndDate     string `json:"end_date" form:"end_date"`
	Location 	string `json:"location" form:"location"`
	Status 		string `json:"status" form:"status"`
}

func RequestToCore(dataRequest TrainingRequest) training.Core {
	return training.Core{
		Title: 			dataRequest.Title,
		Expertise: 		dataRequest.Expertise,
		Category: 		dataRequest.Category,
		Description:	dataRequest.Description,
		Cost: 			dataRequest.Cost,
		StartDate:   	time.Time{},
		EndDate:     	time.Time{},
		Location: 		dataRequest.Location,
		Status: 		dataRequest.Status,
	}
}
