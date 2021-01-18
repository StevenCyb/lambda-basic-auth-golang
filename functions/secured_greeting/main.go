package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler is a lambda handler invoked by the `lambda.Start` function call
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// From ...
	username := fmt.Sprintf("%v", request.RequestContext.Authorizer["username"])

	body, _ := json.Marshal(map[string]string{
		"message":       "Hello from secured lambda function",
		"your_username": username,
	})

	var buf bytes.Buffer
	json.HTMLEscape(&buf, body)

	return events.APIGatewayProxyResponse{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}, nil
}

func main() {
	lambda.Start(Handler)
}
