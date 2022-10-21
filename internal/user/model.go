package user

type UserDTO struct {
	Id    int    `json:"id"`
	Coast string `json:"coast"`
}

type User struct {
	Id   int    `json:"-"`
	CurB string `json:"curB"`
	ResB string `json:"resB"`
}
