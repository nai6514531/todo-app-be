package service

import (
	"todo-app-be/src/server/common"
	"todo-app-be/src/server/model"
)

type UserService struct {
}

func (self *UserService) FindByAccount(account string) (*model.User, error) {
	user := &model.User{}
	r := common.TodoDB_WR.Where("account = ?", account).First(user)
	if r.Error != nil {
		return nil, r.Error
	}
	return user, nil
}
func (self *UserService) GetUserById(id int) (*model.User, error) {
	user := &model.User{}
	r := common.TodoDB_WR.Where("id = ?", id).First(user)
	if r.Error != nil {
		return nil, r.Error
	}
	return user, nil
}
