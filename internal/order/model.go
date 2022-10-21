package order

type OrderDTO struct {
	IdUser  int    `json:"idUser"`
	IdSer   int    `json:"idSer"`
	IdOrder int    `json:"idOrder"`
	Coast   string `json:"coast"`
}

type Order struct {
	Id      int    `json:"-"`
	IdUser  int    `json:"idUser"`
	IdSer   int    `json:"idSer"`
	IdOrder int    `json:"idOrder"`
	Coast   string `json:"coast"`
	State   string `json:"state"`
}
