package data

import (
	"errors"

	"github.com/zahraftrm/mini-project/features/revision"

	"gorm.io/gorm"
)

type revisionQuery struct {
	db *gorm.DB
}

type trainingQuery struct {
	db *gorm.DB
}

type Training struct {
	gorm.Model
	Status 		 string 
}

// Update status implements training.TrainingDataInterface.
func (repo *trainingQuery) UpdateStatusByTrainingID(trainingID uint, newStatus string) error {
    // Menggunakan tipe data dari paket training/data.
    tx := repo.db.Model(&Training{}).Where("id = ?", trainingID).Update("Status", newStatus)
    if tx.Error != nil {
        return tx.Error
    }
    return nil
}


func New(db *gorm.DB) revision.RevisionDataInterface {
	return &revisionQuery{
		db: db,
	}
}


// Delete implements admin.AdminDataInterface.
func (repo *revisionQuery) Delete(id int) (row int, err error) {
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
func (repo *revisionQuery) Update(id int, input revision.Core) (data revision.Core, err error) {

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

	// memeriksa apakah data dengan TrainingID tertentu sudah ada di tabel revision
    var existingRevisionCount int64
    if err := repo.db.Model(&Revision{}).
        Where("training_id = ?", input.TrainingID).
        Count(&existingRevisionCount).Error; err != nil {
        return err
    }

    if existingRevisionCount > 0 {
        return errors.New(": revision dengan training id tersebut sudah ada, silahkan lakukan update data")
    }

    tx := repo.db.Create(&inputGorm) // insert into revisions
    if tx.Error != nil {
        return tx.Error
    }

    if tx.RowsAffected == 0 {
        return errors.New("insert failed, row affected = 0")
    }

    // Membuat objek trainingQuery
    trainingRepo := &trainingQuery{db: repo.db}

    // Update Status pada Training yang sesuai dengan TrainingID.
    err := trainingRepo.UpdateStatusByTrainingID(input.TrainingID, "revisi")

    if err != nil {
        return err
    }

    return nil
}
