package service

import (
	"errors"

	training "github.com/zahraftrm/mini-project/features/training"

	"github.com/go-playground/validator/v10"
)

type trainingService struct {
	trainingData training.TrainingDataInterface
	validate    *validator.Validate
}

func New(repo training.TrainingDataInterface) training.TrainingServiceInterface {
	return &trainingService{
		trainingData: repo,
		validate:    validator.New(),
	}
}

// Create implements teacher.TeacherServiceInterface
func (service *trainingService) Create(input training.Core, role string) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}
	errInsert := service.trainingData.Insert(input, role)

	return errInsert
}

// GetAll implements teacher.TeacherServiceInterface
func (service *trainingService) GetAll(teacherid int) ([]training.Core, error) {
	data, err := service.trainingData.SelectAll(teacherid)
	if err != nil {
		return nil, err
	}
	return data, err
}

// GetAllByAdmin 
func (service *trainingService) GetAllByAdmin() ([]training.Core, error) {
	data, err := service.trainingData.SelectAllByAdmin()
	if err != nil {
		return nil, err
	}
	return data, err
}

// GetById implements teacher.TeacherServiceInterface.
func (service *trainingService) GetById(id int, teacherid int) (training.Core, error) {
	if id == 0 {
		return training.Core{}, errors.New("invalid id")
	}
	result, err := service.trainingData.SelectById(id, teacherid)
	return result, err
}

func (service *trainingService) GetByIdTeacher(teacher_id int) ([]training.Core, error) {
	if teacher_id == 0 {
		return []training.Core{}, errors.New("invalid id")
	}
	result, err := service.trainingData.SelectByIdTeacher(teacher_id)
	return result, err
}

func (service *trainingService) GetByIdTraining(id int) (training.Core, error) {
	if id == 0 {
		return training.Core{}, errors.New("invalid id")
	}
	result, err := service.trainingData.SelectByIdTraining(id)
	return result, err
}

// Update implements teacher.TeacherServiceInterface.
func (service *trainingService) Update(id int, adminid int, role string, input training.Core) (data training.Core, email string, name string, err error) {
	if id == 0 {
		return training.Core{}, "", "", errors.New("id not found")
	}

	if errValidate := service.validate.Struct(input); errValidate != nil {
		return training.Core{}, "", "", errValidate
	}

	data, email, name, err = service.trainingData.Update(id, adminid, role, input)
	if err != nil {
		return training.Core{}, "", "", err
	}
	return
}

// Delete implements teacher.TeacherServiceInterface.
func (service *trainingService) Delete(id int, teacherid int) (err error) {
	if id == 0 {
		return errors.New("id not found")
	}
	row, errData := service.trainingData.Delete(id, teacherid)
	if errData != nil || row == 0 {
		return errData
	}
	return nil
}

func (service *trainingService) DeleteByAdmin(id int) (err error) {
	if id == 0 {
		return errors.New("id not found")
	}
	row, errData := service.trainingData.DeleteByAdmin(id)
	if errData != nil || row == 0 {
		return errData
	}
	return nil
}
