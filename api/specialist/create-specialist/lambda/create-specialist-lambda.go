package main

import (
	"context"
	"fmt"
	"monospec-api/api/specialist/create-specialist/controller"
	sharedHelpers "monospec-api/shared/helpers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func init() {
	dbPool = sharedHelpers.ConnectToPostgres()
}

func CreateSpecialist(context context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	sessionUserId, ok := request.RequestContext.Authorizer["UserId"].(string)

	if !ok {
		response := &events.APIGatewayProxyResponse{
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       fmt.Sprintf(`{"id": "1f6e4da5-0e7a-4d24-b0aa-45d51d783564", "message": "Failed to get session user id", "httpStatusCode": 500}`),
			StatusCode: 500,
		}

		return response, nil
	}

	controller := controller.New(dbPool, context, sessionUserId)

	rawResponseBody, err := controller.Execute(request.Body)

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
	lambda.Start(CreateSpecialist)
}
