package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"monospec-api/repos"
)

func GetMe(context context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Print(context)
	fmt.Print(request)

	userRepo := repos.NewUserRepo()

	fmt.Print(userRepo.GetUser())

	return events.APIGatewayProxyResponse{Body: "It works!", StatusCode: 200}, nil
}

func main() {
	lambda.Start(GetMe)
}
