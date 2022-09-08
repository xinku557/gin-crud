package services

import (
	"gin-crud/models"
	"time"

	"gorm.io/gorm"
)

func CheckActiveStatus(db *gorm.DB, id int) bool {
	var d []models.BAndR
	return db.Where("user_id=?", id).Where("status = 1 or status = 2").Find(&d).RowsAffected > 3
}

func BRList(db *gorm.DB, data *[]models.BAndR, status int) (err error) {
	if status == 0 {
		return db.Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username", "email")
		}).Preload("Book").Find(&data).Order("id desc").Error
	}
	return db.Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username", "email")
	}).Preload("Book").Where("status = ?", status).Find(&data).Error
}

func UserBorrow(db *gorm.DB, data *models.BAndR) (err error) {
	return db.Create(&data).Error
}

func UserBorrowControl(db *gorm.DB, id int, status int, admin uint) (err error) {
	data := models.BAndR{Status: uint(status), BorrowDate: time.Now().Format("2006-01-02 15:04:05"), AdminID: &admin}
	if status == 3 {
		data.ReturnDate = time.Now().Format("2006-01-02 15:04:05")
	}
	return db.Model(models.BAndR{}).Where("status = 1 and id = ?", id).Updates(data).Error
}

func UserReturnControl(db *gorm.DB, id int, userid uint) (err error) {
	return db.Model(models.BAndR{}).Where("status = 2 and id = ? and user_id = ?", id, userid).Updates(models.BAndR{Status: 3, ReturnDate: time.Now().Format("2006-01-02 15:04:05")}).Error
}
