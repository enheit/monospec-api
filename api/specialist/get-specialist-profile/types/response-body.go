package types

import "time"

type ResponseBody struct {
	SpecialistProfile SpecialistProfileResponse `json:"specialist"`
}

type SpecialistProfileResponse struct {
	Id                 int64                    `json:"id"`
	Name               string                   `json:"name"`
	Nickname           string                   `json:"nickname"`
	Avatar             *string                  `json:"avatar,omitempty"`
	Bio                *string                  `json:"bio,omitempty"`
	Verified           bool                     `json:"verified"`
	AverageRating      float64                  `json:"averageRating"`
	AppointmentsNumber int                      `json:"appointmentsNumber"`
	ReviewsNumber      int                      `json:"aeviewsNumber"`
	ServiceGroups      *[]*ServiceGroupResponse `json:"serviceGroups"`
}

type ServiceGroupResponse struct {
	Id        int64               `json:"id"`
	Name      string              `json:"name"`
	Services  *[]*ServiceResponse `json:"services"`
	CreatedAt time.Time           `json:"created_at"`
}

type ServiceResponse struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Duration  int       `json:"duration"`
	CreatedAt time.Time `json:"created_at"`
}
