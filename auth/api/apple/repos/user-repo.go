package repos

import "monospec-api/auth/api/apple/types"

type UserRepo struct {
}

func (u *UserRepo) FindUserByAppleSub(sub string) (*types.User, error) {
	return nil, nil
}
