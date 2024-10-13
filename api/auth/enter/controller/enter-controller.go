package controller

import (
	"monospec-api/api/auth/enter/enums"
	"monospec-api/api/auth/enter/helpers"
	"monospec-api/api/auth/enter/services"
	"monospec-api/api/auth/enter/use-case"
	"monospec-api/api/auth/enter/validators"
	"os"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

type EnterController struct {
}

func (c *EnterController) Execute(rawRequestBody string) (*string, error) {
	requestBody, err := validators.ValidateRequestBody(rawRequestBody)

	if err != nil {
		return nil, err
	}

	cognitoAuthService := &services.CognitoAuthService{
		CognitoClient: cognitoidentityprovider.New(cognitoidentityprovider.Options{}),
		ClientId:      os.Getenv(enums.LambdaEnvsUserPoolClientId),
	}

	useCase := &usecase.EnterUseCase{
		CognitoAuthService: cognitoAuthService,
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
