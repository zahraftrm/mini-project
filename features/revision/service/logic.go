package service

import (
	"errors"

	"github.com/zahraftrm/mini-project/features/revision"

	"github.com/go-playground/validator/v10"
)

type revisionService struct {
	revisionData revision.RevisionDataInterface
	validate *validator.Validate
}

// Delete implements admin.AdminServiceInterface.
func (service *revisionService) Delete(id int, adminid int) (err error) {
	if id == 0 {
		return errors.New("id not found")
	}
	row, errData := service.revisionData.Delete(id, adminid)
	if errData != nil || row == 0 {
		return errData
	}
	return nil
}

// Update implements admin.AdminServiceInterface.
func (service *revisionService) Update(id int, adminid int, input revision.Core) (data revision.Core, err error) {
	if id == 0 {
		return revision.Core{}, errors.New("id not found")
	}

	if errValidate := service.validate.Struct(input); errValidate != nil {
		return revision.Core{}, errValidate
	}

	data, err = service.revisionData.Update(id, adminid, input)
	if err != nil {
		return revision.Core{}, err
	}
	return
}

// Create implements admin.AdminServiceInterface
func (service *revisionService) Create(input revision.Core) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}
	errInsert := service.revisionData.Insert(input)
	return errInsert
}

func New(repo revision.RevisionDataInterface) revision.RevisionServiceInterface {
	return &revisionService{
		revisionData: repo,
		validate: validator.New(),
	}
}
