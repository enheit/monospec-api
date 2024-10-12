package getme

func getMeUseCase_(deps *GetMeUseCaseDeps) func(bag *GetMeUseCaseBag) (*GetMeUseCasePayload, error) {
	return func(bag *GetMeUseCaseBag) (*GetMeUseCasePayload, error) {
		getUserPayload, nil := deps.getUserRepo(&GetUserRepoBag{userId: bag.userId})

		payload := &GetMeUseCasePayload{
			user: getUserPayload.user,
		}

		return payload, nil
	}
}

type GetMeUseCaseDeps struct {
	getUserRepo func(bag *GetUserRepoBag) (*GetUserRepoPayload, error)
}

type GetMeUseCaseBag struct {
	userId string
}

type GetMeUseCasePayload struct {
	user *User
}
