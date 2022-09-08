package controllers

import (
	"gin-crud/models"
	"gin-crud/services"
	"gin-crud/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (serve *Serve) CreateBookRoute(ctx *gin.Context) {
	tokenData, _ := utils.ExtractTokenID(ctx)

	var json models.Book
	if err := ctx.ShouldBindJSON(&json); err != nil {
		utils.ErrorMessage(ctx, http.StatusBadRequest, "Validation Error", err.Error())
		return
	}
	if services.CheckAuthor(serve.DB, json.AuthorID) == 0 {
		utils.ErrorMessage(ctx, http.StatusBadRequest, "Validation Error", []string{"The author with this id not found."})
		return
	}
	if services.CheckCategory(serve.DB, json.CategoryID) == 0 {
		utils.ErrorMessage(ctx, http.StatusBadRequest, "Validation Error", []string{"The category with this id not found."})
		return
	}
	json.AdminID = tokenData["id"]
	err := services.CreateBook(serve.DB, &json)
	if err != nil {
		utils.ErrorMessage(ctx, http.StatusBadGateway, "Internal Server Error", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusCreated, "Created", nil)

}

func (serve *Serve) GetAllBookRoute(ctx *gin.Context) {
	var data []models.Book
	err := services.GetAllBook(serve.DB, &data)
	if err != nil {
		utils.ErrorMessage(ctx, http.StatusBadGateway, "Internal Server Error", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusOK, "Success", data)

}

func (serve *Serve) GetBookRoute(ctx *gin.Context) {
	var id, _ = strconv.Atoi(ctx.Param("id"))
	var data models.Book
	err := services.GetBook(serve.DB, &data, id)
	if err != nil {
		utils.ErrorMessage(ctx, http.StatusNotFound, "No book found", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusOK, "Success", data)
}

func (serve *Serve) SearchBookRoute(ctx *gin.Context) {
	var val = ctx.Param("query")
	var data []models.Book
	err := services.SearchBook(serve.DB, &data, val)
	if err != nil {
		utils.ErrorMessage(ctx, http.StatusNotFound, "No book found", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusOK, "Success", data)
}

func (serve *Serve) UpdateBookRoute(ctx *gin.Context) {
	tokenData, _ := utils.ExtractTokenID(ctx)

	var id, _ = strconv.Atoi(ctx.Param("id"))
	var json models.Book
	if err := ctx.ShouldBindJSON(&json); err != nil {
		utils.ErrorMessage(ctx, http.StatusBadRequest, "Validation Error", err.Error())
		return
	}
	if services.CheckAuthor(serve.DB, json.AuthorID) == 0 {
		utils.ErrorMessage(ctx, http.StatusBadRequest, "Validation Error", []string{"The author with this id not found."})
		return
	}
	if services.CheckCategory(serve.DB, json.CategoryID) == 0 {
		utils.ErrorMessage(ctx, http.StatusBadRequest, "Validation Error", []string{"The category with this id not found."})
		return
	}
	json.AdminID = tokenData["id"]
	err := services.UpdateBook(serve.DB, &json, id)
	if err != nil {
		utils.ErrorMessage(ctx, http.StatusBadGateway, "Internal Server Error", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusAccepted, "Success", json)
}

func (serve *Serve) DeleteBookRoute(ctx *gin.Context) {
	var id, _ = strconv.Atoi(ctx.Param("id"))
	err := services.DeleteBook(serve.DB, id)
	if err != nil {
		utils.ErrorMessage(ctx, http.StatusNotFound, "No book found", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusNoContent, "Success", nil)
}
