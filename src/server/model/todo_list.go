package model

type TodoList struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (TodoList) TableName() string {
	return "todo_list"
}
