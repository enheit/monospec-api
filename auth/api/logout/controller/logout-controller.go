package controller

import "monospec-api/auth/api/logout/use-case"

type LogoutController struct {
}

func New() *LogoutController {
	return &LogoutController{}
}

func (c *LogoutController) Execute() {
	useCase := usecase.New()

	useCase.Logout()
}
