package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/zahraftrm/mini-project/features/recommendation"
	service "github.com/zahraftrm/mini-project/features/recommendation/service"
)

type RecommendationHandler struct {
    RecommendationUsecase service.RecommendationService
}

func NewRecommendationHandler(service service.RecommendationService) *RecommendationHandler {
    return &RecommendationHandler{
        RecommendationUsecase: service,
    }
}

func (h *RecommendationHandler) Recommendation(c echo.Context) error {
    var requestData map[string]interface{}
    err := c.Bind(&requestData)
    if err != nil {
        return c.JSON(http.StatusBadRequest, "Invalid JSON format")
    }

    expertise, checkExpertise := requestData["expertise"].(string)
    educationLevel, checkEducationLevel := requestData["education_level"].(string)
    experience, checkExperience := requestData["experience"].(string)
    purpose, checkPurpose := requestData["purpose"].(string) 

    if !checkExpertise || !checkEducationLevel || !checkExperience || !checkPurpose {
        return c.JSON(http.StatusBadRequest, "Invalid request format")
    }

    applicationDetails := fmt.Sprintf("Rekomendasi pelatihan untuk guru yang cocok dengan bidang %s, pendidikan terakhir %s, lama mengajar %s, dan tujuan mengikuti pelatihan untuk %s.", expertise, educationLevel, experience, purpose)

    app := &recommendation.Recommendation{
        Template: applicationDetails,
        OpenAIKey: os.Getenv("OPENAI_API_KEY"), // Mengambil OpenAIKey dari environment
    }

    answer, err := h.RecommendationUsecase.Recommendation(app)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, "Error dalam pengajuan pendaftaran magang")
    }

    responseData := RecommendationResponse{
        Status:        "success",
        Recommendation: answer,
    }

    return c.JSON(http.StatusOK, responseData)
}
