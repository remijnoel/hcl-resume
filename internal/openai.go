package internal

import (
	"context"
	"os"

	"github.com/invopop/jsonschema"
	openai "github.com/openai/openai-go"
	"github.com/openai/openai-go/option"

	log "github.com/sirupsen/logrus"
)

func NewOpenAIClient() *openai.Client {
	client := openai.NewClient(
		option.WithAPIKey(os.Getenv("OPENAI_API_KEY")),
	)
	return &client
}

func GenerateSchema[T any]() interface{} {
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	var v T
	return reflector.Reflect(v)
}

type CompletionRequest struct {
	Prompt         string         `json:"prompt"`
	SystemPrompt   string         `json:"system_prompt"`
	Model          string         `json:"model"`
	Client         *openai.Client `json:"-"`
	ResponseSchema any            `json:"-"`
}

func RequestCompletion(r CompletionRequest) (string, error) {
	ctx := context.Background()
	messages := []openai.ChatCompletionMessageParamUnion{}

	if r.SystemPrompt != "" {
		messages = append(messages, openai.SystemMessage(r.SystemPrompt))
	}
	messages = append(messages, openai.UserMessage(r.Prompt))

	params := openai.ChatCompletionNewParams{
		Messages: messages,
		Model:    r.Model,
	}
	if r.ResponseSchema != nil {
		params.ResponseFormat = openai.ChatCompletionNewParamsResponseFormatUnion{
			OfJSONSchema: &openai.ResponseFormatJSONSchemaParam{
				JSONSchema: openai.ResponseFormatJSONSchemaJSONSchemaParam{
					Name:   "structured_output",
					Schema: r.ResponseSchema,
					Strict: openai.Bool(true),
				},
			},
		}
	}

	resp, err := r.Client.Chat.Completions.New(ctx, params)
	if err != nil {
		return "", err
	}
	log.Debugf("OpenAI response: %s", resp.Choices[0].Message.Content)

	if r.ResponseSchema != nil {

	}

	// Return raw JSON, user can unmarshal as needed
	return resp.Choices[0].Message.Content, nil
}
