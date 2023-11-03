package handler

import (
	"time"

	_revisionHandler "github.com/zahraftrm/mini-project/features/revision/handler"
	training "github.com/zahraftrm/mini-project/features/training"
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
	Location 		string                      `json:"location,omitempty"`
	Status 		 	string                      `json:"status"`
	TeacherID     	uint                        `json:"teacher_id,omitempty"`
	TeacherName    	string                      `json:"teacher_name,omitempty"`
	AdminID     	uint                        `json:"admin_id,omitempty"`
	AdminName    	string                      `json:"admin_name,omitempty"`
	CreatedAt   	time.Time                   `json:"created_at"`
	Revisions       []_revisionHandler.RevisionResponse `json:"revisions,omitempty"`
}

func ListRevisionCoreToResponse(dataCore training.Core) []_revisionHandler.RevisionResponse {
	var result []_revisionHandler.RevisionResponse
	if len(dataCore.Revisions) > 0 {
    	latestRevision := dataCore.Revisions[len(dataCore.Revisions)-1]
		
    	var revisionCore = _revisionHandler.RevisionResponse{
    	    Id:        latestRevision.Id,
    	    TrainingID: latestRevision.TrainingID,
    	    Message:   latestRevision.Message,
    	    CreatedAt: latestRevision.CreatedAt,
    	}
	
    	result = append(result, revisionCore)
	}
	return result
}

func CoreToResponse(dataCore training.Core) TrainingResponse {
	if dataCore.Status == "disetujui" {
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
			AdminID:      		dataCore.AdminID,
			AdminName:    		dataCore.Admin.Name,
			CreatedAt:   		dataCore.CreatedAt,
			Revisions:       	ListRevisionCoreToResponse(dataCore),
		}
	} else {
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
}

func CoreToResponseList(dataCore []training.Core) []TrainingResponse {
	var result []TrainingResponse
	for _, v := range dataCore {
		result = append(result, CoreToResponse(v))
	}
	return result
}
