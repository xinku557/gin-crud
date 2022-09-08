package controllers

import (
	"gin-crud/models"
	"gin-crud/services"
	"gin-crud/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (serve *Serve) RegisterRoute(ctx *gin.Context) {
	var json models.User
	if err := ctx.ShouldBindJSON(&json); err != nil {
		utils.ErrorMessage(ctx, http.StatusBadRequest, "Validation Error", err.Error())
		return
	}
	if c := services.CheckDuplicate(serve.DB, "email", json.Email); c != 0 {
		utils.ErrorMessage(ctx, http.StatusBadRequest, "Validation Error", []string{"Email already exits."})
		return
	}
	if c := services.CheckDuplicate(serve.DB, "username", json.Username); c != 0 {
		utils.ErrorMessage(ctx, http.StatusBadRequest, "Validation Error", []string{"Username already exits."})
		return
	}
	if err := services.Register(serve.DB, &json); err != nil {
		utils.ErrorMessage(ctx, http.StatusBadGateway, "Internal Server Error", nil)
		return
	}
	utils.SuccessMessage(ctx, http.StatusCreated, "Created", nil)
}

func (serve *Serve) AddUserRoute(ctx *gin.Context) {
	var id, _ = strconv.Atoi(ctx.Param("type"))
	var json models.User
	if err := ctx.ShouldBindJSON(&json); err != nil {
		utils.ErrorMessage(ctx, http.StatusBadRequest, "Validation Error", err.Error())
		return
	}
	switch id {
	case 1, 2:
		json.Type = uint(id)
		if c := services.CheckDuplicate(serve.DB, "email", json.Email); c != 0 {
			utils.ErrorMessage(ctx, http.StatusBadRequest, "Validation Error", []string{"Email already exits."})
			return
		}
		if c := services.CheckDuplicate(serve.DB, "username", json.Username); c != 0 {
			utils.ErrorMessage(ctx, http.StatusBadRequest, "Validation Error", []string{"Username already exits."})
			return
		}
		if err := services.Register(serve.DB, &json); err != nil {
			utils.ErrorMessage(ctx, http.StatusBadGateway, "Internal Server Error", nil)
			return
		}
		utils.SuccessMessage(ctx, http.StatusCreated, "Created", nil)
		return
	default:
		utils.ErrorMessage(ctx, http.StatusBadRequest, "Validation Error", []string{"Please enter validate type."})
		return
	}
}

func (serve *Serve) LoginRoute(ctx *gin.Context) {
	var json models.Login
	var data models.User
	if err := ctx.ShouldBindJSON(&json); err != nil {
		utils.ErrorMessage(ctx, http.StatusBadRequest, "Validation Error", err.Error())
		return
	}
	if err := services.Login(serve.DB, &data, json); err != nil {
		utils.ErrorMessage(ctx, http.StatusBadRequest, "Login Fail", []string{"Username/Email Not found."})
		return
	}

	if key := bcrypt.CompareHashAndPassword([]byte(data.Password), []byte(json.Password)); key != nil {
		utils.ErrorMessage(ctx, http.StatusBadRequest, "Login Fail", []string{"These credentials do not match our records."})
		return
	}

	var res models.ReturnUser
	res.Email = data.Email
	res.Username = data.Username
	res.CreatedAt = data.CreatedAt
	res.Token, _ = utils.CreateToken(data.ID, data.Type)
	utils.SuccessMessage(ctx, http.StatusOK, "Success", res)
}

func (serve *Serve) ProfileRoute(ctx *gin.Context) {
	tokenData, _ := utils.ExtractTokenID(ctx)
	var user models.User
	if err := services.Me(serve.DB, &user, int(tokenData["id"])); err != nil {
		utils.ErrorMessage(ctx, http.StatusForbidden, "Invalid Token", nil)
		return
	}
	res := models.ReturnUser{Username: user.Username, Email: user.Email, Books: user.Books, Authors: user.Authors, Categories: user.Categories}
	utils.SuccessMessage(ctx, http.StatusOK, "Success", res)

}
