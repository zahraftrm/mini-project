package data

import (
	"errors"
	"fmt"

	training "github.com/zahraftrm/mini-project/features/teacher_training"
	"gorm.io/gorm"
)

type trainingQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) training.TeacherTrainingDataInterface {
	return &trainingQuery{
		db: db,
	}
}

// Delete implements teacher.TeacherDataInterface.
func (repo *trainingQuery) Delete(id int, teacherid int) (row int, err error) {
	result := repo.db.Where("teacher_id = ?", teacherid).Delete(&Training{}, id)
	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("failed delete data. Data is not exits row = 0")
	}
	return int(result.RowsAffected), nil
}

// Update implements teacher.TeacherDataInterface.
func (repo *trainingQuery) Update(id int, teacherid int, input training.Core) (data training.Core, err error) {

	tx := repo.db.Model(&Training{}).Where("id = ? AND teacher_id = ?", id, teacherid).Updates(CoreToModel(input))
	if tx.Error != nil {
		return data, tx.Error
	}
	if tx.RowsAffected == 0 {
		return data, errors.New("failed update data, row affected = 0")
	}
	var trainingData Training
	resultFind := repo.db.Find(&trainingData, id)
	if resultFind.Error != nil {
		return training.Core{}, resultFind.Error
	}
	data = ModelToCore(trainingData)
	return data, nil
}

// SelectById implements teacher.TeacherDataInterface.
func (repo *trainingQuery) SelectById(id int, teacherid int) (training.Core, error) {
	var trainingData Training
	tx := repo.db.Where("id = ? and teacher_id = ?", id, teacherid).Preload("Teacher").Preload("Tasks").First(&trainingData) // select * from teachers
	if tx.Error != nil {
		return training.Core{}, tx.Error
	}
	// mapping dari struct gorm model ke struct entities core
	var trainingCore = ModelToCore(trainingData)
	fmt.Println(trainingCore)

	return trainingCore, nil
}

// Insert implements teacher.TeacherDataInterface
func (repo *trainingQuery) Insert(input training.Core) error {

	trainingInputGorm := CoreToModel(input)

	tx := repo.db.Create(&trainingInputGorm) // insert into teachers set name = .....
	if tx.Error != nil {
		return tx.Error
	}

	if tx.RowsAffected == 0 {
		return errors.New("insert failed, row affected = 0")
	}

	return nil
}

// SelectAll implements teacher.TeacherDataInterface
func (repo *trainingQuery) SelectAll(teacherid int) ([]training.Core, error) {
	var trainingData []Training
	tx := repo.db.Where("teacher_id = ?", teacherid).Preload("Teacher").Preload("Tasks").Find(&trainingData) // select * from teachers
	if tx.Error != nil {
		return nil, tx.Error
	}
	// mapping dari struct gorm model ke struct entities core
	var trainingCoreAll []training.Core = ModelToCoreList(trainingData)
	fmt.Println(trainingCoreAll)

	return trainingCoreAll, nil
}
