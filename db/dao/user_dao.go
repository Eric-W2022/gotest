package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const userTable = "Users"

// UserInterface 计数器数据模型接口
type UserInterface interface {
	CreateUser(user *model.UserModel) error
	UpsertUser(user *model.UserModel) error
	DeleteUser(id int32) error
	GetUser(id int32) (*model.UserModel, error)
	GetUserByName(name string) (*model.UserModel, error)
	GetUserByPhone(phone string) (*model.UserModel, error)
	GetUserByEmail(email string) (*model.UserModel, error)
}

// UserInterfaceImp 计数器数据模型实现
type UserInterfaceImp struct{}

// GetUserByPhone 通过手机查询用户
func (imp *UserInterfaceImp) GetUserByPhone(phone string) (*model.UserModel, error) {
	var err error
	var user = new(model.UserModel)

	cli := db.Get()
	err = cli.Table(userTable).Where("phone = ?", phone).First(user).Error

	return user, err
}

// GetUserByEmail 通过邮箱查询用户
func (imp *UserInterfaceImp) GetUserByEmail(email string) (*model.UserModel, error) {
	var err error
	var user = new(model.UserModel)

	cli := db.Get()
	err = cli.Table(userTable).Where("email = ?", email).First(user).Error

	return user, err
}

// UserImp 实现实例
var UserImp UserInterface = &UserInterfaceImp{}

// CreateUser 添加用户
func (imp *UserInterfaceImp) CreateUser(userModel *model.UserModel) error {
	cli := db.Get()
	return cli.Table(userTable).Create(userModel).Error
}

// UpsertUser 更新/写入user
func (imp *UserInterfaceImp) UpsertUser(user *model.UserModel) error {
	cli := db.Get()
	return cli.Table(userTable).Save(user).Error
}

// DeleteUser 更新/写入user
func (imp *UserInterfaceImp) DeleteUser(id int32) error {
	cli := db.Get()
	return cli.Table(userTable).Delete(id).Error
}

// GetUser 查询Counter
func (imp *UserInterfaceImp) GetUser(id int32) (*model.UserModel, error) {
	var err error
	var user = new(model.UserModel)

	cli := db.Get()
	err = cli.Table(userTable).Where("id = ?", id).First(user).Error

	return user, err
}

// GetUserByName 查询Counter
func (imp *UserInterfaceImp) GetUserByName(name string) (*model.UserModel, error) {
	var err error
	var user = new(model.UserModel)

	cli := db.Get()
	err = cli.Table(userTable).Where("name = ?", name).First(user).Error

	return user, err
}
