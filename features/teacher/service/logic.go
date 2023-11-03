package service

import (
	"errors"
	"math/rand"
	"strings"
	"time"

	"github.com/zahraftrm/mini-project/features/teacher"

	"github.com/go-playground/validator/v10"
)

type teacherService struct {
	teacherData teacher.TeacherDataInterface
	validate *validator.Validate
}

func New(repo teacher.TeacherDataInterface) teacher.TeacherServiceInterface {
	return &teacherService{
		teacherData: repo,
		validate: validator.New(),
	}
}

func isNumeric(s string) bool {
    for _, char := range s {
        if char < '0' || char > '9' {
            return false
        }
    }
    return true
}

// Delete implements teacher.TeacherServiceInterface.
func (service *teacherService) Delete(id int) (err error) {
	if id == 0 {
		return errors.New("id not found")
	}
	row, errData := service.teacherData.Delete(id)
	if errData != nil || row == 0 {
		return errData
	}
	return nil
}

// Update implements teacher.TeacherServiceInterface.
func (service *teacherService) Update(id int, input teacher.Core) (data teacher.Core, err error) {
	if id == 0 {
		return teacher.Core{}, errors.New("id not found")
	}

	// validasi atribut "phone"
    if len(input.Phone) < 10 || len(input.Phone) > 14 || !isNumeric(input.Phone) {
        return teacher.Core{}, errors.New(": nomor telepon harus berupa angka antara 10-14 karakter")
    }

	if errValidate := service.validate.Struct(input); errValidate != nil {
		return teacher.Core{}, errValidate
	}

	data, err = service.teacherData.Update(id, input)
	if err != nil {
		return teacher.Core{}, err
	}
	return
}

// GetById implements teacher.TeacherServiceInterface.
func (service *teacherService) GetById(id int) (teacher.Core, error) {
	if id == 0 {
		return teacher.Core{}, errors.New("invalid id")
	}
	result, err := service.teacherData.SelectById(id)
	return result, err
}

// Login implements teacher.TeacherServiceInterface
func (service *teacherService) Login(email string, password string) (teacher.Core, string, error) {
	if email == "" || password == "" {
		return teacher.Core{}, "", errors.New("error validation: nama, email, password harus diisi")
	}
	dataLogin, token, err := service.teacherData.Login(email, password)
	return dataLogin, token, err
}

// Create implements teacher.TeacherServiceInterface
func (service *teacherService) Create(input teacher.Core) error {
	// validasi atribut "phone"
    if len(input.Phone) < 10 || len(input.Phone) > 14 || !isNumeric(input.Phone) {
        return errors.New(": nomor telepon harus berupa angka antara 10-14 karakter")
    }
	
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate
	}
	errInsert := service.teacherData.Insert(input)
	return errInsert
}

// GetAll implements teacher.TeacherServiceInterface
func (service *teacherService) GetAll() ([]teacher.Core, error) {
	data, err := service.teacherData.SelectAll()
	if err != nil {
		return nil, err
	}
	return data, err
}

func RandomPassword() string {
	// posibilitas karakter
	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := strings.ToLower(lower)
	symbols := "!@#$%^&*()"
	numbers := "0123456789"
	all := lower + upper + symbols + numbers

	length := 6

	// random
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	password := ""

	for i := 0; i < length; i++ {
		password += string(all[random.Intn(len(all))])
	}

	return password
}