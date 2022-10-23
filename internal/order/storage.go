package order

type Repository interface {
	Create(dto OrderDTO) error
	GetStateByParams(dto OrderDTO) Order
	UpdateOrder(dto OrderDTO) error
}
