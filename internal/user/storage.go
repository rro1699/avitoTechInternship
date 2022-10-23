package user

type Repository interface {
	GetUserById(user UserDTO) User
	UpdateUser(localUser User) error
	CreateUser(localUser User) (string, error)
}
