package validators

import (
	"monospec-api/api/specialist/get-specialist-profile/types"
	"monospec-api/shared/problems"

	"github.com/go-playground/validator/v10"
)

func ValidatePathParams(rawPathParams map[string]string) (*types.PathParams, error) {
	pathParams := &types.PathParams{}

	validate := validator.New()

	err := validate.Struct(pathParams)

	if err != nil {
		return nil, &problems.Problem{
			Id:             "9367b91f-39aa-40f6-bf0e-42c173094b71",
			Message:        "Invalid path params",
			HttpStatusCode: 400,
		}
	}

	return pathParams, nil
}
