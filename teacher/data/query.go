package data

import (
	"errors"
	"fmt"

	"github.com/zahraftrm/mini-project/app/middlewares"
	"github.com/zahraftrm/mini-project/features/teacher"
	"github.com/zahraftrm/mini-project/helper"

	"gorm.io/gorm"
)

type teacherQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) teacher.TeacherDataInterface {
	return &teacherQuery{
		db: db,
	}
}

// Delete implements teacher.TeacherDataInterface.
func (repo *teacherQuery) Delete(id int) (row int, err error) {
	result := repo.db.Delete(&Teacher{}, id)
	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("failed delete data. Data is not exits row = 0")
	}
	return int(result.RowsAffected), nil
}

// Update implements teacher.TeacherDataInterface.
func (repo *teacherQuery) Update(id int, input teacher.Core) (data teacher.Core, err error) {
	if input.Password != "" {
		hashedPassword, errHash := helper.HashPassword(input.Password)
		if errHash != nil {
			return teacher.Core{}, errors.New("error hash password")
		}
		input.Password = hashedPassword
	}

	tx := repo.db.Model(&Teacher{}).Where("id = ?", id).Updates(CoreToModel(input))
	if tx.Error != nil {
		return data, tx.Error
	}
	if tx.RowsAffected == 0 {
		return data, errors.New("failed update data, row affected = 0")
	}
	var teachersData Teacher
	resultFind := repo.db.Find(&teachersData, id)
	if resultFind.Error != nil {
		return teacher.Core{}, resultFind.Error
	}
	data = ModelToCore(teachersData)
	return data, nil
}

// SelectById implements teacher.TeacherDataInterface.
func (repo *teacherQuery) SelectById(id int) (teacher.Core, error) {
	var teachersData Teacher
	tx := repo.db.First(&teachersData) // select * from teachers
	if tx.Error != nil {
		return teacher.Core{}, tx.Error
	}
	// mapping dari struct gorm model ke struct entities core
	var teachersCore = ModelToCore(teachersData)
	fmt.Println(teachersCore)

	return teachersCore, nil
}

// Login implements teacher.TeacherDataInterface
func (repo *teacherQuery) Login(email string, password string) (teacher.Core, string, error) {
	var teacherGorm Teacher
	tx := repo.db.Where("email = ?", email).First(&teacherGorm) // select * from teachers limit 1
	if tx.Error != nil {
		return teacher.Core{}, "", tx.Error
	}

	if tx.RowsAffected == 0 {
		return teacher.Core{}, "", errors.New("login failed, email dan password salah")
	}

	checkPassword := helper.CheckPasswordHash(password, teacherGorm.Password)
	if !checkPassword {
		return teacher.Core{}, "", errors.New("login failed, password salah")
	}

	token, errToken := middlewares.CreateToken(int(teacherGorm.ID))
	if errToken != nil {
		return teacher.Core{}, "", errToken
	}

	dataCore := ModelToCore(teacherGorm)
	return dataCore, token, nil
}

// Insert implements teacher.TeacherDataInterface
func (repo *teacherQuery) Insert(input teacher.Core) error {
	// mapping dari struct entities core ke gorm model
	// teacherInputGorm := Teacher{
	//	NUPTK:    	input.NUPTK,
	// 	Name:     	input.Name,
	// 	Phone:    	input.Phone,
	// 	Expertise:  input.Expertise,
	// 	Email:    	input.Email,
	// 	Password: 	input.Password,
	// }
	hashedPassword, errHash := helper.HashPassword(input.Password)
	if errHash != nil {
		return errors.New("error hash password")
	}
	teacherInputGorm := CoreToModel(input)
	teacherInputGorm.Password = hashedPassword

	tx := repo.db.Create(&teacherInputGorm) // insert into teachers set name = .....
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}

	return nil
}

// SelectAll implements teacher.TeacherDataInterface
func (repo *teacherQuery) SelectAll() ([]teacher.Core, error) {
	var teachersData []Teacher
	tx := repo.db.Find(&teachersData) // select * from teacher
	if tx.Error != nil {
		return nil, tx.Error
	}
	fmt.Println(teachersData)
	// mapping dari struct gorm model ke struct entities core
	var teachersCoreAll []teacher.Core = ModelToCoreList(teachersData)
	// for _, value := range teachersData {
	// 	var teacherCore = teacher.Core{
	// 		Id:        	value.ID,
	// 		NUPTK:     	value.NUPTK,
	//		Name:      	value.Name,
	// 		Phone:     	value.Phone,	
	// 		Expertise:  value.Expertise,
	// 		Email:     	value.Email,
	// 		Password:  	value.Password,
	// 		CreatedAt: 	value.CreatedAt,
	// 		UpdatedAt: 	value.UpdatedAt,
	// 	}
	// 	teachersCoreAll = append(teachersCoreAll, teacherCore)
	// }
	return teachersCoreAll, nil
}
