package handler

type RecommendationResponse struct {
	Status         string `json:"status"`
	Recommendation string `json:"recommendation"`
}