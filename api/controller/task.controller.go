package controller

import (
	"golang-rest-api/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	usecase *usecase.TaskUsecase
}

func NewTaskController(u *usecase.TaskUsecase) *TaskController {
	return &TaskController{u}
}

func (c *TaskController) CreateTask(ctx *gin.Context) {
	var body struct {
		Title string `json:"title"`
	}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := ctx.GetUint("user_id")

	err := c.usecase.CreateTask(body.Title, userId)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Task created successfully"})
}

func (c *TaskController) GetTasks(ctx *gin.Context) {
	userID := ctx.GetUint("user_id")

	tasks, err := c.usecase.GetTasks(userID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, tasks)
}
