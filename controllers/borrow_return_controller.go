package controllers

import (
	"gin-crud/models"
	"gin-crud/services"
	"gin-crud/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (serve *Serve) UserBorrowRoute(ctx *gin.Context) {
	var id, _ = strconv.Atoi(ctx.Param("id"))
	tokenData, _ := utils.ExtractTokenID(ctx)
	if services.CheckActiveStatus(serve.DB, int(tokenData["id"])) {
		utils.ErrorMessage(ctx, http.StatusBadRequest, "Your borrow permission is on limit", nil)
		return
	}
	err := services.UserBorrow(serve.DB, &models.BAndR{UserID: tokenData["id"], BookID: uint(id), Status: 1})
	if err != nil {
		utils.ErrorMessage(ctx, http.StatusNotFound, "No book found", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusAccepted, "Please wait admin accept.", nil)
}

func (serve *Serve) UserBRListRoute(ctx *gin.Context) {
	var query = ctx.Param("query")
	var err error
	var data []models.BAndR
	if query == "all" {
		err = services.BRList(serve.DB, &data, 0)
	} else if query == "allow" {
		err = services.BRList(serve.DB, &data, 2)
	} else if query == "pending" {
		err = services.BRList(serve.DB, &data, 1)
	} else {
		err = services.BRList(serve.DB, &data, 3)
	}

	if err != nil {
		utils.ErrorMessage(ctx, http.StatusNotFound, "Not found", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusOK, "Success", data)

}

func (serve *Serve) UserBorrowControlRoute(ctx *gin.Context) {
	var id, _ = strconv.Atoi(ctx.Param("id"))
	var i int
	var status = ctx.Param("status")
	tokenData, _ := utils.ExtractTokenID(ctx)
	if status == "allow" {
		i = 2
	} else {
		i = 3
	}
	err := services.UserBorrowControl(serve.DB, id, i, tokenData["id"])
	if err != nil {
		utils.ErrorMessage(ctx, http.StatusNotFound, "No data found", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusAccepted, "Success", nil)
}

func (serve *Serve) UserReturnRoute(ctx *gin.Context) {
	var id, _ = strconv.Atoi(ctx.Param("id"))
	tokenData, _ := utils.ExtractTokenID(ctx)
	err := services.UserReturnControl(serve.DB, id, tokenData["id"])
	if err != nil {
		utils.ErrorMessage(ctx, http.StatusNotFound, "No data found", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusAccepted, "Success", nil)
}
