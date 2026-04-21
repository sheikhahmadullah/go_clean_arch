package controller

import (
	"golang-rest-api/usecase"

	"github.com/gin-gonic/gin"
)

type RefreshController struct {
	usecase *usecase.RefreshUsecase
}

func NewRefreshUsecase(u *usecase.RefreshUsecase) *RefreshController {
	return &RefreshController{u}
}

func (c *RefreshController) Refresh(ctx *gin.Context) {
	var body struct {
		Token string `json:"refresh_token"`
	}

	ctx.BindJSON(&body)

	newToken, err := c.usecase.Refresh(body.Token)

	if err != nil {
		ctx.JSON(401, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"access_token": newToken})
}
