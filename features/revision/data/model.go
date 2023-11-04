package data

import (
	"github.com/zahraftrm/mini-project/features/revision"
	"gorm.io/gorm"
)

// struct gorm model
type Revision struct {
	gorm.Model
	TrainingID uint
	Message  string
}

// mapping dari core ke model
func CoreToModel(dataCore revision.Core) Revision {
	return Revision{
		TrainingID: dataCore.TrainingID,
		Message:  dataCore.Message,
	}
}

func ModelToCore(dataModel Revision) revision.Core {
	return revision.Core{
		Id:        dataModel.ID,
		TrainingID: dataModel.TrainingID,
		Message:  dataModel.Message,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

func ModelToCoreList(dataModel []Revision) []revision.Core {
	var coreList []revision.Core
	for _, v := range dataModel {
		coreList = append(coreList, ModelToCore(v))
	}
	return coreList
}
