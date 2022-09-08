package controllers

import (
	"gin-crud/models"
	"gin-crud/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-crud/utils"
)

func (serve *Serve) CreateCategoryRoute(ctx *gin.Context) {
	tokenData, _ := utils.ExtractTokenID(ctx)
	var json models.Category
	if err := ctx.ShouldBindJSON(&json); err != nil {
		utils.ErrorMessage(ctx, http.StatusBadRequest, "Validation Error", err.Error())
		return
	}
	json.AdminID = tokenData["id"]
	err := services.CreateCategory(serve.DB, &json)
	if err != nil {
		utils.ErrorMessage(ctx, http.StatusBadGateway, "Internal Server Error", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusCreated, "Created", nil)

}

func (serve *Serve) GetAllCategoryRoute(ctx *gin.Context) {
	var data []models.Category
	err := services.GetAllCategory(serve.DB, &data)
	if err != nil {
		utils.ErrorMessage(ctx, http.StatusBadGateway, "Internal Server Error", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusOK, "Success", data)

}

func (serve *Serve) GetCategoryRoute(ctx *gin.Context) {
	var id, _ = strconv.Atoi(ctx.Param("id"))
	var data models.Category
	err := services.GetCategory(serve.DB, &data, id)
	if err != nil {
		utils.ErrorMessage(ctx, http.StatusNotFound, "No Category found", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusOK, "Success", data)
}

func (serve *Serve) SearchCategoryRoute(ctx *gin.Context) {
	var val = ctx.Param("query")
	var data []models.Category
	err := services.SearchCategory(serve.DB, &data, val)
	if err != nil {
		utils.ErrorMessage(ctx, http.StatusNotFound, "No category found", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusOK, "Success", data)
}

func (serve *Serve) UpdateCategoryRoute(ctx *gin.Context) {
	tokenData, _ := utils.ExtractTokenID(ctx)

	var id, _ = strconv.Atoi(ctx.Param("id"))
	var json models.Category
	if err := ctx.ShouldBindJSON(&json); err != nil {
		utils.ErrorMessage(ctx, http.StatusBadRequest, "Validation Error", err.Error())
		return
	}
	json.AdminID = tokenData["id"]
	err := services.UpdateCategory(serve.DB, &json, id)
	if err != nil {
		utils.ErrorMessage(ctx, http.StatusBadGateway, "Internal Server Error", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusAccepted, "Success", json)
}

func (serve *Serve) DeleteCategoryRoute(ctx *gin.Context) {
	var id, _ = strconv.Atoi(ctx.Param("id"))
	err := services.DeleteCategory(serve.DB, id)
	if err != nil {
		utils.ErrorMessage(ctx, http.StatusNotFound, "No book found", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusNoContent, "Success", nil)
}
