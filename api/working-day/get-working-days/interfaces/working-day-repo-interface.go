package interfaces

import "monospec-api/api/working-day/get-working-days/types"

type WorkingDayRepo interface {
	GetWorkingDays() ([]*types.WorkingDay, error)
}
