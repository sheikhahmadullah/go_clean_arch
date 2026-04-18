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

//

func (c *SignupController) Signup(ctx *gin.Context) {
	var user domain.User

	// check if user sends bad request
	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.usecase.Signup(&user)
	// check if we failed to create a user in signup usecase step
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created"})

}
