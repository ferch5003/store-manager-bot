package vertex

import (
	aiplatform "cloud.google.com/go/aiplatform/apiv1"
	"cloud.google.com/go/aiplatform/apiv1/aiplatformpb"
	"context"
	"errors"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/protobuf/types/known/structpb"
	"log"
)

type Client struct {
	projectID  string
	locationID string
	endpointID string

	predictionClient *aiplatform.PredictionClient
}

func NewVertexClient(ctx context.Context, projectID, locationID, endpointID, credentialsPath string) (*Client, error) {
	client, err := aiplatform.NewPredictionClient(ctx, option.WithCredentialsFile(credentialsPath))
	if err != nil {
		log.Fatalf("Failed to create Vertex AI client: %v", err)
	}
	defer func(client *aiplatform.PredictionClient) {
		err := client.Close()
		if err != nil {
			log.Printf("Failed to close Vertex AI client: %v\n", err)
		}
	}(client)

	return &Client{
		projectID:        projectID,
		locationID:       locationID,
		endpointID:       endpointID,
		predictionClient: client,
	}, nil
}

// Response returns the first response for the Vertex API
func (c *Client) Response(text string) (string, error) {
	resp, err := c.makePrediction(text)
	if err != nil {
		return "", err
	}

	if len(resp.Predictions) > 0 {
		// Access the first prediction
		firstPrediction := resp.Predictions[0]

		// Assuming the model returns a struct with a field that contains the text
		if firstPrediction.GetStructValue() != nil {
			fields := firstPrediction.GetStructValue().Fields

			if textValue, ok := fields["text"]; ok {
				return textValue.GetStringValue(), nil
			}
		}
	}

	return "", errors.New("no response")
}

func (c *Client) makePrediction(text string) (*aiplatformpb.PredictResponse, error) {
	ctx := context.Background()

	// Prepare the input
	inputData, err := c.prepareInput(text)
	if err != nil {
		return nil, err
	}

	// Prepare the request
	req := &aiplatformpb.PredictRequest{
		Endpoint:  fmt.Sprintf("projects/%s/locations/%s/endpoints/%s", c.projectID, c.locationID, c.endpointID),
		Instances: []*structpb.Value{inputData},
	}

	// Send the prediction request
	resp, err := c.predictionClient.Predict(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *Client) prepareInput(text string) (*structpb.Value, error) {
	data, err := structpb.NewStruct(map[string]interface{}{
		"instances": []interface{}{
			map[string]interface{}{
				"content": text,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return structpb.NewStructValue(data), nil
}
