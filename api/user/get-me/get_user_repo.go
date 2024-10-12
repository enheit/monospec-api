package getme

func getUserRepo_(deps *GetUserRepoDeps) func(bag *GetUserRepoBag) (*GetUserRepoPayload, error) {
	return func(bag *GetUserRepoBag) (*GetUserRepoPayload, error) {
		user := &User{
			Id:    "1",
			Name:  "John Doe",
			Email: "joe@me.com",
		}

		payload := &GetUserRepoPayload{
			user: user,
		}

		return payload, nil
	}
}

type GetUserRepoDeps struct {
	connectionId string
}

type GetUserRepoBag struct {
	userId string
}

type GetUserRepoPayload struct {
	user *User
}
