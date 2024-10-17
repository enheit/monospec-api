package usecase

type LogoutUseCase struct {
}

func New() *LogoutUseCase {
	return &LogoutUseCase{}
}

func (u *LogoutUseCase) Logout() {
}
