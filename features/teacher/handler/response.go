package handler

import (
	"time"

	"github.com/zahraftrm/mini-project/features/teacher"
)

type TeacherResponse struct {
	Id        uint      `json:"id"`
	NUPTK     string    `json:"nuptk"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone,omitempty"`
	Expertise string    `json:"expertise"`
	CreatedAt time.Time `json:"created_at"`
}

func CoreToResponse(dataCore teacher.Core) TeacherResponse {
	return TeacherResponse{
		Id:        dataCore.Id,
		NUPTK:     dataCore.NUPTK,
		Name:      dataCore.Name,
		Phone:     dataCore.Phone,
		Expertise: dataCore.Expertise,
		Email:     dataCore.Email,
		CreatedAt: dataCore.CreatedAt,
	}
}

func CoreToResponseList(dataCore []teacher.Core) []TeacherResponse {
	var result []TeacherResponse
	for _, v := range dataCore {
		result = append(result, CoreToResponse(v))
	}
	return result
}
