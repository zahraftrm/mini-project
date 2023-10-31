package handler

import (
	"time"

	_revisionHandler "github.com/zahraftrm/mini-project/features/revision/handler"
	training "github.com/zahraftrm/mini-project/features/teacher_training"
)

type TrainingResponse struct {
	Id          	uint                        `json:"id"`
	Title        	string                      `json:"title"`
	Expertise    	string                      `json:"expertise"`
	Category 	 	string                      `json:"category"`
	Description 	string                      `json:"description"`
	Cost 		 	string                      `json:"cost"`
	StartDate   	time.Time                   `json:"start_date"`
	EndDate     	time.Time                   `json:"end_date"`
	Location 		string                      `json:"location"`
	Status 		 	string                     `json:"status"`
	TeacherID     	uint                        `json:"teacher_id,omitempty"`
	TeacherName    	string                      `json:"teacher_name,omitempty"`
	CreatedAt   	time.Time                   `json:"created_at"`
	Revisions       []_revisionHandler.RevisionResponse `json:"revisions,omitempty"`
}

func ListRevisionCoreToResponse(dataCore training.Core) []_revisionHandler.RevisionResponse {
	var result []_revisionHandler.RevisionResponse
	for _, v := range dataCore.Revisions {
		revision := _revisionHandler.RevisionResponse{
			Id:        	v.Id,
			TrainingID: v.TrainingID,
			Message:  	v.Message,
			CreatedAt: 	v.CreatedAt,
		}
		result = append(result, revision)
	}
	return result
}
func CoreToResponse(dataCore training.Core) TrainingResponse {
	return TrainingResponse{
		Id: 	         	dataCore.Id,
		Title:  	     	dataCore.Title,
		Expertise: 		 	dataCore.Expertise,
		Category: 		 	dataCore.Category,
		Description: 	 	dataCore.Description,
		Cost: 			 	dataCore.Cost,
		StartDate:   		dataCore.StartDate,
		EndDate:     		dataCore.EndDate,
		Location: 		 	dataCore.Location,
		Status: 		 	dataCore.Status,
		TeacherID:      	dataCore.TeacherID,
		TeacherName:    	dataCore.Teacher.Name,
		CreatedAt:   		dataCore.CreatedAt,
		Revisions:       	ListRevisionCoreToResponse(dataCore),
	}
}

func CoreToResponseList(dataCore []training.Core) []TrainingResponse {
	var result []TrainingResponse
	for _, v := range dataCore {
		result = append(result, CoreToResponse(v))
	}
	return result
}
