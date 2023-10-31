package revision

import "time"

type Core struct {
	Id         uint
	TrainingID uint
	Message   string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RevisionDataInterface interface {
	Insert(input Core) error
	Update(id int, adminid int, input Core) (data Core, err error)
	Delete(id int, adminid int) (row int, err error)
}

type RevisionServiceInterface interface {
	Create(input Core) error
	Update(id int, adminid int, input Core) (data Core, err error)
	Delete(id int, adminid int) (err error)
}
