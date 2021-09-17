package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Response struct {
	Definitions []string `json:"definitions"`
}

// Called like <API_ENDPOINT>?word=<WORD>

func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	word := req.QueryStringParameters["word"]

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: "hello",
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
