package services

type AuthService struct {
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (a *AuthService) SignIn() string {
	return "asd"
}

func (a *AuthService) SignUp() string {
  return "asd"
}

func (a *AuthService) SignOut() string {
  return "asd"
}
