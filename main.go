package main

import (
	"gin-crud/controllers"
	"gin-crud/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	var controllers controllers.Serve
	controllers.Init()
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Server is running.....",
		})
	})
	v1 := router.Group("/api/v1")
	{

		v1.POST("/auth/login", controllers.LoginRoute)
		v1.POST("/auth/register", controllers.RegisterRoute)
		v1.POST("/auth/add-user/:type", middlewares.NormalUserAuthMiddleware(), middlewares.AuthMiddleware(), controllers.AddUserRoute)
		v1.GET("/profile", middlewares.NormalUserAuthMiddleware(), controllers.ProfileRoute)
		v1.GET("/borrow/:id/:status", middlewares.NormalUserAuthMiddleware(), middlewares.AuthMiddleware(), controllers.UserBorrowControlRoute)
		v1.GET("/borrow-return-list/:query", middlewares.NormalUserAuthMiddleware(), middlewares.AuthMiddleware(), controllers.UserBRListRoute)
		v1.GET("/authors", controllers.GetAllAuthorRoute)
		v1.POST("/authors", middlewares.AuthMiddleware(), controllers.CreateAuthorRoute)
		author := v1.Group("/author/:id")
		{
			author.GET("", controllers.GetAuthorRoute)
			author.GET("/search/:query", controllers.SearchAuthorRoute)
			author.PUT("", middlewares.AuthMiddleware(), controllers.UpdateAuthorRoute)
			author.DELETE("", middlewares.AuthMiddleware(), controllers.DeleteAuthorRoute)
		}
		author.Use(middlewares.NormalUserAuthMiddleware())
		v1.GET("/categories", controllers.GetAllCategoryRoute)
		v1.POST("/categories", middlewares.AuthMiddleware(), controllers.CreateCategoryRoute)
		category := v1.Group("/category/:id")
		{
			category.GET("", controllers.GetCategoryRoute)
			category.GET("/search/:query", controllers.SearchCategoryRoute)
			category.PUT("", middlewares.AuthMiddleware(), controllers.UpdateCategoryRoute)
			category.DELETE("", middlewares.AuthMiddleware(), controllers.DeleteCategoryRoute)
		}
		category.Use(middlewares.NormalUserAuthMiddleware())

		v1.GET("/books", controllers.GetAllBookRoute)
		v1.POST("/books", middlewares.AuthMiddleware(), controllers.CreateBookRoute)
		book := v1.Group("/book/:id")
		{
			book.GET("", controllers.GetBookRoute)
			book.GET("/borrow", controllers.UserBorrowRoute)

			book.GET("/search/:query", controllers.SearchBookRoute)
			book.PUT("", middlewares.AuthMiddleware(), controllers.UpdateBookRoute)
			book.DELETE("", middlewares.AuthMiddleware(), controllers.DeleteBookRoute)
		}
		book.Use(middlewares.NormalUserAuthMiddleware())
	}
	router.Run(":8888")
}
