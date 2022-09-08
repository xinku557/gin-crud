package models

import "time"

type BAndR struct {
	ID         uint      `gorm:"primary_key;auto_increment" json:"id"`
	UserID     uint      `json:"user_id" gorm:"null"`
	AdminID    *uint     `json:"admin_id" gorm:"null"`
	BookID     uint      `json:"book_id"`
	Status     uint      `gorm:"null;default:1;" json:"status"`
	BorrowDate string    `gorm:"null" json:"borrow_date"`
	ReturnDate string    `gorm:"null" json:"return_date"`
	CreatedAt  time.Time `gorm:"autoCreateTime:true;autoUpdateTime:true;" json:"created_at"`
	Book       *Book     `gorm:"foreignKey:BookID;references:ID;constraint:OnDelete:CASCADE;" json:"book"`
	User       *User     `gorm:"foreignKey:UserID;references:ID;constraint:OnDelete:CASCADE;" json:"user"`
	Admin      *User     `gorm:"foreignKey:AdminID;references:ID;constraint:OnDelete:CASCADE;" json:"-"`
}
