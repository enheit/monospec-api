package usecase

import (
	"monospec-api/api/specialist/get-specialist-profile/mappers"
	"monospec-api/api/specialist/get-specialist-profile/repos"
	"monospec-api/api/specialist/get-specialist-profile/types"
)

type GetSpecialistProfileUseCase struct {
	specialistProfileRepo *repos.SpecialistProfileRepo
	sessionUserId         int64
}

func NewGetSpecialistProfileUseCase(sessionUserId int64, specialistProfileRepo *repos.SpecialistProfileRepo) *GetSpecialistProfileUseCase {
	return &GetSpecialistProfileUseCase{
		sessionUserId:         sessionUserId,
		specialistProfileRepo: specialistProfileRepo,
	}
}

func (u *GetSpecialistProfileUseCase) Execute(specialistId int64) (*types.ResponseBody, error) {
	specialistProfile, err := u.specialistProfileRepo.GetSpecialistProfileById(specialistId)

	if err != nil {
		return nil, err
	}

	responseBody := mappers.ResponseBodyMapper(specialistProfile)

	return responseBody, nil
}
