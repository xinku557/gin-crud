package services

import (
	"gin-crud/models"

	"gorm.io/gorm"
)

func Register(db *gorm.DB, data *models.User) (err error) {
	return db.Create(&data).Error
}

func Login(db *gorm.DB, data *models.User, login models.Login) (err error) {
	return db.Where("username = ? OR email = ?", login.User, login.User).First(&data).Error
}

func Me(db *gorm.DB, data *models.User, id int) (err error) {
	return db.Preload("Authors").Preload("Books").Preload("Categories").Where("id = ?", id).First(&data).Error
}

func CheckDuplicate(db *gorm.DB, typ string, val string) int64 {
	return db.Where(typ+" = ?", val).Find(&models.User{}).RowsAffected
}
