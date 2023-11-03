package data

import (
	"time"

	"github.com/zahraftrm/mini-project/features/revision"
	training "github.com/zahraftrm/mini-project/features/training"

	_revisionModel "github.com/zahraftrm/mini-project/features/revision/data"

	dataAdmin "github.com/zahraftrm/mini-project/features/admin/data"
	dataTeacher "github.com/zahraftrm/mini-project/features/teacher/data"

	"gorm.io/gorm"
)

// struct gorm model
type Training struct {
	gorm.Model
	Teacher      dataTeacher.Teacher `gorm:"foreignKey:TeacherID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Admin     	 dataAdmin.Admin	 `gorm:"foreignKey:AdminID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TeacherID    uint
	AdminID      uint	`gorm:"default:NULL"`
	Title        string  
	Expertise    string  
	Category 	 string  
	Description	 string  
	Cost 		 string  
	StartDate   time.Time
	EndDate     time.Time
	Location	 string  
	Status 		 string `gorm:"default:'menunggu persetujuan'"`
	Revisions   []_revisionModel.Revision `gorm:"foreignKey:TrainingID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	TeacherRole	string `gorm:"-"`
	AdminRole	string `gorm:"-"`

}

// mapping dari core ke model
func CoreToModel(dataCore training.Core, role string) Training {
	if role == "teacher" {
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
	} else if role == "admin" {
		return Training{
			TeacherID:   	dataCore.TeacherID,
			AdminID:   		dataCore.AdminID,
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
	
	return Training{}
}

func ListRevisionModelToTrainingCore(dataModel Training) []revision.Core {
	var result []revision.Core
	if len(dataModel.Revisions) > 0 {
    	latestRevision := dataModel.Revisions[len(dataModel.Revisions)-1]
		
    	var revisionCore = revision.Core{
    	    Id:        latestRevision.ID,
    	    TrainingID: latestRevision.TrainingID,
    	    Message:   latestRevision.Message,
    	    CreatedAt: latestRevision.CreatedAt,
    	    UpdatedAt: latestRevision.UpdatedAt,
    	}
	
    	result = append(result, revisionCore)
	}
	return result
}


func ModelToCore(dataModel Training) training.Core {
	
	return training.Core{
		Id:          	dataModel.ID,
		TeacherID:      dataModel.TeacherID,
		AdminID:      	dataModel.AdminID,
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
		Admin: training.Admin{
			Id:    dataModel.AdminID,
			Name:  dataModel.Admin.Name,
			Email: dataModel.Admin.Email,
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
