package main

import (
	b64 "encoding/base64"
	"errors"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler handles basic auth requests
func Handler(request events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {
	token := request.AuthorizationToken
	if len(token) == 0 {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	tokenSlice := strings.Split(token, " ")
	if len(tokenSlice) != 2 {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	basicString, _ := b64.StdEncoding.DecodeString(tokenSlice[1])
	basic := strings.Split(string(basicString), ":")
	if len(basic) != 2 {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	if len(basic[0]) <= 3 || basic[0] != os.Getenv("USERNAME") || basic[1] != os.Getenv("PASSWORD") {
		return events.APIGatewayCustomAuthorizerResponse{}, errors.New("Unauthorized")
	}

	return generatePolicy(basic[0], "Allow", request.MethodArn), nil
}

func main() {
	lambda.Start(Handler)
}

func generatePolicy(principalID string, effect string, resource string) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{
		PrincipalID: principalID,
		Context: map[string]interface{}{
			"username": principalID, // Username is already on as principal ID but this could be used for roles etc.
		},
	}

	if effect != "" && resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
	}
	return authResponse
}
