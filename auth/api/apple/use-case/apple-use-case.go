package usecase

import (
	"fmt"
	"monospec-api/auth/api/apple/repos"
	"monospec-api/auth/api/apple/services"
	"monospec-api/auth/api/apple/types"

	authServices "monospec-api/auth/services"
)

type AppleUseCase struct {
	TokenService              *authServices.TokenService
	AppleIdentityTokenService *services.AppleIdentityTokenService
	UserRepo                  *repos.UserRepo
}

func (u *AppleUseCase) Enter(rawIdentityToken string, firstName string) (*AppleUseCasePayload, error) {
	publicKeys, err := u.AppleIdentityTokenService.GetPublicKeys()

	fmt.Printf("publicKeys: %+v\n", publicKeys)

	if err != nil {
		return nil, err
	}

	appleIdentityToken, err := u.AppleIdentityTokenService.VerifyIdentityToken(rawIdentityToken, publicKeys)

	println("appleIdentityToken: ", appleIdentityToken)

	if err != nil {
		return nil, err
	}

	user, err := u.UserRepo.GetUserByAppleSub(appleIdentityToken.Payload.Subject)

	println("user: ", user)

	if err != nil {
		return nil, err
	}

	if user == nil {
		appleSub := appleIdentityToken.Payload.Subject
		email := appleIdentityToken.Payload.Email
		isEmailVerified := appleIdentityToken.Payload.IsEmailVerified

		newUser, err := u.UserRepo.CreateUser(appleSub, firstName, email, isEmailVerified)

		println("newUser: ", newUser)

		if err != nil {
			return nil, err
		}

		user = newUser
	} else {
		u.UserRepo.UpdateUserLoginAt(user.Id)
	}

	rawAccessToken, err := u.TokenService.CreateAccessToken(appleIdentityToken.Payload.Subject)

	println("rawAccessToken: ", rawAccessToken)

	if err != nil {
		return nil, err
	}

	rawRefreshToken, err := u.TokenService.CreateRefreshToken(appleIdentityToken.Payload.Subject)

	println("rawRefreshToken: ", rawRefreshToken)

	if err != nil {
		return nil, err
	}

	payload := &AppleUseCasePayload{
		User:         *user,
		AccessToken:  *rawAccessToken,
		RefreshToken: *rawRefreshToken,
	}

	println("payload: ", payload)

	return payload, nil
}

type AppleUseCasePayload struct {
	User         types.User
	AccessToken  string
	RefreshToken string
}
