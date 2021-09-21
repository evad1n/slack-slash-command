package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// https://api.slack.com/interactivity/slash-commands

// Will be POST request

type Response struct {
	Type string `json:"response_type"`
	Text string `json:"text"`
}

// A simple echo command
func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Get slack params
	text := req.QueryStringParameters["text"]

	// Do something

	// Construct response data
	r := Response{
		Type: "in_channel",
		Text: fmt.Sprintf("You said %q", text),
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

func main() {
	lambda.Start(HandleRequest)
}
