package service

import (
	"errors"

	"github.com/zahraftrm/mini-project/features/admin"

	"github.com/go-playground/validator/v10"
)

type adminService struct {
	adminData admin.AdminDataInterface
	validate *validator.Validate
}

func New(repo admin.AdminDataInterface) admin.AdminServiceInterface {
	return &adminService{
		adminData: repo,
		validate: validator.New(),
	}
}

// Delete implements admin.AdminServiceInterface.
func (service *adminService) Delete(id int) (err error) {
	if id == 0 {
		return errors.New("id not found")
	}
	row, errData := service.adminData.Delete(id)
	if errData != nil || row == 0 {
		return errData
	}
	return nil
}

// Update implements admin.AdminServiceInterface.
func (service *adminService) Update(id int, input admin.Core) (data admin.Core, err error) {
	if id == 0 {
		return admin.Core{}, errors.New("id not found")
	}

	if errValidate := service.validate.Struct(input); errValidate != nil {
		return admin.Core{}, errValidate
	}

	data, err = service.adminData.Update(id, input)
	if err != nil {
		return admin.Core{}, err
	}
	return
}

// GetById implements admin.AdminServiceInterface.
func (service *adminService) GetById(id int) (admin.Core, error) {
	if id == 0 {
		return admin.Core{}, errors.New("invalid id")
	}
	result, err := service.adminData.SelectById(id)
	return result, err
}

// Login implements admin.AdminServiceInterface
func (service *adminService) Login(email string, password string) (admin.Core, string, error) {
	if email == "" || password == "" {
		return admin.Core{}, "", errors.New("error validation: nama, email, password harus diisi")
	}
	dataLogin, token, err := service.adminData.Login(email, password)
	
	return dataLogin, token, err
}

// Create implements admin.AdminServiceInterface
func (service *adminService) Create(input admin.Core) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}
	errInsert := service.adminData.Insert(input)
	return errInsert
}

// GetAll implements admin.AdminServiceInterface
func (service *adminService) GetAll() ([]admin.Core, error) {
	data, err := service.adminData.SelectAll()
	if err != nil {
		return nil, err
	}
	return data, err
}
