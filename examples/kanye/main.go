package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type (
	Response struct {
		Type string `json:"response_type"`
		Text string `json:"text"`
	}

	Kanye struct {
		Quote string `json:"quote"`
	}
)

// Get a random Kanye quote
func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	resp, err := http.Get("https://api.kanye.rest/")
	if err != nil {
		return internalError(fmt.Errorf("accessing kanye api: %v", err))
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return internalError(fmt.Errorf("decoding kanye response: %v", err))
	}

	var ye Kanye
	if err := json.Unmarshal(bytes, &ye); err != nil {
		return internalError(fmt.Errorf("error unmarshalling: %s", err))
	}

	// Construct response data
	r := Response{
		Type: "in_channel",
		Text: ye.Quote,
	}

	data, err := json.Marshal(r)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(data),
	}, nil
}

func internalError(err error) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body:       err.Error(),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
