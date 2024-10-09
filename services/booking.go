package services 

type BookingService struct {
}

func NewBookingService() *BookingService {
	return &BookingService{}
}

func (b *BookingService) GetAvailableWorkingDays() string {
  return "working days"
}

func (b *BookingService) GetAvailableTimeSlots() string {
  return "time slots"
}

func (b *BookingService) Book() string {
  return "asd"
}
