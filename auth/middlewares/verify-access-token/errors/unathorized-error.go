package errors

type Unathorized struct{}

func (u *Unathorized) Error() string {
	return "Unathorized"
}
