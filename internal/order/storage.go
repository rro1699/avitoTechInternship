package order

type Repository interface {
	Reservation(dto OrderDTO) error
	Recognition(dto OrderDTO) error
}
