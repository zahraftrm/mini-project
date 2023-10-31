package data

import (
	"errors"

	"github.com/zahraftrm/mini-project/features/revision"

	"gorm.io/gorm"
)

type revisionQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) revision.RevisionDataInterface {
	return &revisionQuery{
		db: db,
	}
}

// Delete implements admin.AdminDataInterface.
func (repo *revisionQuery) Delete(id int, adminid int) (row int, err error) {
	result := repo.db.Delete(&Revision{}, id)
	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("failed delete data. Data is not exits row = 0")
	}
	return int(result.RowsAffected), nil
}

// Update implements admin.AdminDataInterface.
func (repo *revisionQuery) Update(id int, adminid int, input revision.Core) (data revision.Core, err error) {

	tx := repo.db.Model(&Revision{}).Where("id = ?", id).Updates(CoreToModel(input))
	if tx.Error != nil {
		return data, tx.Error
	}
	if tx.RowsAffected == 0 {
		return data, errors.New("failed update data, row affected = 0")
	}
	var resultData Revision
	resultFind := repo.db.Find(&resultData, id)
	if resultFind.Error != nil {
		return revision.Core{}, resultFind.Error
	}
	data = ModelToCore(resultData)
	return data, nil
}

// Insert implements admin.AdminDataInterface
func (repo *revisionQuery) Insert(input revision.Core) error {

	inputGorm := CoreToModel(input)

	tx := repo.db.Create(&inputGorm) // insert into admins 
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}

	return nil
}
