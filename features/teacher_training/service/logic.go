package service

import (
	"errors"

	training "github.com/zahraftrm/mini-project/features/teacher_training"

	"github.com/go-playground/validator/v10"
)

type trainingService struct {
	trainingData training.TeacherTrainingDataInterface
	validate    *validator.Validate
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

// Update implements teacher.TeacherServiceInterface.
func (service *trainingService) Update(id int, teacherid int, input training.Core) (data training.Core, err error) {
	if id == 0 {
		return training.Core{}, errors.New("id not found")
	}

	if errValidate := service.validate.Struct(input); errValidate != nil {
		return training.Core{}, errValidate
	}

	data, err = service.trainingData.Update(id, teacherid, input)
	if err != nil {
		return training.Core{}, err
	}
	return
}

// GetById implements teacher.TeacherServiceInterface.
func (service *trainingService) GetById(id int, teacherid int) (training.Core, error) {
	if id == 0 {
		return training.Core{}, errors.New("invalid id")
	}
	result, err := service.trainingData.SelectById(id, teacherid)
	return result, err
}

// Create implements teacher.TeacherServiceInterface
func (service *trainingService) Create(input training.Core) error {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}
	errInsert := service.trainingData.Insert(input)
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

func New(repo training.TeacherTrainingDataInterface) training.TeacherTrainingServiceInterface {
	return &trainingService{
		trainingData: repo,
		validate:    validator.New(),
	}
}
