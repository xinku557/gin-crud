package services

import (
	"gin-crud/models"

	"gorm.io/gorm"
)

func GetAllBook(db *gorm.DB, data *[]models.Book) (err error) {
	return db.Find(&data).Order("id desc").Error
}

func GetBook(db *gorm.DB, data *models.Book, id int) (err error) {
	return db.Preload("Author").Preload("Category").First(&data, id).Error
}

func SearchBook(db *gorm.DB, data *[]models.Book, val string) (err error) {
	return db.Where("title Like ?", "%"+val+"%").Error
}

func CreateBook(db *gorm.DB, data *models.Book) (err error) {
	return db.Create(&data).Error
}

func UpdateBook(db *gorm.DB, data *models.Book, id int) (err error) {
	return db.Model(&models.Book{}).Where("id = ?", id).Save(&data).Error
}

func DeleteBook(db *gorm.DB, id int) (err error) {

	return db.Where("id = ?", id).Delete(&models.Book{}).Error
}
