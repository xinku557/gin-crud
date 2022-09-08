package services

import (
	"gin-crud/models"

	"gorm.io/gorm"
)

func CheckCategory(db *gorm.DB, id uint) int64 {
	i := db.Where("id = ?", id).Find(&models.Category{}).RowsAffected
	return i
}

func GetAllCategory(db *gorm.DB, data *[]models.Category) (err error) {
	return db.Find(&data).Order("id desc").Error
}

func GetCategory(db *gorm.DB, data *models.Category, id int) (err error) {
	return db.Preload("Books").First(&data, id).Error
}

func SearchCategory(db *gorm.DB, data *[]models.Category, val string) (err error) {
	return db.Where("name Like ?", "%"+val+"%").Error
}

func CreateCategory(db *gorm.DB, data *models.Category) (err error) {
	return db.Create(&data).Error
}

func UpdateCategory(db *gorm.DB, data *models.Category, id int) (err error) {
	return db.Model(&models.Category{}).Where("id = ?", id).Update("name", data.Name).Error
}

func DeleteCategory(db *gorm.DB, id int) (err error) {

	return db.Where("id = ?", id).Delete(&models.Category{}).Error
}
