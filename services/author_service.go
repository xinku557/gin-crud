package services

import (
	"gin-crud/models"

	"gorm.io/gorm"
)

func CheckAuthor(db *gorm.DB, id uint) int64 {
	i := db.Where("id = ?", id).Find(&models.Author{}).RowsAffected
	return i
}

func GetAllAuthor(db *gorm.DB, data *[]models.Author) (err error) {
	return db.Find(&data).Order("id desc").Error
}

func GetAuthor(db *gorm.DB, data *models.Author, id int) (err error) {
	return db.Preload("Books").First(&data, id).Error
}

func SearchAuthor(db *gorm.DB, data *[]models.Author, val string) (err error) {
	return db.Where("name Like ?", "%"+val+"%").Error
}

func CreateAuthor(db *gorm.DB, data *models.Author) (err error) {
	return db.Create(&data).Error
}

func UpdateAuthor(db *gorm.DB, data *models.Author, id int) (err error) {
	return db.Model(&models.Author{}).Where("id = ?", id).Update("name", data.Name).Error
}

func DeleteAuthor(db *gorm.DB, id int) (err error) {
	return db.Unscoped().Where("id = ?", id).Delete(&models.Author{}).Error
}
