package service

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
	"wxcloudrun-golang/db/dao"
	"wxcloudrun-golang/db/model"
	"wxcloudrun-golang/utils"
)

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

	// 密码校验
	// 确认密码一致校验
	if password != rePassword {
		c.JSON(200, gin.H{
			"status":  "fail",
			"message": "两次密码不一致",
		})
		return
	}
	// 密码长度校验
	if len(password) < 8 || len(password) > 16 {
		c.JSON(200, gin.H{
			"status":  "fail",
			"message": "密码需要8-16个字符",
		})
		return
	}
	// 用户名校验
	// 用户名非空校验
	if name == "" {
		c.JSON(200, gin.H{
			"status":  "fail",
			"message": "用户名不能为空",
		})
		return
	}
	// 用户名长度校验
	if len(name) < 6 || len(name) > 40 {
		c.JSON(200, gin.H{
			"status":  "fail",
			"message": "用户名需在6-20字符内",
		})
		return
	}
	// 用户名重复注册校验
	registeredMember, _ := dao.UserImp.GetUserByName(name)
	fmt.Println(registeredMember)
	if registeredMember.Name != "" {
		c.JSON(200, gin.H{
			"status":  "fail",
			"message": "用户名已存在",
		})
		return
	}

	// 手机重复注册校验
	registeredMember, _ = dao.UserImp.GetUserByEmail(email)
	fmt.Println(registeredMember)
	if registeredMember.Name != "" {
		c.JSON(200, gin.H{
			"status":  "fail",
			"message": "邮箱已存在",
		})
		return
	}

	// 建立用户对象
	randomNumber := fmt.Sprintf("%06d", rand.Int31())
	password = utils.EncodePassword(password, randomNumber)
	user := &model.UserModel{
		Name:         name,
		Password:     password,
		RandomNumber: randomNumber,
		Phone:        phone,
		Email:        email,
		LoginTime:    time.Now(),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// 创建用户
	CreateUserErr := dao.UserImp.CreateUser(user)
	if CreateUserErr != nil {
		c.JSON(200, gin.H{
			"status":  "fail",
			"message": "创建用户失败",
		})
		return
	} else {
		c.JSON(200, gin.H{
			"status":  "successful",
			"message": "创建用户成功",
		})
		return
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
	//user := model.UserModel{}
	//user.Name = c.Query("name")
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

// GetUserInfo
// @Tags 用户
// @Summary 查询用户
// @param id formData int false "id"
// @Success 200 {string} Success
// @Router /user [get]
func GetUserInfo(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	//user, _ := c.Get(utils.IdentityKey)
	//fmt.Println(user)
	c.JSON(200, gin.H{
		"userID":   claims[utils.UserIdKey],
		"userName": claims[utils.UserNameKey],
		"identity": claims[utils.IdentityKey],
	})
}
