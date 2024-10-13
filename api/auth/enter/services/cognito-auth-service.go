package services

import (
	"context"
	customTypes "monospec-api/api/auth/enter/types"
	"monospec-api/shared/helpers"
	"monospec-api/shared/problems"

	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
)

type CognitoAuthService struct {
	CognitoClient *cognitoidentityprovider.Client
	ClientId      string
}

func (c *CognitoAuthService) InitAuth(appleUserId string) (*string, error) {
	input := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow: types.AuthFlowTypeCustomAuth,
		ClientId: &c.ClientId,
		AuthParameters: map[string]string{
			"USERNAME": appleUserId,
		},
	}

	result, err := c.CognitoClient.InitiateAuth(context.Background(), input)

	if err != nil {
		originalErrorMessage := err.Error()

		problem := problems.Problem{
			Id:             "1742599d-c49b-4db1-a05c-322eb86bfef3",
			Message:        "Failed to initiate auth",
			Description:    &originalErrorMessage,
			HttpStatusCode: 500,
		}

		return nil, problem
	}

	return result.Session, nil
}

func (c *CognitoAuthService) RespondToAuthChallenge(appleUserId string, appleIdToken string, session *string) (*customTypes.TokensBundle, error) {
	input := &cognitoidentityprovider.RespondToAuthChallengeInput{
		ChallengeName: types.ChallengeNameTypeCustomChallenge,
		ClientId:      &c.ClientId,
		Session:       session,
		ChallengeResponses: map[string]string{
			"USERNAME": appleUserId,
			"ANSWER":   appleIdToken,
		},
	}

	output, err := c.CognitoClient.RespondToAuthChallenge(context.Background(), input)

	if err != nil {
		originalErrorMessage := err.Error()

		problem := problems.Problem{
			Id:             "2eed86e4-a75a-44a7-afd3-0549a66ec665",
			Message:        "Failed to respond to auth challenge",
			Description:    &originalErrorMessage,
			HttpStatusCode: 500,
		}

		return nil, problem
	}

	tokensBundle := &customTypes.TokensBundle{
		AccessToken:  output.AuthenticationResult.AccessToken,
		RefreshToken: output.AuthenticationResult.RefreshToken,
	}

	return tokensBundle, nil
}

func (c *CognitoAuthService) SignUp(userEmail string, appleUserId string) (*string, error) {
	username := appleUserId
	password, err := helpers.GenerateRandomPassword()

	if err != nil {
		return nil, err
	}

	emailAttributeName := "email"

	input := &cognitoidentityprovider.SignUpInput{
		ClientId: &c.ClientId,
		Username: &username,
		Password: &password,
		UserAttributes: []types.AttributeType{
			{Name: &emailAttributeName, Value: &userEmail},
		},
	}

	_, err = c.CognitoClient.SignUp(context.Background(), input)

	if err != nil {
		originalErrorMessage := err.Error()

		problem := problems.Problem{
			Id:             "1a46c20e-774d-4e43-84c8-1486a166477d",
			Message:        "Failed to sign up in cognito",
			Description:    &originalErrorMessage,
			HttpStatusCode: 500,
		}

		return nil, problem
	}

	str := "TODO_DELETE"

	return &str, nil
}
