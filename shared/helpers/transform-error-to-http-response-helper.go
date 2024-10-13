package helpers

import (
	"encoding/json"
	"fmt"
	"monospec-api/shared/problems"

	"github.com/aws/aws-lambda-go/events"
)

func TransformErrorToHttpResponse(err error) *events.APIGatewayProxyResponse {
	switch e := err.(type) {
	case *problems.Problem:
		responseBodyBytes, err := json.Marshal(e)

		if err != nil {
			return &events.APIGatewayProxyResponse{
				Headers:    map[string]string{"Content-Type": "application/json"},
				Body:       fmt.Sprintf(`{"id": "eb30e676-6eff-4775-bd92-f3a8aa09e801", "message": "Failed to transform Problem to http response", "description": "%s", "httpStatusCode": 500}"`, err.Error()),
				StatusCode: 500,
			}
		}

		responseBodyString := string(responseBodyBytes)

		return &events.APIGatewayProxyResponse{
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       responseBodyString,
			StatusCode: e.HttpStatusCode,
		}
	default:
		return &events.APIGatewayProxyResponse{
			Headers:    map[string]string{"Content-Type": "application/json"},
			Body:       fmt.Sprintf(`{"id": "fce417ab-ce83-4b8a-88d7-fee9788aed42", "message": "Unmapped error", "description": "%s", "httpStatusCode": 500}`, err.Error()),
			StatusCode: 500,
		}
	}
}
