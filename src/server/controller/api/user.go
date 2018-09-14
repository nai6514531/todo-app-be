package controller

import (
	"net/http"
	"todo-app-be/src/server/enity"
	"todo-app-be/src/server/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 01表是系统是todo
// 01表示表是user
var (
	userMsg = map[string]string{

		"010100": "登录成功",
		"010101": "账号不存在",
		"010102": "密码错误",
		"010103": "获取用户信息成功",
		"010104": "获取用户信息失败",
		"010105": "登出成功",
		"010106": "登出失败",
	}
)

type UserController struct {
}

type Login struct {
	Account  string `form:"account" json:"account" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (User *UserController) Login(c *gin.Context) {
	userService := &service.UserService{}
	var params Login
	if err := c.ShouldBindJSON(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	account := params.Account
	password := params.Password
	user, err := userService.FindByAccount(account)
	if err != nil {
		result := &enity.Result{"010101", false, nil, userMsg["010101"]}
		c.JSON(http.StatusOK, result)
		return
	}
	if user.Password != password {
		result := &enity.Result{"010102", false, nil, userMsg["010102"]}
		c.JSON(http.StatusOK, result)
		return
	}
	result := &enity.Result{"010100", true, nil, userMsg["010100"]}
	session := c.MustGet("session").(sessions.Session)
	session.Set("userId", user.Id)
	session.Save()
	c.JSON(http.StatusOK, result)
}

func (User *UserController) GetProfile(c *gin.Context) {
	userService := &service.UserService{}
	session := c.MustGet("session").(sessions.Session)
	id := session.Get("userId").(int)
	user, err := userService.GetUserById(id)
	if err != nil {
		result := &enity.Result{"010104", false, nil, userMsg["010104"]}
		c.JSON(http.StatusOK, result)
		return
	}
	result := &enity.Result{"010103", true, user, userMsg["010103"]}
	c.JSON(http.StatusOK, result)
}
func (User *UserController) Logout(c *gin.Context) {
	session := c.MustGet("session").(sessions.Session)
	session.Clear()
	err := session.Save()
	if err != nil {
		result := &enity.Result{"010106", true, nil, userMsg["010106"]}
		c.JSON(http.StatusOK, result)
		return
	}
	result := &enity.Result{"010105", true, nil, userMsg["010105"]}
	c.JSON(http.StatusOK, result)
}
