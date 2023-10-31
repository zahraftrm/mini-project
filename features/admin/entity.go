package admin

import "time"

type Core struct {
	Id        uint
	Name      string `validate:"required"`
	Email     string `validate:"required,email"`
	Password  string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AdminDataInterface interface {
	SelectAll() ([]Core, error)
	Insert(input Core) error
	Login(email string, password string) (Core, string, error)
	SelectById(id int) (Core, error)
	Update(id int, input Core) (data Core, err error)
	Delete(id int) (row int, err error)
}

type AdminServiceInterface interface {
	GetAll() ([]Core, error)
	Create(input Core) error
	Login(email string, password string) (Core, string, error)
	GetById(id int) (Core, error)
	Update(id int, input Core) (data Core, err error)
	Delete(id int) (err error)
}
