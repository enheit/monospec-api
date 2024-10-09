package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func GetUserAppointments(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Print(context)
	fmt.Print(request)

	return events.APIGatewayProxyResponse{Body: "Hello, CDK 2!", StatusCode: 200}, nil
}

func main() {
	lambda.Start(GetUserAppointments)
}
