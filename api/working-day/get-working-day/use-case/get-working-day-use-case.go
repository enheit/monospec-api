package usecase

import (
	"monospec-api/api/working-day/get-working-day/interfaces"
	"monospec-api/api/working-day/get-working-day/types"
)

type GetWorkingDayUseCase struct {
	WorkingDayRepo interfaces.WorkingDayRepo
}

func (u *GetWorkingDayUseCase) GetWorkingDay(workingDayId string, sessionUserId string) (*types.WorkingDay, error) {
	workingDay, err := u.WorkingDayRepo.GetWorkingDay(workingDayId)

	return workingDay, err
}
