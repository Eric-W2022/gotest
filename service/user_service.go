package service

import (
	"github.com/gin-gonic/gin"
	"wxcloudrun-golang/db/model"
)

func GetUserList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// CreateUser
// @Summary 用户模块
// @Tags 新建用户
// @param name formData string false "name"
// @param password formData string false "password"
// @param email formData string false "email"
// @param phone formData string false "phone"
// @Success 200 {string} Success
// @Router /user [post]
func CreateUser(c *gin.Context) {
	user := model.UserModel{}
	user.Name = c.Query("name")
}

func UpdateUser(c *gin.Context) {
	user := model.UserModel{}
	user.Name = c.Query("name")
}

func DeleteUser(c *gin.Context) {
	user := model.UserModel{}
	user.Name = c.Query("name")
}

func GetUser(c *gin.Context) {
	user := model.UserModel{}
	user.Name = c.Query("name")
}
