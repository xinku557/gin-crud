package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID         uint       `gorm:"primary_key;auto_increment" json:"id"`
	Username   string     `gorm:"type:varchar(100);unique;not_null;" json:"username" binding:"required"`
	Email      string     `gorm:"type:varchar(255);unique;not_null" json:"email" binding:"required"`
	Password   string     `gorm:"type:varchar(255);not_null" json:"password,omitempty" binding:"required"`
	Type       uint       `gorm:"default:1;" json:"-"`
	CreatedAt  time.Time  `gorm:"autoCreateTime:true;autoUpdateTime:true;" json:"created_at"`
	Books      []Book     `json:"books,omitempty" gorm:"foreignKey:AdminID;references:ID;constraint:OnDelete:CASCADE"`
	Authors    []Author   `json:"authors,omitempty" gorm:"foreignKey:AdminID;references:ID;constraint:OnDelete:CASCADE"`
	Categories []Category `json:"categories,omitempty" gorm:"foreignKey:AdminID;references:ID;constraint:OnDelete:CASCADE"`
}

type Login struct {
	User     string `json:"user" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ReturnUser struct {
	Username   string     `json:"username"`
	Email      string     `json:"email"`
	Token      string     `json:"token,omitempty"`
	CreatedAt  time.Time  `json:"created_at"`
	Books      []Book     `json:"books,omitempty"`
	Authors    []Author   `json:"authors,omitempty"`
	Categories []Category `json:"categories,omitempty"`
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("password", hashedPassword)

	return nil
}
