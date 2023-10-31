package handler

import (
	"time"

	"github.com/zahraftrm/mini-project/features/admin"
)

type AdminResponse struct {
	Id        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func CoreToResponse(dataCore admin.Core) AdminResponse {
	return AdminResponse{
		Id:        dataCore.Id,
		Name:      dataCore.Name,
		Email:     dataCore.Email,
		CreatedAt: dataCore.CreatedAt,
	}
}

func CoreToResponseList(dataCore []admin.Core) []AdminResponse {
	var result []AdminResponse
	for _, v := range dataCore {
		result = append(result, CoreToResponse(v))
	}
	return result
}
