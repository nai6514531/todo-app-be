package controller

import (
	"fmt"
	"net/http"
	"todo-app-be/src/server/enity"
	"todo-app-be/src/server/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 01表是系统是todo
// 02表示表是toda
var (
	todoMsg = map[string]string{
		"010200": "获取todo列表成功",
	}
)

type TodoController struct {
}

func (todo *TodoController) GetList(c *gin.Context) {
	todoService := &service.TodoService{}
	session := c.MustGet("session").(sessions.Session)
	userId := session.Get("userId").(int)
	list, err := todoService.GetListById(userId)
	if err != nil {
		fmt.Println("err", err)
	}
	result := &enity.Result{"010200", true, list, todoMsg["010200"]}
	c.JSON(http.StatusOK, result)
}
