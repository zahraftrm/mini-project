package teacher_training

import (
	"time"

	"github.com/zahraftrm/mini-project/features/revision"
)

type Core struct {
	Id           uint
	TeacherID       uint   `validate:"required"`
	Title        string `validate:"required"`
	Expertise    string `validate:"required"`
	Category 	 string `validate:"required"`
	Description	 string `validate:"required"`
	Cost 		 string `validate:"required"`
	StartDate   time.Time `validate:"required"`
	EndDate     time.Time `validate:"required"`
	Location	 string `validate:"required"`
	Status 		 string 
	Teacher      Teacher
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Revisions    []revision.Core
}

type Teacher struct {
	Id    uint
	Name  string
	Email string
}

type TeacherTrainingDataInterface interface {
	Insert(input Core) error
	SelectAll(teacherid int) ([]Core, error)
	SelectById(id int, teacherid int) (Core, error)
	Update(id int, teacherid int, input Core) (data Core, err error)
	Delete(id int, teacherid int) (row int, err error)
}

type TeacherTrainingServiceInterface interface {
	Create(input Core) error
	GetAll(teacherid int) ([]Core, error)
	GetById(id int, teacherid int) (Core, error)
	Update(id int, teacherid int, input Core) (data Core, err error)
	Delete(id int, teacherid int) (err error)
}
