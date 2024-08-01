package vertex

import (
	"backend/internal/domain"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/oauth2/google"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const _endpointURL = "https://%s-aiplatform.googleapis.com/v1/projects/%s/locations/%s/endpoints/%s:streamGenerateContent"

type Role string

const (
	User  Role = "user"
	Model Role = "model"
)

type Content struct {
	Role  Role   `json:"role"`
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text"`
}

type GeneratedResponse struct {
	Candidates []Candidate `json:"candidates"`
}

type Candidate struct {
	Content *Content `json:"content"`
}

type Client struct {
	ctx context.Context

	url string

	vertexClient    *http.Client
	sessionContents []Content
}

func NewVertexClient(
	ctx context.Context,
	projectID,
	region,
	endpointID,
	credentialsPath string,
	histories []domain.History) (*Client, error) {
	url := fmt.Sprintf(_endpointURL, region, projectID, region, endpointID)
	client, err := generateVertexClient(ctx, credentialsPath)
	if err != nil {
		return nil, err
	}

	contents := serializeHistories(histories)

	return &Client{
		url: url,

		vertexClient:    client,
		sessionContents: contents,
	}, nil
}

// SendMessage returns the first response from the Vertex API
func (c *Client) SendMessage(msg string) (string, error) {
	// Add user prompt to contents history.
	c.sessionContents = appendContent(c.sessionContents, User, msg)

	requestBody := map[string]interface{}{
		"contents": c.sessionContents,
		"generationConfig": map[string]interface{}{
			"maxOutputTokens": 2048,
			"temperature":     1,
			"topP":            1,
		},
		"safetySettings": []map[string]interface{}{
			{
				"category":  "HARM_CATEGORY_HATE_SPEECH",
				"threshold": "BLOCK_MEDIUM_AND_ABOVE",
			},
			{
				"category":  "HARM_CATEGORY_DANGEROUS_CONTENT",
				"threshold": "BLOCK_MEDIUM_AND_ABOVE",
			},
			{
				"category":  "HARM_CATEGORY_SEXUALLY_EXPLICIT",
				"threshold": "BLOCK_MEDIUM_AND_ABOVE",
			},
			{
				"category":  "HARM_CATEGORY_HARASSMENT",
				"threshold": "BLOCK_MEDIUM_AND_ABOVE",
			},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	// Make the API request
	req, err := http.NewRequest("POST", c.url, strings.NewReader(string(jsonData)))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.vertexClient.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	// Read the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response []GeneratedResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	botResponse, err := processResponse(response)
	if err != nil {
		return "", err
	}

	// Add model response to contents history.
	c.sessionContents = appendContent(c.sessionContents, Model, botResponse)

	return botResponse, nil
}

func generateVertexClient(ctx context.Context, credentialsPath string) (*http.Client, error) {
	jsonCredentials, err := os.ReadFile(credentialsPath)
	if err != nil {
		return nil, err
	}

	// Get the JWT config for OAuth 2.0 access token
	config, err := google.JWTConfigFromJSON(jsonCredentials, "https://www.googleapis.com/auth/cloud-platform")
	if err != nil {
		return nil, err
	}

	return config.Client(ctx), nil
}

func appendContent(contents []Content, role Role, msg string) []Content {
	userContent := Content{
		Role:  role,
		Parts: []Part{{Text: msg}},
	}
	contents = append(contents, userContent)

	return contents
}

func processResponse(generatedResponses []GeneratedResponse) (string, error) {
	if generatedResponses == nil {
		return "", errors.New("no candidates")
	}

	var botResponse string
	for _, generatedResponse := range generatedResponses {
		for _, candidate := range generatedResponse.Candidates {
			if candidate.Content != nil && len(candidate.Content.Parts) > 0 {
				botResponse = fmt.Sprintf("%s%s", botResponse, candidate.Content.Parts[0].Text)
			}
		}
	}

	return botResponse, nil
}

func serializeHistories(histories []domain.History) []Content {
	var contents []Content
	for _, history := range histories {
		userContent := Content{
			Role:  User,
			Parts: []Part{{Text: history.UserMessage}},
		}
		contents = append(contents, userContent)

		modelContent := Content{
			Role:  Model,
			Parts: []Part{{Text: history.BotResponse}},
		}
		contents = append(contents, modelContent)
	}

	return contents
}
