package main

import (
	"fmt"
	"golang-rest-api/api/controller"
	"golang-rest-api/api/route"
	"golang-rest-api/bootstrap"
	"golang-rest-api/repository"
	"golang-rest-api/usecase"

	"github.com/gin-gonic/gin"
)

func main() {

	db := bootstrap.ConnectDB()

	router := gin.Default()

	// wiring

	userRepo := repository.NewUserRepository(db)

	signupUsecase := usecase.NewSignupUsecase(userRepo)
	loginUsecase := usecase.NewLoginUsecase(userRepo)
	signupController := controller.NewSignupController(signupUsecase)
	loginControlelr := controller.NewLoginController(loginUsecase)

	route.NewSignupRouter(router, signupController)
	route.NewLoginRouter(router, loginControlelr)

	// protected := router.Group("/")
	// protected.Use(middleware.JWTAuthMiddleware())

	// protected.GET("/profile", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{"message": "protected route"})

	// })

	fmt.Println("Server running on port: 8080")
	router.Run(":8080")

}
