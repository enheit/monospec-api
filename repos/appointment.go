package repos

type AppointmentRepo struct {
}

func NewAppointmentRepo() *AppointmentRepo {
	return &AppointmentRepo{}
}

func (a *AppointmentRepo) GetAppointmentDetails(appointmentId string) string {
	return "appointments"
}

func (a *AppointmentRepo) CreateAppointment() string {
	return "removed"
}

func (a *AppointmentRepo) ConfirmAppointment() string {
	return "confirmed"
}

func (a *AppointmentRepo) DeclineAppointment() string {
	return "declined"
}

func (a *AppointmentRepo) CancelAppointment() string {
	return "cancelled"
}
