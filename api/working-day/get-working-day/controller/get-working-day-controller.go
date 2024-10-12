package controller

import (
	"monospec-api/api/working-day/get-working-day/repos"
	"monospec-api/api/working-day/get-working-day/types"
	"monospec-api/api/working-day/get-working-day/use-case"
)

type Controller struct {
}

func (c *Controller) Execute(workingDayId string, sessionUserId string) *types.WorkingDay {
	workingDayRepo := repos.WorkingDayRepo{Db: nil}

	useCase := usecase.GetWorkingDayUseCase{WorkingDayRepo: &workingDayRepo}

	workingDay, _ := useCase.GetWorkingDay(workingDayId, sessionUserId)

	return workingDay
}
