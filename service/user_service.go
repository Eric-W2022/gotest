package service

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
)

func GetUserList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// CreateUser
// @Tags 用户
// @Summary 新建用户
// @param name formData string false "name"
// @param password formData string false "password"
// @param re_password formData string false "re_password"
// @param email formData string false "email"
// @param phone formData string false "phone"
// @Success 200 {string} Success
// @Router /user [post]
func CreateUser(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")
	rePassword := c.PostForm("re_password")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	if password == rePassword {
		user := &model.UserModel{
			Name:      name,
			Password:  password,
			Phone:     phone,
			Email:     email,
			LoginTime: time.Now(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		CreateUserErr := dao.UserImp.CreateUser(user)
		if CreateUserErr != nil {
			c.JSON(200, gin.H{
				"status":  "fail",
				"message": "创建用户失败",
			})
		}
	} else {
		c.JSON(200, gin.H{
			"status":  "fail",
			"message": "密码不一致",
		})
	}

}

// UpdateUser
// @Tags 用户
// @Summary 更新用户
// @param name formData string false "name"
// @param password formData string false "password"
// @param new_password formData string false "new_password"
// @param email formData string false "email"
// @param phone formData string false "phone"
// @Success 200 {string} Success
// @Router /user [patch]
func UpdateUser(c *gin.Context) {
	user := model.UserModel{}
	user.Name = c.Query("name")
}

// DeleteUser
// @Tags 用户
// @Summary 删除模块
// @param id formData int false "id"
// @Success 200 {string} Success
// @Router /user [delete]
func DeleteUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	DeleteUserErr := dao.UserImp.DeleteUser(int32(id))
	if DeleteUserErr != nil {
		c.JSON(200, gin.H{
			"status":  "fail",
			"message": "删除用户失败",
		})
	} else {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "删除用户成功",
		})
	}

}

// GetUser
// @Tags 用户
// @Summary 删除用户
// @param id formData int false "id"
// @Success 200 {string} Success
// @Router /user [get]
func GetUser(c *gin.Context) {
	user := model.UserModel{}
	user.Name = c.Query("name")
	c.JSON(200, gin.H{
		"status":  "fail",
		"message": "创建用户失败",
	})
}
