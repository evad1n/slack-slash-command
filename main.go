package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Will be POST request

// URL encoded data will be in the BODY. Not as query parameterss.

type Response struct {
	Type string `json:"response_type"`
	Text string `json:"text"`
}

// A simple echo command
func HandleRequest(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// Slack sends its parameters as url encoded data in the request body. These need to be parsed to obtain the key/values. A list of the data slack sends can be seen [here](https://api.slack.com/interactivity/slash-commands).

	// Get slack params
	params, err := url.ParseQuery(req.Body)
	if err != nil {
		return internalError(fmt.Errorf("decoding slack params: %v", err))
	}
	text := params.Get("text")

	// Do something. Anything you want really
	// Some cool code

	// Construct response data
	r := Response{
		Type: "in_channel",
		Text: fmt.Sprintf("You said '%s'", text),
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
