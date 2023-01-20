package user

type UserRepository interface {
	FindByUsernameAndPassword(username, password string) (*User, error)
}
