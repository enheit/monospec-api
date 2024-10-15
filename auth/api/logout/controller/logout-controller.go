package controller

import usecase "monospec-api/api/auth/logout/use-case"

type LogoutController struct {
}

func (c *LogoutController) Execute() {
	useCase := &usecase.LogoutUseCase{}

	useCase.Logout()
}
