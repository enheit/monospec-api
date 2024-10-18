package types

import "time"

type SpecialistProfile struct {
	Id                 int64
	Name               string
	Nickname           string
	Avatar             *string
	Bio                *string
	Verified           bool
	AverageRating      float64
	AppointmentsNumber int
	ReviewsNumber      int
	ServiceGroups      *[]*ServiceGroup
}

type ServiceGroup struct {
	Id        int64
	Name      string
	Services  *[]*Service
	CreatedAt time.Time
}

type Service struct {
	Id        int64
	Name      string
	Price     float64
	Duration  int
	CreatedAt time.Time
}
