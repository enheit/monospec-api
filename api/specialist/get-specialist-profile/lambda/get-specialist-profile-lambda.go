package main

import (
	"context"
	"fmt"
	"monospec-api/api/specialist/get-specialist-profile/controller"
	sharedHelpers "monospec-api/shared/helpers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jackc/pgx/v5/pgxpool"
)

var dbPool *pgxpool.Pool

func init() {
	dbPool = sharedHelpers.ConnectToPostgres()
}

func GetSpecialistProfile(context context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	sessionUserId, ok := request.RequestContext.Authorizer["UserId"].(int64)

	if !ok {
		response := &events.APIGatewayProxyResponse{
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       fmt.Sprintf(`{"id": "9a869a83-f6e9-4efb-9487-060cf1269f08", "message": "Failed to get session user id", "httpStatusCode": 500}`),
			StatusCode: 500,
		}

		return response, nil
	}

	controller := controller.NewGetSpecialistProfileController(dbPool, context, sessionUserId)

	rawResponseBody, err := controller.Execute(request.PathParameters)

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
	lambda.Start(GetSpecialistProfile)
}
