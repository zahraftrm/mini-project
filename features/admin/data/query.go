package data

import (
	"errors"
	"fmt"

	"github.com/zahraftrm/mini-project/app/middlewares"
	"github.com/zahraftrm/mini-project/features/admin"
	"github.com/zahraftrm/mini-project/helper"

	"gorm.io/gorm"
)

type adminQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) admin.AdminDataInterface {
	return &adminQuery{
		db: db,
	}
}

// Delete implements admin.AdminDataInterface.
func (repo *adminQuery) Delete(id int) (row int, err error) {
	result := repo.db.Delete(&Admin{}, id)
	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("failed delete data. Data is not exits row = 0")
	}
	return int(result.RowsAffected), nil
}

// Update implements admin.AdminDataInterface.
func (repo *adminQuery) Update(id int, input admin.Core) (data admin.Core, err error) {
	if input.Password != "" {
		hashedPassword, errHash := helper.HashPassword(input.Password)
		if errHash != nil {
			return admin.Core{}, errors.New("error hash password")
		}
		input.Password = hashedPassword
	}

	tx := repo.db.Model(&Admin{}).Where("id = ?", id).Updates(CoreToModel(input))
	if tx.Error != nil {
		return data, tx.Error
	}
	if tx.RowsAffected == 0 {
		return data, errors.New("failed update data, row affected = 0")
	}
	var adminsData Admin
	resultFind := repo.db.Find(&adminsData, id)
	if resultFind.Error != nil {
		return admin.Core{}, resultFind.Error
	}
	data = ModelToCore(adminsData)
	return data, nil
}

// SelectById implements admin.AdminDataInterface.
func (repo *adminQuery) SelectById(id int) (admin.Core, error) {
	var adminsData Admin
	tx := repo.db.Where("id = ?", id).First(&adminsData) // select * from admins
	if tx.Error != nil {
		return admin.Core{}, tx.Error
	}
	// mapping dari struct gorm model ke struct entities core
	var adminsCore = ModelToCore(adminsData)
	fmt.Println(adminsCore)

	return adminsCore, nil
}

// Login implements admin.AdminDataInterface
func (repo *adminQuery) Login(email string, password string) (admin.Core, string, error) {
	var adminGorm Admin
	tx := repo.db.Where("email = ?", email).First(&adminGorm) // select * from admins limit 1
	if tx.Error != nil {
		return admin.Core{}, "", tx.Error
	}

	if tx.RowsAffected == 0 {
		return admin.Core{}, "", errors.New("login failed, email dan password salah")
	}

	checkPassword := helper.CheckPasswordHash(password, adminGorm.Password)
	if !checkPassword {
		return admin.Core{}, "", errors.New("login failed, password salah")
	}

	token, errToken := middlewares.CreateToken(int(adminGorm.ID), "admin")
	if errToken != nil {
		return admin.Core{}, "", errToken
	}

	dataCore := ModelToCore(adminGorm)
	return dataCore, token, nil
}

// Insert implements admin.AdminDataInterface
func (repo *adminQuery) Insert(input admin.Core) error {
	hashedPassword, errHash := helper.HashPassword(input.Password)
	if errHash != nil {
		return errors.New("error hash password")
	}
	adminInputGorm := CoreToModel(input)
	adminInputGorm.Password = hashedPassword

	tx := repo.db.Create(&adminInputGorm) // insert into admins set name = .....
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}

	return nil
}

// SelectAll implements admin.AdminDataInterface
func (repo *adminQuery) SelectAll() ([]admin.Core, error) {
	var adminsData []Admin
	tx := repo.db.Find(&adminsData) // select * from admin
	if tx.Error != nil {
		return nil, tx.Error
	}
	fmt.Println(adminsData)
	// mapping dari struct gorm model ke struct entities core
	var adminsCoreAll []admin.Core = ModelToCoreList(adminsData)
	
	return adminsCoreAll, nil
}
