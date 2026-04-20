package controller

import (
	"golang-rest-api/domain"
	"golang-rest-api/internal/tokenutil"
	"golang-rest-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	usecase *usecase.LoginUsecase
}

func NewLoginController(u *usecase.LoginUsecase) *LoginController {
	return &LoginController{u}
}

func (c *LoginController) Login(ctx *gin.Context) {
	var req domain.LoginRequest

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := c.usecase.Login(&req)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	token, _ := tokenutil.GenerateToken(user.Id)

	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})

}
