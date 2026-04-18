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

	signupController := controller.NewSignupController(signupUsecase)

	route.NewSignupRouter(router, signupController)

	router.Run(":8080")

	fmt.Println("Server running on port: 8080")

}
