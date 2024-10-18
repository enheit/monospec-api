package repos

import (
	"context"
	"monospec-api/api/specialist/get-specialist-profile/types"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type SpecialistProfileRepo struct {
	dbPool  *pgxpool.Pool
	context context.Context
}

func NewSpecialistProfileRepo(dbPool *pgxpool.Pool, context context.Context) *SpecialistProfileRepo {
	return &SpecialistProfileRepo{
		dbPool:  dbPool,
		context: context,
	}
}

func (s *SpecialistProfileRepo) GetSpecialistProfileById(specialistId int64) (*types.SpecialistProfile, error) {
	query := `
    SELECT 
      s.id AS specialist_id,
      s.name AS specialist_name,
      s.nickname,
      s.avatar,
      s.bio,
      s.average_rating,
      s.appointments_number,
      s.reviews_number,
      s.verified,
      sg.id AS service_group_id,
      sg.name AS service_group_name,
      sv.id AS service_id,
      sv.name AS service_name,
      sv.price,
      sv.duration,
      sv.created_at AS service_created_at
    FROM specialists s
    LEFT JOIN service_groups sg ON sg.specialist_id = s.id AND sg.deleted_at IS NULL
    LEFT JOIN services sv ON sv.service_group_id = sg.id AND sv.deleted_at IS NULL
    WHERE s.id = $1 AND s.deleted_at IS NULL;
  `

	rows, err := s.dbPool.Query(s.context, query, specialistId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var specialist *types.SpecialistProfile
	serviceGroupMap := make(map[int64]*types.ServiceGroup)

	for rows.Next() {
		var (
			serviceGroupId   *int64
			serviceGroupName *string
			serviceId        *int64
			serviceName      *string
			servicePrice     *float64
			serviceDuration  *int
			serviceCreatedAt *string
		)

		specialistTemp := &types.SpecialistProfile{}

		// Scan each row
		err := rows.Scan(
			&specialistTemp.Id,
			&specialistTemp.Name,
			&specialistTemp.Nickname,
			&specialistTemp.Avatar,
			&specialistTemp.Bio,
			&specialistTemp.AverageRating,
			&specialistTemp.AppointmentsNumber,
			&specialistTemp.ReviewsNumber,
			&specialistTemp.Verified,
			&serviceGroupId,
			&serviceGroupName,
			&serviceId,
			&serviceName,
			&servicePrice,
			&serviceDuration,
			&serviceCreatedAt,
		)
		if err != nil {
			return nil, err
		}

		// Initialize the specialist only once
		if specialist == nil {
			specialist = specialistTemp
			specialist.ServiceGroups = &[]*types.ServiceGroup{}
		}

		// Handle service groups
		if serviceGroupId != nil {
			group, exists := serviceGroupMap[*serviceGroupId]

			if !exists {
				// Create a new service group
				group = &types.ServiceGroup{
					Id:       *serviceGroupId,
					Name:     *serviceGroupName,
					Services: &[]*types.Service{},
				}

				// Add the new group to the serviceGroupMap
				serviceGroupMap[*serviceGroupId] = group

				*specialist.ServiceGroups = append(*specialist.ServiceGroups, group)
			}

			// Handle services
			if serviceId != nil {
				service := &types.Service{
					Id:       *serviceId,
					Name:     *serviceName,
					Price:    *servicePrice,
					Duration: *serviceDuration,
					CreatedAt: func() time.Time {
						if serviceCreatedAt != nil {
							parsedTime, _ := time.Parse(time.RFC3339, *serviceCreatedAt)
							return parsedTime
						}
						return time.Time{}
					}(),
				}
				// Add the service to the group
				*group.Services = append(*group.Services, service)
			}
		}
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	if specialist == nil {
		return nil, nil
	}

	return specialist, nil
}
