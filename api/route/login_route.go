package route

import (
	"golang-rest-api/api/controller"

	"github.com/gin-gonic/gin"
)

func NewLoginRouter(r *gin.Engine, c *controller.LoginController) {
	r.POST("/login", c.Login)
}
