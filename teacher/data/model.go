package data

import (
	"github.com/zahraftrm/mini-project/features/teacher"

	"gorm.io/gorm"
)

// struct gorm model
type Teacher struct {
	gorm.Model
	// ID        string `gorm:"primaryKey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
	NUPTK 	 	string `gorm:"unique"`
	Name     	string
	Phone    	string `gorm:"unique"`
	Expertise   string
	Email    	string `gorm:"unique" `
	Password 	string
}

// mapping dari core ke model
func CoreToModel(dataCore teacher.Core) Teacher {
	return Teacher{
		NUPTK:    	dataCore.NUPTK,
		Name:     	dataCore.Name,
		Phone:    	dataCore.Phone,
		Expertise:  dataCore.Expertise,
		Email:    	dataCore.Email,
		Password: 	dataCore.Password,
	}
}

func ModelToCore(dataModel Teacher) teacher.Core {
	return teacher.Core{
		Id:        	dataModel.ID,
		NUPTK:     	dataModel.NUPTK,
		Name:      	dataModel.Name,
		Phone:     	dataModel.Phone,
		Expertise:  dataModel.Expertise,
		Email:     	dataModel.Email,
		Password:  	dataModel.Password,
		CreatedAt: 	dataModel.CreatedAt,
		UpdatedAt: 	dataModel.UpdatedAt,
	}
}

func ModelToCoreList(dataModel []Teacher) []teacher.Core {
	var coreList []teacher.Core
	for _, v := range dataModel {
		coreList = append(coreList, ModelToCore(v))
	}
	return coreList
}
