package main

import (
	"context"

	"monospec-api/api/auth/logout/controller"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Logout(context context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	controller := controller.LogoutController{}

	controller.Execute()

	return &events.APIGatewayProxyResponse{
		Body:       "It works!",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Logout)
}
