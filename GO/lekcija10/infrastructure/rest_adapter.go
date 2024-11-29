package infrastructure

import (
	"blazperic/lekcija10/config"
	"blazperic/lekcija10/service"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	router  *gin.Engine
	service service.Todo
}
type CreateTaskRequestDTO struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Deadline    time.Time `json:"deadline"`
	Completed   bool      `json:"completed"`
}

func NewRestController(service service.Todo) Controller {
	c := Controller{
		router:  gin.Default(),
		service: service,
	}

	c.setupRoutes()

	return c
}

func (c *Controller) setupRoutes() {
	c.router.Use(gin.Logger())

	c.router.GET("/", c.health)

	v1 := c.router.Group("v1")
	{
		v1.GET("/health", c.health)
		v1.GET("//tasks:id", c.getTask)
		v1.GET("/tasks", c.getAllTasks)
		v1.POST("/tasks", c.createNewTask)
		v1.DELETE("/tasks/:id", c.deleteTask)
		v1.PUT("/tasks/:id", c.completeTask)
	}
}

func (c *Controller) createNewTask(ctx *gin.Context) {
	var taskDTO CreateTaskRequestDTO

	if err := ctx.ShouldBindJSON(&taskDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Errorf("error while binding json: %v", err.Error()),
		})
		return
	}

	err := c.service.CreateNewTask(taskDTO.Title, taskDTO.Description, taskDTO.Deadline, taskDTO.Completed)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("error while saving task: %v", err.Error()),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "success",
	})
}
func (c *Controller) completeTask(ctx *gin.Context) {
	taskId := ctx.Param("id")
	id, err := strconv.Atoi(taskId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id",
		})
		return
	}
	err = c.service.CompleteTask(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("error while completing task: %v", err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (c *Controller) getAllTasks(ctx *gin.Context) {
	tasks, err := c.service.GetAllTasks()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("error while getting tasks: %v", err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (c *Controller) deleteTask(ctx *gin.Context) {
	taskId := ctx.Param("id")
	id, err := strconv.Atoi(taskId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id",
		})
		return
	}
	err = c.service.DeleteTask(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("error while deleting task: %v", err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func (c *Controller) getTask(ctx *gin.Context) {
	taskId := ctx.Param("id")
	id, err := strconv.Atoi(taskId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid id",
		})
		return
	}
	task, err := c.service.GetTask(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("error while getting task: %v", err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (c *Controller) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "blazperic/lekcija10 is running...",
	})
}

func (c *Controller) Run() {
	err := c.router.Run(config.SERVER_PORT)
	if err != nil {
		panic(err)
	}
}
