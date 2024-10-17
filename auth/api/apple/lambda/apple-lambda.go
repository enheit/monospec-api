package main

import (
	"context"

	"monospec-api/auth/api/apple/controller"
	sharedHelpers "monospec-api/shared/helpers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func init() {
	println("Connecting to Postgres...")

	dbPool = sharedHelpers.ConnectToPostgres()

	println("Postgres connection is established")
}

func Apple(context context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	println(request.Body)
	println(request.Headers)

	controller := controller.New(dbPool, context)

	println("Controller is created")
	rawResponseBody, err := controller.Execute(request.Body)
	println("Controller is executed")

	if err != nil {
		return sharedHelpers.TransformErrorToHttpResponse(err), nil
	}

	response := &events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       *rawResponseBody,
		StatusCode: 200,
	}

	return response, nil
}

func main() {
	lambda.Start(Apple)
}
