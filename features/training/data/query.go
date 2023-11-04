package data

import (
	"errors"
	"fmt"

	training "github.com/zahraftrm/mini-project/features/training"
	"gorm.io/gorm"
)

type trainingQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) training.TrainingDataInterface {
	return &trainingQuery{
		db: db,
	}
}

// Insert implements teacher.TeacherDataInterface
func (repo *trainingQuery) Insert(input training.Core, role string) error {
	var expertise string
	result := repo.db.Model(&Training{}).
		Joins("JOIN teachers ON trainings.teacher_id = teachers.id").
		Where("trainings.teacher_id = ?", input.TeacherID).
		Pluck("teachers.Expertise", &expertise)
	if result.Error != nil {
		return result.Error
	}

	input.Expertise = expertise

	trainingInputGorm := CoreToModel(input, role)

	tx := repo.db.Create(&trainingInputGorm) 
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
	tx := repo.db.Where("teacher_id = ?", teacherid).Preload("Teacher").Preload("Admin").Preload("Revisions").Find(&trainingData) // select * from teachers
	if tx.Error != nil {
		return nil, tx.Error
	}

	// mapping dari struct gorm model ke struct entities core
	var trainingCoreAll []training.Core = ModelToCoreList(trainingData)
	fmt.Println(trainingCoreAll)

	return trainingCoreAll, nil
}

// SelectAllByAdmin
func (repo *trainingQuery) SelectAllByAdmin() ([]training.Core, error) {
    var trainingData []Training
    tx := repo.db.Preload("Teacher").Preload("Admin").Preload("Revisions").Find(&trainingData) // Select * from training
    if tx.Error != nil {
        return nil, tx.Error
    }

    // Mapping dari struct GORM model ke struct entities core
    var trainingCoreAll []training.Core = ModelToCoreList(trainingData)
    fmt.Println(trainingCoreAll)

    return trainingCoreAll, nil
}


// SelectById implements teacher.TeacherDataInterface.
func (repo *trainingQuery) SelectById(id int, teacherid int) (training.Core, error) {
	var trainingData Training
	tx := repo.db.Where("id = ? and teacher_id = ?", id, teacherid).Preload("Teacher").Preload("Admin").Preload("Revisions").First(&trainingData) // select * from teachers
	if tx.Error != nil {
		return training.Core{}, tx.Error
	}
	// mapping dari struct gorm model ke struct entities core
	var trainingCore = ModelToCore(trainingData)
	fmt.Println(trainingCore)

	return trainingCore, nil
}

func (repo *trainingQuery) SelectByIdTeacher(teacher_id int) ([]training.Core, error) {
	var trainingData []Training
	tx := repo.db.Where("teacher_id = ?", teacher_id).Preload("Teacher").Preload("Admin").Preload("Revisions").Find(&trainingData)
	if tx.Error != nil {
		return nil, tx.Error
	}

	// mapping dari struct gorm model ke struct entities core
	var trainingCoreAll []training.Core = ModelToCoreList(trainingData)
    fmt.Println(trainingCoreAll)

	return trainingCoreAll, nil
}

func (repo *trainingQuery) SelectByIdTraining(id int) (training.Core, error) {
	var trainingData Training
	tx := repo.db.Where("id = ?", id).First(&trainingData)
	if tx.Error != nil {
		return training.Core{}, tx.Error
	}
	// mapping dari struct gorm model ke struct entities core
	var trainingCore = ModelToCore(trainingData)
	fmt.Println(trainingCore)

	return trainingCore, nil
}

// Update implements teacher.TeacherDataInterface.
func (repo *trainingQuery) Update(id int, adminid int, role string, input training.Core) (data training.Core, email string, name string, err error) {
	
	tx := repo.db.Model(&Training{}).Where("id = ?", id).Updates(CoreToModel(input, role))
	if tx.Error != nil {
		return data, "", "", tx.Error
	}
	if tx.RowsAffected == 0 {
		return data, "", "", errors.New("failed update data, row affected = 0")
	}
	var trainingData Training
	resultFind := repo.db.Find(&trainingData, id)
	if resultFind.Error != nil {
		return training.Core{}, "", "", resultFind.Error
	}
	data = ModelToCore(trainingData)

	var mail, nama string
	result := repo.db.Model(&Training{}).
		Joins("JOIN teachers ON trainings.teacher_id = teachers.id").
		Where("trainings.teacher_id = ?", data.TeacherID).
		Pluck("teachers.email", &mail)
	if result.Error != nil {
		return data, "yesh", "", result.Error
	}
	result = repo.db.Model(&Training{}).
		Joins("JOIN teachers ON trainings.teacher_id = teachers.id").
		Where("trainings.teacher_id = ?", data.TeacherID).
		Pluck("teachers.name", &nama)
	if result.Error != nil {
		return data, "", "lol", result.Error
	}
	email = mail
	name = nama

	return data, email, name, nil
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

func (repo *trainingQuery) DeleteByAdmin(id int) (row int, err error) {
	result := repo.db.Where("id = ?", id).Delete(&Training{})
	if result.Error != nil {
		return -1, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("failed delete data. Data is not exits row = 0")
	}
	return int(result.RowsAffected), nil
}