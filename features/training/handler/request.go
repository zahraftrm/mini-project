package handler

import (
	"errors"
	"time"

	"github.com/zahraftrm/mini-project/constants"
	training "github.com/zahraftrm/mini-project/features/training"
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

func RequestToCore(dataRequest TrainingRequest, role string) (training.Core, error) {
	if dataRequest.Category == "offline" && dataRequest.Location == "" {
		return training.Core{}, errors.New("location harus diisi jika category adalah offline")
	}

    if role == constants.RolesAdmin {
        return training.Core{
            Title:       dataRequest.Title,
            Category:    dataRequest.Category,
            Description: dataRequest.Description,
            Cost:        dataRequest.Cost,
            Status:      dataRequest.Status,
        }, nil
    } else if role == constants.RolesTeacher {
        return training.Core{
            Title:       dataRequest.Title,
            Expertise:   dataRequest.Expertise,
            Category:    dataRequest.Category,
            Description: dataRequest.Description,
            Cost:        dataRequest.Cost,
            StartDate:   time.Time{},
            EndDate:     time.Time{},
            Location:    dataRequest.Location,
        }, nil
    }

    return training.Core{}, errors.New("role tidak valid")
}
