package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todo-cc/config"
)

type Controller struct {
	router *gin.Engine
}

func NewRestController() Controller {
	c := Controller{
		router: gin.Default(),
	}

	c.setupRoutes()

	return c
}

func (c *Controller) setupRoutes() {
	c.router.Use(gin.Logger())

	c.router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello world!",
		})
	})
}

func (c *Controller) Run() {
	err := c.router.Run(config.SERVER_PORT)
	if err != nil {
		panic(err)
	}
}
