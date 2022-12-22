package models

import (
	"fmt"

	"myapp/utils"

	"gorm.io/gorm"
)

type User struct {
	ID           uint64 `json:"id" gorm:"primaryKey"`
	FirstName    string `json:"first_name" gorm:"not null"`
	LastName     string `json:"last_name" gorm:"not null"`
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique;not null"`
	PasswordHash string `json:"password_hash" gorm:"not null"`
	Password     string `json:"-" gorm:"-"`
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hash, err := utils.HashPassword(u.Password)
		if err != nil {
			return nil
		}
		tx.Statement.SetColumn("PasswordHash", hash)
	}

	fmt.Println(u.PasswordHash)
	return tx.Error
}
