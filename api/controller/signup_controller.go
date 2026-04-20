package controller

import (
	"golang-rest-api/domain"
	"golang-rest-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	usecase *usecase.SignupUsecase
}

func NewSignupController(u *usecase.SignupUsecase) *SignupController {
	return &SignupController{u}
}

// the above part is dependency injection

func (c *SignupController) Signup(ctx *gin.Context) {
	var req domain.SignupRequest // Use the Request DTO

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Map the request data to the User model
	user := domain.User{
		Email:    req.Email,
		Password: req.Password, // Now the password is NOT empty!
	}

	err := c.usecase.Signup(&user)
	if err != nil {
		// Tip: If the error is "Email already exists",
		// you might want to return http.StatusConflict (409)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created"})
}
