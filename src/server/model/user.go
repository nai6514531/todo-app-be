package model

type User struct {
	Id       int    `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Name     string `json:"name"`
	NickName string `json:"nickName"`
}

func (User) TableName() string {
	return "user"
}
