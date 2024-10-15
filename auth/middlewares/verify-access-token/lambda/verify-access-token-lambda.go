package main

import (
	"context"
	"monospec-api/auth/middlewares/verify-access-token/controller"
	"monospec-api/auth/middlewares/verify-access-token/errors"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"monospec-api/auth/types"
)

func VerifyAccessToken(context context.Context, request events.APIGatewayCustomAuthorizerRequest) (*events.APIGatewayCustomAuthorizerResponse, error) {
	rawAccessToken := request.AuthorizationToken

	controller := controller.VerifyAccessTokenController{}

	accessToken, err := controller.Execute(rawAccessToken)

	if err != nil {
		return nil, &errors.Unathorized{}
	}

	return allow(accessToken), nil
}

// HINT: See https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-lambda-authorizer-output.html @ Roman
func allow(accessToken *types.AccessToken) *events.APIGatewayCustomAuthorizerResponse {
	return &events.APIGatewayCustomAuthorizerResponse{
		PrincipalID: accessToken.Payload.Subject,
		PolicyDocument: events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   "Allow",
					Resource: []string{"*"},
				},
			},
		},
		Context: map[string]interface{}{
			"UserId": accessToken.Payload.Subject,
		},
	}
}

func main() {
	lambda.Start(VerifyAccessToken)
}
