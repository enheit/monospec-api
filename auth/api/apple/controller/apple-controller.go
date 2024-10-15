package controller

import (
	"monospec-api/api/auth/apple/enums"
	"monospec-api/api/auth/apple/helpers"
	"monospec-api/api/auth/apple/services"
	"monospec-api/api/auth/apple/use-case"
	"monospec-api/api/auth/apple/validators"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type AppleController struct {
}

func (c *AppleController) Execute(rawRequestBody string) (*string, error) {
	requestBody, err := validators.ValidateRequestBody(rawRequestBody)

	if err != nil {
		return nil, err
	}

	cognitoAuthService := &services.CognitoAuthService{
		CognitoClient: cognitoidentityprovider.New(cognitoidentityprovider.Options{}),
		ClientId:      os.Getenv(enums.LambdaEnvsUserPoolClientId),
	}

	useCase := &usecase.AppleUseCase{
		TokenService: cognitoAuthService,
	}

	user, err := useCase.Enter(requestBody.Token)

	if err != nil {
		return nil, err
	}

	rawResponseBody, err := helpers.PrepareResponseBody(user)

	if err != nil {
		return nil, err
	}

	return rawResponseBody, nil
}
