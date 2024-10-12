package controller

import (
	"monospec-api/api/working-day/get-working-days/repos"
	"monospec-api/api/working-day/get-working-days/types"
	"monospec-api/api/working-day/get-working-days/use-case"
	"time"
)

type GetWorkingDaysController struct {
}

func (c *GetWorkingDaysController) Execute(sessionUserId string) *types.WorkingDay {
	workingDayRepo := repos.WorkingDayRepo{Db: nil}

	useCase := usecase.GetWorkingDaysUseCase{WorkingDayRepo: &workingDayRepo}

	useCase.GetWorkingDays(sessionUserId)

	return &types.WorkingDay{Id: "1", CreatedAt: time.Now()}
}
