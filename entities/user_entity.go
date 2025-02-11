package entities

import (
	"auth-sederhana-go-fiber/utilities"
	"gorm.io/gorm"
)

type User struct {
	Id       int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"type:varchar(255);not null" json:"username"`
	Email    string `gorm:"type:varchar(255);not null;uniqueIndex" json:"email"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var err error
	u.Password, err = utilities.HashPassword(u.Password)
	if err != nil {
		return err
	}

	return nil
}
