package controllers

import (
	"gin-crud/models"
	"gin-crud/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"gin-crud/utils"
)

func (serve *Serve) CreateAuthorRoute(ctx *gin.Context) {
	tokenData, _ := utils.ExtractTokenID(ctx)
	var json models.Author
	if err := ctx.ShouldBindJSON(&json); err != nil {
		utils.ErrorMessage(ctx, http.StatusBadRequest, "Validation Error", err.Error())
		return
	}
	json.AdminID = tokenData["id"]
	err := services.CreateAuthor(serve.DB, &json)
	if err != nil {
		utils.ErrorMessage(ctx, http.StatusBadGateway, "Internal Server Error", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusCreated, "Created", nil)
}

func (serve *Serve) GetAllAuthorRoute(ctx *gin.Context) {
	var data []models.Author
	err := services.GetAllAuthor(serve.DB, &data)
	if err != nil {
		utils.ErrorMessage(ctx, http.StatusBadGateway, "Internal Server Error", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusOK, "Success", data)

}

func (serve *Serve) GetAuthorRoute(ctx *gin.Context) {
	var id, _ = strconv.Atoi(ctx.Param("id"))
	var data models.Author
	err := services.GetAuthor(serve.DB, &data, id)
	if err != nil {
		utils.ErrorMessage(ctx, http.StatusNotFound, "No author found", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusOK, "Success", data)
}

func (serve *Serve) SearchAuthorRoute(ctx *gin.Context) {
	var val = ctx.Param("query")
	var data []models.Author
	err := services.SearchAuthor(serve.DB, &data, val)
	if err != nil {
		utils.ErrorMessage(ctx, http.StatusNotFound, "No author found", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusOK, "Success", data)
}

func (serve *Serve) UpdateAuthorRoute(ctx *gin.Context) {
	tokenData, _ := utils.ExtractTokenID(ctx)

	var id, _ = strconv.Atoi(ctx.Param("id"))
	var json models.Author
	if err := ctx.ShouldBindJSON(&json); err != nil {
		utils.ErrorMessage(ctx, http.StatusBadRequest, "Validation Error", err.Error())
		return
	}
	json.AdminID = tokenData["id"]
	err := services.UpdateAuthor(serve.DB, &json, id)
	if err != nil {
		utils.ErrorMessage(ctx, http.StatusBadGateway, "Internal Server Error", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusAccepted, "Success", json)
}

func (serve *Serve) DeleteAuthorRoute(ctx *gin.Context) {
	var id, _ = strconv.Atoi(ctx.Param("id"))
	err := services.DeleteAuthor(serve.DB, id)
	if err != nil {
		utils.ErrorMessage(ctx, http.StatusNotFound, "No author found", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusNoContent, "Success", nil)
}
