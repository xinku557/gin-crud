package models

import (
	"time"
)

type Book struct {
	ID          uint      `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `gorm:"type:varchar(255);not_null" json:"title"`
	Description string    `gorm:"type:text;null" json:"description"`
	AuthorID    uint      `gorm:"null" json:"author_id,omitempty"`
	AdminID     uint      `json:"-"`
	CategoryID  uint      `gorm:"null" json:"category_id,omitempty"`
	CreatedAt   time.Time `gorm:"autoCreateTime:true" json:"created_at"`
	Author      *Author   `gorm:"foreignKey:AuthorID;references:ID;constraint:OnDelete:CASCADE;" json:"author,omitempty"`
	Admin       *User     `gorm:"foreignKey:AdminID;references:ID;constraint:OnDelete:CASCADE;" json:"-"`
	Category    *Category `gorm:"foreignKey:CategoryID;references:ID;constraint:OnDelete:CASCADE;" json:"category,omitempty"`
}
