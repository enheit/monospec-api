package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func GetUserAppointments(context context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		Body:       "It works!",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(GetUserAppointments)
}
