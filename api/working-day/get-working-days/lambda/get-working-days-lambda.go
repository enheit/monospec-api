package main

import (
	"context"
	"monospec-api/api/working-day/get-working-days/controller"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func GetWorkingDays(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	controller := controller.GetWorkingDaysController{}

	controller.Execute("User Id goes here")

	return events.APIGatewayProxyResponse{Body: "It works!", StatusCode: 200}, nil
}

func main() {
	lambda.Start(GetWorkingDays)
}
