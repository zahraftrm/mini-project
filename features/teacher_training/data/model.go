package data

import (
	"time"

	"github.com/zahraftrm/mini-project/features/revision"
	training "github.com/zahraftrm/mini-project/features/teacher_training"

	_revisionModel "github.com/zahraftrm/mini-project/features/revision/data"

	"github.com/zahraftrm/mini-project/features/teacher/data"

	"gorm.io/gorm"
)

// struct gorm model
type Training struct {
	gorm.Model
	Teacher     data.Teacher `gorm:"foreignKey:TeacherID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TeacherID   uint
	Title        string  
	Expertise    string  
	Category 	 string  
	Description	 string  
	Cost 		 string  
	StartDate   time.Time
	EndDate     time.Time
	Location	 string  
	Status 		 string 
	Revisions   []_revisionModel.Revision `gorm:"foreignKey:TrainingID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

// mapping dari core ke model
func CoreToModel(dataCore training.Core) Training {
	return Training{
		TeacherID:   	dataCore.TeacherID,
		Title:        	dataCore.Title,
		Expertise: 		dataCore.Expertise,
		Category: 		dataCore.Category,
		Description: 	dataCore.Description,
		Cost: 			dataCore.Cost,
		StartDate:   	dataCore.StartDate,
		EndDate:     	dataCore.EndDate,
		Location: 		dataCore.Location,
		Status: 		dataCore.Status,
	}
}

func ListRevisionModelToTrainingCore(dataModel Training) []revision.Core {
	var result []revision.Core
	for _, v := range dataModel.Revisions {
		var revisionCore = revision.Core{
			Id:        	v.ID,
			TrainingID: v.TrainingID,
			Message:  	v.Message,
			CreatedAt: 	v.CreatedAt,
			UpdatedAt: 	v.UpdatedAt,
		}

		result = append(result, revisionCore)
	}
	return result
}


func ModelToCore(dataModel Training) training.Core {
	
	return training.Core{
		Id:          	dataModel.ID,
		TeacherID:      dataModel.TeacherID,
		Title:        	dataModel.Title,
		Expertise: 		dataModel.Expertise,
		Category: 		dataModel.Category,
		Description: 	dataModel.Description,
		Cost: 			dataModel.Cost,
		StartDate:   	dataModel.StartDate,
		EndDate:     	dataModel.EndDate,
		Location: 		dataModel.Location,
		Status: 		dataModel.Status,
		
		Teacher: training.Teacher{
			Id:    dataModel.TeacherID,
			Name:  dataModel.Teacher.Name,
			Email: dataModel.Teacher.Email,
		},
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
		Revisions: ListRevisionModelToTrainingCore(dataModel),
	}
}


func ModelToCoreList(dataModel []Training) []training.Core {
	
	var coreList []training.Core
	for _, v := range dataModel {
		coreList = append(coreList, ModelToCore(v))
	}
	return coreList
}
