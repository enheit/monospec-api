package repos

type UserRepo struct {

}

func NewUserRepo() *UserRepo {
  return &UserRepo{}
}

func (u *UserRepo) GetUser() string {
  return "asd"
}

func (u *UserRepo) RemoveUser() string {
  return "removed"
}
