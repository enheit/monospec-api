package repos

import (
	"monospec-api/api/working-day/get-working-days/types"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type WorkingDayRepo struct {
	Db *pgxpool.Pool
}

func (r *WorkingDayRepo) GetWorkingDays() ([]*types.WorkingDay, error) {
	workingDays := []*types.WorkingDay{
		{Id: "1", CreatedAt: time.Now()},
		{Id: "2", CreatedAt: time.Now()},
		{Id: "3", CreatedAt: time.Now()},
		{Id: "4", CreatedAt: time.Now()},
	}

	return workingDays, nil
}
