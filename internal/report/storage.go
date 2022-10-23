package report

type Repository interface {
	GetServiceById(idSer int) Serv
}
