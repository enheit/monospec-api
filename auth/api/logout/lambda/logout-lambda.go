package main

import (
	"context"
	"monospec-api/auth/api/logout/controller"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Logout(context context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	controller := controller.New()

	controller.Execute()

	return &events.APIGatewayProxyResponse{
		Body:       "It works!",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Logout)
}
