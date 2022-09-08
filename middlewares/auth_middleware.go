package middlewares

import (
	"gin-crud/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NormalUserAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/json")
		if utils.TokenValid(ctx) != nil {
			utils.ErrorMessage(ctx, http.StatusUnauthorized, "Unauthorized", nil)
			ctx.Abort()
			return
		}
	}
}

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		data, _ := utils.ExtractTokenID(ctx)
		if data["level"] == 1 {
			utils.ErrorMessage(ctx, http.StatusUnauthorized, "Unauthorized", nil)
			ctx.Abort()
			return
		}
	}
}
