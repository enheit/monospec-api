package usecase

import (
	"monospec-api/api/working-day/get-working-days/interfaces"
	"monospec-api/api/working-day/get-working-days/types"
)

type GetWorkingDaysUseCase struct {
	WorkingDayRepo interfaces.WorkingDayRepo
}

func (u *GetWorkingDaysUseCase) GetWorkingDays(sessionUserId string) ([]*types.WorkingDay, error) {
	workingDays, nil := u.WorkingDayRepo.GetWorkingDays()

	return workingDays, nil
}
