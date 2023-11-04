
package service

import (
	"context"

	openai "github.com/sashabaranov/go-openai"
	"github.com/zahraftrm/mini-project/features/recommendation"
)

type RecommendationService interface {
    Recommendation(app *recommendation.Recommendation) (string, error)
}

type recommendationService struct{}

func NewRecommendationService() RecommendationService {
    return &recommendationService{}
}

func (s *recommendationService) Recommendation(app *recommendation.Recommendation) (string, error) {
    ctx := context.Background()
    client := openai.NewClient(app.OpenAIKey)
    model := openai.GPT3Dot5Turbo
    messages := []openai.ChatCompletionMessage{
        {
            Role:    openai.ChatMessageRoleSystem,
            Content: "Halo, saya adalah sistem untuk memberikan rekomendasi pelatihan untuk guru yang sesuai",
        },
		{
            Role:    openai.ChatMessageRoleUser,
            Content: app.Template,
        },
    }

    resp, err := s.getCompletionFromMessages(ctx, client, messages, model)
    if err != nil {
        return "", err
    }
    answer := resp.Choices[0].Message.Content
    return answer, nil
}

func (s *recommendationService) getCompletionFromMessages(
    ctx context.Context,
    client *openai.Client,
    messages []openai.ChatCompletionMessage,
    model string,
) (openai.ChatCompletionResponse, error) {
    if model == "" {
        model = openai.GPT3Dot5Turbo
    }

    resp, err := client.CreateChatCompletion(
        ctx,
        openai.ChatCompletionRequest{
            Model:    model,
            Messages: messages,
        },
    )
    return resp, err
}
