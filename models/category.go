package models

import (
	"time"
)

type Category struct {
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"type:varchar(255);not_null;" json:"name" binding:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime:true;autoUpdateTime:true;" json:"created_at"`
	Books     []Book    `json:"books,omitempty" gorm:"one2many:books;constraint:OnDelete:CASCADE"`
	AdminID   uint      `gorm:"null" json:"-"`
	Admin     *User     `gorm:"foreignKey:AdminID;references:ID;constraint:OnDelete:CASCADE;" json:"-"`
}
