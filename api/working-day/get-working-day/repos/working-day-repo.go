package repos

import (
	"monospec-api/api/working-day/get-working-day/types"

	"github.com/jackc/pgx/v5/pgxpool"
)

type WorkingDayRepo struct {
	Pool *pgxpool.Pool
}

func (r *WorkingDayRepo) GetWorkingDay(workingDayId string) (*types.WorkingDay, error) {
	workingDay := &types.WorkingDay{Id: workingDayId}

	return workingDay, nil
}
