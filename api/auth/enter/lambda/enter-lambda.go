package main

import (
	"context"

	"monospec-api/api/auth/enter/controller"
	"monospec-api/shared/helpers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Enter(context context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	controller := controller.EnterController{}

	rawResponseBody, err := controller.Execute(request.Body)

	if err != nil {
		return helpers.TransformErrorToHttpResponse(err), err
	}

	return &events.APIGatewayProxyResponse{
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       *rawResponseBody,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Enter)
}
