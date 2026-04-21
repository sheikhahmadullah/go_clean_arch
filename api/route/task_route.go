package route

import (
	"golang-rest-api/api/controller"
	"golang-rest-api/api/middleware"

	"github.com/gin-gonic/gin"
)

func NewTaskRouter(r *gin.Engine, c *controller.TaskController) {
	protected := r.Group("/tasks")

	protected.Use(middleware.JWTAuthMiddleware())

	protected.POST("/", c.CreateTask)
	protected.GET("/", c.GetTasks)

}
