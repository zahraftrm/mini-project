package training

import (
	"time"

	"github.com/zahraftrm/mini-project/features/revision"
)

type Core struct {
	Id           uint
	TeacherID    uint   
	AdminID      uint
	Title        string `validate:"required"`
	Expertise    string 
	Category 	 string `validate:"required"`
	Description	 string `validate:"required"`
	Cost 		 string `validate:"required"`
	StartDate   time.Time 
	EndDate     time.Time
	Location	 string
	Status 		 string 
	Teacher      Teacher
	Admin		 Admin
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Revisions    []revision.Core
}

type Teacher struct {
	Id    uint
	Name  string
	Email string
}

type Admin struct {
	Id    uint
	Name  string
	Email string
}

type TrainingDataInterface interface {
	Insert(input Core, role string) error
	SelectAll(teacherid int) ([]Core, error)	
	SelectAllByAdmin() ([]Core, error)											// accessed by admin only
	SelectById(id int, teacherid int) (Core, error)
	SelectByIdTeacher(teacher_id int) ([]Core, error)								// accessed by admin only
	SelectByIdTraining(id int) (Core, error)									// accessed by admin only
	Update(id int, adminid int, role string, input Core) (data Core, email string, name string, err error)
	Delete(id int, teacherid int) (row int, err error)
	DeleteByAdmin(id int) (row int, err error)									// accessed by admin only
}

type TrainingServiceInterface interface {
	Create(input Core, role string) error
	GetAll(teacherid int) ([]Core, error)
	GetAllByAdmin() ([]Core, error)													// accessed by admin only
	GetById(id int, teacherid int) (Core, error)
	GetByIdTeacher(teacher_id int) ([]Core, error)									// accessed by admin only
	GetByIdTraining(id int) (Core, error)											// accessed by admin only
	Update(id int, adminid int, role string, input Core) (data Core, email string, name string, err error)
	Delete(id int, teacherid int) (err error)
	DeleteByAdmin(id int) (err error)												// accessed by admin only
}
