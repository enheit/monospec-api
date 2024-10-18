package mappers

import "monospec-api/api/specialist/get-specialist-profile/types"

func ResponseBodyMapper(specialistProfile *types.SpecialistProfile) *types.ResponseBody {
	var serviceGroupsResponse []*types.ServiceGroupResponse

	for _, serviceGroup := range *specialistProfile.ServiceGroups {
		var servicesResponse []*types.ServiceResponse

		for _, service := range *serviceGroup.Services {
			serviceResponse := &types.ServiceResponse{
				Id:       service.Id,
				Name:     service.Name,
				Price:    service.Price,
				Duration: service.Duration,
			}

			servicesResponse = append(servicesResponse, serviceResponse)
		}

		serviceGroupResponse := &types.ServiceGroupResponse{
			Id:        serviceGroup.Id,
			Name:      serviceGroup.Name,
			CreatedAt: serviceGroup.CreatedAt,
			Services:  &servicesResponse,
		}

		serviceGroupsResponse = append(serviceGroupsResponse, serviceGroupResponse)
	}

	responseBody := &types.ResponseBody{
		SpecialistProfile: types.SpecialistProfileResponse{
			Id:                 specialistProfile.Id,
			Name:               specialistProfile.Name,
			Bio:                specialistProfile.Bio,
			Avatar:             specialistProfile.Avatar,
			Nickname:           specialistProfile.Nickname,
			Verified:           specialistProfile.Verified,
			AppointmentsNumber: specialistProfile.AppointmentsNumber,
			AverageRating:      specialistProfile.AverageRating,
			ReviewsNumber:      specialistProfile.ReviewsNumber,
			ServiceGroups:      &serviceGroupsResponse,
		},
	}

	return responseBody
}
