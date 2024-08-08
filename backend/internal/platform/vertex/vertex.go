package vertex

import (
	"backend/internal/domain"
	"backend/internal/platform/files"
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

type Config struct {
	ProjectTunedID  string
	RegionTuned     string
	EndpointTunedID string

	ProjectFlashID string
	RegionFlash    string
	ModelFlashID   string
}

const (
	_endpointURL    = "https://%s-aiplatform.googleapis.com/v1/projects/%s/locations/%s/endpoints/%s:streamGenerateContent"
	_geminiFlashURL = "https://%s-aiplatform.googleapis.com/v1/projects/%s/locations/%s/publishers/google/models/%s:streamGenerateContent"
)

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
	Text       string `json:"text,omitempty"`
	InlineData *Blob  `json:"inlineData,omitempty"`
}

type Blob struct {
	MimeType string `json:"mimeType"`
	Data     string `json:"data"`
}

type GeneratedResponse struct {
	Candidates []Candidate `json:"candidates"`
}

type Candidate struct {
	Content *Content `json:"content"`
}

type Client struct {
	ctx context.Context

	tunedURL string
	flashURL string

	vertexClient         *http.Client
	sessionTunedContents []Content
	sessionFlashContents []Content
}

func NewVertexClient(
	ctx context.Context,
	config Config,
	credentialsPath string,
	histories []domain.History) (*Client, error) {
	client, err := generateVertexClient(ctx, credentialsPath)
	if err != nil {
		return nil, err
	}

	tunedContents, flashContents, err := serializeHistories(histories)
	if err != nil {
		return nil, err
	}

	tunedURL := fmt.Sprintf(
		_endpointURL, config.RegionTuned, config.ProjectTunedID, config.RegionTuned, config.EndpointTunedID)
	flashURL := fmt.Sprintf(
		_geminiFlashURL, config.RegionFlash, config.ProjectFlashID, config.RegionFlash, config.ModelFlashID)

	return &Client{
		tunedURL: tunedURL,
		flashURL: flashURL,

		vertexClient:         client,
		sessionTunedContents: tunedContents,
		sessionFlashContents: flashContents,
	}, nil
}

// SendMessage returns the first response from the Vertex API.
func (c *Client) SendMessage(data string) (string, error) {
	// Add user prompt to contents history.
	content, err := getContent(User, data)
	if err != nil {
		return "", err
	}

	var contents []Content
	_, isMIMEContent := files.GetMIMEType(data)
	if isMIMEContent {
		contents = append(c.sessionFlashContents, content)
	} else {
		contents = append(c.sessionTunedContents, content)
	}

	requestBody := map[string]interface{}{
		"contents": contents,
		"generationConfig": map[string]interface{}{
			"maxOutputTokens": 8192,
			"temperature":     1,
			"topP":            0.95,
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
	url := c.tunedURL
	if isMIMEContent {
		url = c.flashURL
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(jsonData)))
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

	log.Printf("VERTEX RESPONSE: %s", body)

	var response []GeneratedResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	botResponse, err := processResponse(response)
	if err != nil {
		return "", err
	}

	// Add model response to contents history.
	botContent, err := getContent(Model, botResponse)
	if err != nil {
		return "", err
	}

	if len(botContent.Parts) <= 0 {
		return "", errors.New("no response content from bot")
	}

	// If there is positive response then add content user to histories.
	if !isMIMEContent {
		c.sessionTunedContents = append(c.sessionTunedContents, content, botContent)
	}

	c.sessionFlashContents = append(c.sessionFlashContents, content, botContent)

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

func getContent(role Role, data string) (Content, error) {
	if len(data) == 0 {
		return Content{}, errors.New("data is empty")
	}

	// Check if data is different from audio or image.
	var parts []Part
	mimeType, ok := files.GetMIMEType(data)
	if ok {
		base64Data, err := files.GetBase64Image(data, mimeType)
		if err != nil {
			return Content{}, err
		}

		parts = []Part{
			{
				InlineData: &Blob{
					MimeType: mimeType,
					Data:     base64Data,
				},
			},
			{
				Text: "Si esta imagen es algún producto, me puedes decir cual es y algún " +
					"de MercadoLibre Colombia link para buscarlo",
			},
		}
	} else {
		parts = []Part{
			{
				Text: data,
			},
		}
	}

	userContent := Content{
		Role:  role,
		Parts: parts,
	}

	return userContent, nil
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

func serializeHistories(histories []domain.History) (tunedContents []Content, flashContents []Content, err error) {
	for _, history := range histories {
		userContent, err := getContent(User, history.UserMessage)
		if err != nil {
			return nil, nil, err
		}

		modelContent, err := getContent(Model, history.BotResponse)
		if err != nil {
			return nil, nil, err
		}

		flashContents = append(flashContents, userContent, modelContent)

		// Tuned models (Using gemini-1.0-pro-002) doesn't support multimedia.
		if !history.Multimedia {
			tunedContents = append(tunedContents, userContent, modelContent)
		}
	}

	return tunedContents, flashContents, nil
}
