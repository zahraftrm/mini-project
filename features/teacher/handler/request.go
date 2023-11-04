package handler

import "github.com/zahraftrm/mini-project/features/teacher"

type TeacherRequest struct {
	NUPTK    	string `json:"nuptk" form:"nuptk"`
	Name     	string `json:"name" form:"name"`
	Phone    	string `json:"phone" form:"phone"`
	Expertise   string `json:"expertise" form:"expertise"`
	Email    	string `json:"email" form:"email"`
	Password 	string `json:"password" form:"password"`
}

type AuthRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func RequestToCore(dataRequest TeacherRequest) teacher.Core {
	return teacher.Core{
		NUPTK:    	dataRequest.NUPTK,
		Name:     	dataRequest.Name,
		Phone:    	dataRequest.Phone,
		Expertise:  dataRequest.Expertise,
		Email:    	dataRequest.Email,
		Password: 	dataRequest.Password,
	}
}
