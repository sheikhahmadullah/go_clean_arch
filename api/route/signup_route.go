package route

import (
	"golang-rest-api/api/controller"

	"github.com/gin-gonic/gin"
)

func NewSignupRouter(r *gin.Engine, c *controller.SignupController) {
	r.POST("/signup", c.Signup)

}
