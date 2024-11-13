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
	c.setupRouter()

	return c
}

func (c *Controller) setupRouter() {
	c.router.Use(gin.Logger())
	c.router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})
}

func (c *Controller) Run() error {
	err := c.router.Run()
	if err != nil {
		panic(err)
	}
	return nil
}
