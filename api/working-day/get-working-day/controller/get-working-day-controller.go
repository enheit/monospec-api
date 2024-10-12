package controller

import (
	"monospec-api/api/working-day/get-working-day/repos"
	"monospec-api/api/working-day/get-working-day/types"
	"monospec-api/api/working-day/get-working-day/use-case"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Controller struct {
	Pool *pgxpool.Pool
}

func (c *Controller) Execute(workingDayId string, sessionUserId string) *types.WorkingDay {
	workingDayRepo := repos.WorkingDayRepo{Pool: c.Pool}

	useCase := usecase.GetWorkingDayUseCase{WorkingDayRepo: &workingDayRepo}

	workingDay, _ := useCase.GetWorkingDay(workingDayId, sessionUserId)

	return workingDay
}
