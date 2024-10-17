package controller

import (
	"context"
	"encoding/json"
	"monospec-api/auth/api/apple/repos"
	"monospec-api/auth/api/apple/services"
	"monospec-api/auth/api/apple/types"
	"monospec-api/auth/api/apple/use-case"
	"monospec-api/auth/api/apple/validators"
	authServices "monospec-api/auth/services"
	"monospec-api/shared/enums"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AppleController struct {
	context context.Context
	dbPool  *pgxpool.Pool
}

func New(dbPool *pgxpool.Pool, context context.Context) *AppleController {
	return &AppleController{
		dbPool:  dbPool,
		context: context,
	}
}

func (c *AppleController) Execute(rawRequestBody string) (*string, error) {
	println("Request body is received")
	requestBody, err := validators.ValidateRequestBody(rawRequestBody)
	println("Request body is validated")

	if err != nil {
		return nil, err
	}

	userRepo := repos.NewUserRepo(c.dbPool, c.context)
	tokenService := authServices.NewTokenService(os.Getenv(enums.JWTPrivateKey))
	appleIdentityTokenService := services.NewAppleIdentityTokenService()

	useCase := &usecase.AppleUseCase{
		TokenService:              tokenService,
		AppleIdentityTokenService: appleIdentityTokenService,
		UserRepo:                  userRepo,
	}

	useCasePayload, err := useCase.Enter(requestBody.IdentityToken, requestBody.FirstName)

	if err != nil {
		return nil, err
	}

	responseBody := &types.ResponseBody{
		User: types.UserResponse{
			Id:    useCasePayload.User.Id,
			Name:  useCasePayload.User.Name,
			Email: useCasePayload.User.Email,
		},
		AccessToken:  useCasePayload.AccessToken,
		RefreshToken: useCasePayload.RefreshToken,
	}

	rawResponseBodyBytes, err := json.Marshal(responseBody)

	if err != nil {
		return nil, err
	}

	rawResponseBody := string(rawResponseBodyBytes)

	return &rawResponseBody, nil
}
