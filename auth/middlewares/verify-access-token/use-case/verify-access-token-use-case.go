package usecase

import (
	"monospec-api/auth/services"
	"monospec-api/auth/types"
)

type VerifyAccessTokenUseCase struct {
	TokenService *services.TokenService
}

func (v *VerifyAccessTokenUseCase) VerifyAccessToken(rawAccessToken string) (*types.AccessToken, error) {
	accessToken, err := v.TokenService.VerifyAccessToken(rawAccessToken)

	if err != nil {
		return nil, err
	}

	return accessToken, nil
}
