package dao

import (
	"wxcloudrun-golang/db"
	"wxcloudrun-golang/db/model"
)

const userTable = "users"

// CreateUser 添加用户
func (imp *CounterInterfaceImp) CreateUser(userModel *model.UserModel) error {
	cli := db.Get()
	return cli.Table(userTable).Create(userModel).Error
}

// UpsertUser 更新/写入user
func (imp *CounterInterfaceImp) UpsertUser(counter *model.CounterModel) error {
	cli := db.Get()
	return cli.Table(userTable).Save(counter).Error
}

// GetUser 查询Counter
func (imp *CounterInterfaceImp) GetUser(id int32) (*model.CounterModel, error) {
	var err error
	var counter = new(model.CounterModel)

	cli := db.Get()
	err = cli.Table(userTable).Where("id = ?", id).First(counter).Error

	return counter, err
}
