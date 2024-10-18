package controller

import (
	"context"
	"encoding/json"
	"monospec-api/api/specialist/get-specialist-profile/repos"
	"monospec-api/api/specialist/get-specialist-profile/use-case"
	"monospec-api/api/specialist/get-specialist-profile/validators"

	"github.com/jackc/pgx/v5/pgxpool"
)

type GetSpecialistProfileController struct {
	dbPool        *pgxpool.Pool
	context       context.Context
	sessionUserId int64
}

func NewGetSpecialistProfileController(dbPool *pgxpool.Pool, context context.Context, sessionUserId int64) *GetSpecialistProfileController {
	return &GetSpecialistProfileController{
		dbPool:        dbPool,
		context:       context,
		sessionUserId: sessionUserId,
	}
}

func (c *GetSpecialistProfileController) Execute(rawPathParams map[string]string) (*string, error) {
	pathParams, err := validators.ValidatePathParams(rawPathParams)

	if err != nil {
		return nil, err
	}

	specialistProfileRepo := repos.NewSpecialistProfileRepo(c.dbPool, c.context)
	useCase := usecase.NewGetSpecialistProfileUseCase(c.sessionUserId, specialistProfileRepo)

	responseBody, err := useCase.Execute(pathParams.SpecialistId)

	if err != nil {
		return nil, err
	}

	rawResponseBodyBytes, err := json.Marshal(responseBody)

	if err != nil {
		return nil, err
	}

	rawResponseBody := string(rawResponseBodyBytes)

	return &rawResponseBody, nil
}
