package infrastructure

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	router *gin.Engine
}

func NewRestController() *Controller {
	c := &Controller{
		router: gin.Default(),
	}

	c.router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	return c
}

func (c *Controller) Run() error {
	err := c.router.Run()
	if err != nil {
		panic(err)
	}
	return nil
}
