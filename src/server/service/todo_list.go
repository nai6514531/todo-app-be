package service

import (
	"todo-app-be/src/server/common"
	"todo-app-be/src/server/model"
)

type TodoService struct {
}

func (self *TodoService) GetListById(id int) (*[]*model.TodoList, error) {
	todoList := &[]*model.TodoList{}
	r := common.TodoDB_WR.Where("user_id = ?", id).Find(todoList)
	if r.Error != nil {
		return nil, r.Error
	}
	return todoList, r.Error
}
