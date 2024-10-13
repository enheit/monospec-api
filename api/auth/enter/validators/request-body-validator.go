package validators

import (
	"encoding/json"
	"monospec-api/api/auth/enter/types"
	"monospec-api/shared/problems"

	"github.com/go-playground/validator/v10"
)

func ValidateRequestBody(rawRequestBody string) (*types.RequestBody, error) {
	var requestBody types.RequestBody

	err := json.Unmarshal([]byte(rawRequestBody), &requestBody)

	if err != nil {
		return nil, &problems.Problem{
			Id:             "8e7942fd-a06d-45c8-a496-342a92aadba9",
			Message:        "Invalid request body",
			HttpStatusCode: 400,
		}
	}

	validate := validator.New()

	err = validate.Struct(requestBody)

	if err != nil {
		return nil, &problems.Problem{
			Id:             "40050bcb-6648-4f38-9f7b-a3d27e7b130f",
			Message:        "Invalid request body",
			HttpStatusCode: 400,
		}
	}

	return &requestBody, nil
}
