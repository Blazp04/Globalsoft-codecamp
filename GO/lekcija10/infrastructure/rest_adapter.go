package infrastructure

import (
	"blazperic/lekcija10/config"
	"blazperic/lekcija10/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	router  *gin.Engine
	service service.Todo
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
		v1.POST("/task", c.createNewTask)
	}
}

func (c *Controller) createNewTask(ctx *gin.Context) {
	//TODO Implement
	ctx.JSON(http.StatusOK, gin.H{
		"message": "blazperic/lekcija10 is running...",
	})
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
