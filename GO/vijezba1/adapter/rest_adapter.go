package adapter

import (
	"blazperic/santa/config"
	"blazperic/santa/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RestAdapter struct {
	router       *gin.Engine
	santaService *service.SantaService
}

func NewRestController(santaService *service.SantaService) RestAdapter {
	c := RestAdapter{
		router:       gin.Default(),
		santaService: santaService,
	}

	c.setupRoutes()
	return c
}
func (c *RestAdapter) Run() {
	err := c.router.Run(config.SERVER_PORT)
	if err != nil {
		panic(err)
	}
}

func (c *RestAdapter) setupRoutes() {
	c.router.Use(gin.Logger())

	c.router.GET("/", c.health)

	v1 := c.router.Group("api/v1")
	{
		v1.GET("/allSantas", c.getAllSantas)
		v1.GET("/santa", c.choiceSanta)
	}
}

func (c *RestAdapter) health(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "blazperic/vijezba1 is running...",
	})
}

func (c *RestAdapter) getAllSantas(ctx *gin.Context) {
	santa, err := c.santaService.GetAllSantas()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("error while getting tasks: %v", err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, santa)
}

func (c *RestAdapter) choiceSanta(ctx *gin.Context) {
	santa, err := c.santaService.ChoiceSanta()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Errorf("error while getting tasks: %v", err.Error()),
		})
		return
	}
	ctx.JSON(http.StatusOK, santa)

}
