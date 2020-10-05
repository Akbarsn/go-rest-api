package model

import "github.com/jinzhu/gorm"

type User struct {
	ID           int    `gorm:"primaryKey;autoIncrement"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	FullName     string `json:"fullName"`
	PhotoProfile string `json:"photoProfile"`
}

func MigrateUser(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
