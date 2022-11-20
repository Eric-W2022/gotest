package model

import (
	"gorm.io/gorm"
	"time"
)

type UserModel struct {
	Id           int32          `gorm:"column:id" json:"id"`
	Name         string         `gorm:"column:name" json:"name"`
	Password     string         `gorm:"column:password" json:"password"`
	RandomNumber string         `gorm:"column:random_number" json:"random_number"`
	Phone        string         `gorm:"column:phone" json:"phone"`
	Email        string         `gorm:"column:email" json:"email"`
	LoginTime    time.Time      `gorm:"column:login_time" json:"login_time"`
	CreatedAt    time.Time      `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt    time.Time      `gorm:"column:updatedAt" json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deletedAt" json:"deletedAt"`
}
