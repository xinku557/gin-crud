package utils

import (
	"github.com/gin-gonic/gin"
)

func SuccessMessage(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.JSON(status, gin.H{"status": status, "message": message, "data": data, "error": nil})
}

func ErrorMessage(ctx *gin.Context, status int, message string, err interface{}) {
	ctx.JSON(status, gin.H{"status": status, "message": message, "data": nil, "error": err})
}
