package model

import "time"

type UserModel struct {
	Id        int32     `gorm:"column:id" json:"id"`
	Name      string    `gorm:"column:name" json:"name"`
	Password  string    `gorm:"column:password" json:"password"`
	Phone     string    `gorm:"column:phone" json:"phone"`
	Email     string    `gorm:"column:email" json:"email"`
	LoginTime time.Time `gorm:"column:login_time" json:"login_time"`
	CreatedAt time.Time `gorm:"column:createdAt" json:"createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt" json:"updatedAt"`
	DeletedAt time.Time `gorm:"column:deletedAt" json:"deletedAt"`
}
