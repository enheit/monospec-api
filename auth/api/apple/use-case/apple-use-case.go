package usecase

import (
	"monospec-api/auth/api/apple/services"
	"monospec-api/auth/api/apple/types"
	sharedServices "monospec-api/auth/services"
)

type AppleUseCase struct {
	TokenService              *sharedServices.TokenService
	AppleIdentityTokenService *services.AppleIdentityTokenService
}

func (u *AppleUseCase) Enter(rawIdentityToken string) (*AppleUseCasePayload, error) {
	publicKeys, err := u.AppleIdentityTokenService.GetPublicKeys()

	if err != nil {
		return nil, err
	}

	appleIdentityToken, err := u.AppleIdentityTokenService.VerifyIdentityToken(rawIdentityToken, publicKeys)

	if err != nil {
		return nil, err
	}

	rawAccessToken, err := u.TokenService.CreateAccessToken(appleIdentityToken.Payload.Subject)

	if err != nil {
		return nil, err
	}

	rawRefreshToken, err := u.TokenService.CreateRefreshToken(appleIdentityToken.Payload.Subject)

	if err != nil {
		return nil, err
	}

	user := &types.User{
		Id:    "TODO_USER_ID",
		Name:  "TODO_USER_NAME",
		Email: "TODO_USER_EMAIL",
	}

	payload := &AppleUseCasePayload{
		User:         *user,
		AccessToken:  *rawAccessToken,
		RefreshToken: *rawRefreshToken,
	}

	return payload, nil
}

type AppleUseCasePayload struct {
	User         types.User
	AccessToken  string
	RefreshToken string
}
