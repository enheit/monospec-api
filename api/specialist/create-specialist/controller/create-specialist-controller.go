package controller

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type CreateSpecialistController struct {
	context       context.Context
	dbPool        *pgxpool.Pool
	sessionUserId string
}

func New(dbPool *pgxpool.Pool, context context.Context, sessionUserId string) *CreateSpecialistController {
	return &CreateSpecialistController{
		dbPool:        dbPool,
		context:       context,
		sessionUserId: sessionUserId,
	}
}

func (c *CreateSpecialistController) Execute(rawRequestBody string) (*string, error) {
	return nil, nil
}
