package model

type User struct {
	Id       int    `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (User) TableName() string {
	return "user"
}
