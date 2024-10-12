package main

import (
	"context"
	"monospec-api/api/working-day/get-working-day/controller"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func GetWorkingDay(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	controller := controller.Controller{}

	controller.Execute("wokringDayId", "sessionUserId")

	return events.APIGatewayProxyResponse{Body: "It works!", StatusCode: 200}, nil
}

func main() {
	lambda.Start(GetWorkingDay)
}
