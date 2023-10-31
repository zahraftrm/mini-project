package database

import (
	"fmt"

	"github.com/zahraftrm/mini-project/app/config"
	_adminData "github.com/zahraftrm/mini-project/features/admin/data"

	//_trainingData "github.com/zahraftrm/mini-project/features/teacher_training/data"
	_revisionData "github.com/zahraftrm/mini-project/features/revision/data"
	_teacherData "github.com/zahraftrm/mini-project/features/teacher/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// database connection
func InitDBMysql(cfg *config.AppConfig) *gorm.DB {

	// declare struct config & variable connectionString
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}

func InitialMigration(db *gorm.DB) {
	db.AutoMigrate(&_teacherData.Teacher{})
	db.AutoMigrate(&_adminData.Admin{})
	//db.AutoMigrate(&_trainingData.Training{})
	db.AutoMigrate(&_revisionData.Revision{})
}
