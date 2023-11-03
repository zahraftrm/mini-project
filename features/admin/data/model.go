package data

import (
	"github.com/zahraftrm/mini-project/features/admin"

	"gorm.io/gorm"
)

// struct gorm model
type Admin struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique" `
	Password string
}

// mapping dari core ke model
func CoreToModel(dataCore admin.Core) Admin {
	return Admin{
		Name:     dataCore.Name,
		Email:    dataCore.Email,
		Password: dataCore.Password,
	}
}

func ModelToCore(dataModel Admin) admin.Core {
	return admin.Core{
		Id:        dataModel.ID,
		Name:      dataModel.Name,
		Email:     dataModel.Email,
		Password:  dataModel.Password,
		CreatedAt: dataModel.CreatedAt,
		UpdatedAt: dataModel.UpdatedAt,
	}
}

func ModelToCoreList(dataModel []Admin) []admin.Core {
	var coreList []admin.Core
	for _, v := range dataModel {
		coreList = append(coreList, ModelToCore(v))
	}
	return coreList
}
