package interfaces

import "monospec-api/api/working-day/get-working-day/types"

type WorkingDayRepo interface {
	GetWorkingDay(workingDayId string) (*types.WorkingDay, error)
}
