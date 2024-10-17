package controller

import (
	"monospec-api/auth/middlewares/verify-access-token/use-case"
	"monospec-api/auth/services"
	"monospec-api/auth/types"
)

type VerifyAccessTokenController struct {
}

func New() *VerifyAccessTokenController {
	return &VerifyAccessTokenController{}
}

func (c *VerifyAccessTokenController) Execute(rawAccessToken string) (*types.AccessToken, error) {
	tokenService := services.TokenService{}

	useCase := usecase.VerifyAccessTokenUseCase{
		TokenService: &tokenService,
	}

	accessToken, err := useCase.VerifyAccessToken(rawAccessToken)

	if err != nil {
		return nil, err
	}

	return accessToken, nil
}
