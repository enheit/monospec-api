package getme

func getMeController(userId string) {
	getUserRepo := getUserRepo_(&GetUserRepoDeps{connectionId: ""})
	getMeUseCase := getMeUseCase_(&GetMeUseCaseDeps{getUserRepo: getUserRepo})

	user, err := getMeUseCase(&GetMeUseCaseBag{userId: userId})

	print(err)

	println(user)
}
