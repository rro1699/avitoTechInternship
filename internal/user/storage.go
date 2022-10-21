package user

type Repository interface {
	Accrual(dto UserDTO) error
	GetBalance(dto UserDTO) (User, error)
}
