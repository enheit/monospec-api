package usecase

import (
	"monospec-api/api/auth/enter/services"
	"monospec-api/api/auth/enter/types"
)

type EnterUseCase struct {
	CognitoAuthService *services.CognitoAuthService
}

func (u *EnterUseCase) Enter(appleIdToken string) (*types.User, error) {

	u.CognitoAuthService.InitAuth()

	u.CognitoAuthService.SignUp()

	u.CognitoAuthService.RespondToAuthChallenge()

	return nil, nil
}
