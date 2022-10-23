package report

type ReportDTO struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

type Serv struct {
	IdSer int    `json:"idSer"`
	Name  string `json:"name"`
}
