package adapter

import (
	"blazperic/vjezba2/config"
	"blazperic/vjezba2/service"
	"blazperic/vjezba2/shared"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RestAdapter struct {
	router  *gin.Engine
	service *service.NoteService
}

type NoteDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func NewRestAdapter(service *service.NoteService) RestAdapter {
	c := RestAdapter{
		router:  gin.Default(),
		service: service,
	}

	c.setupRoutes()

	return c
}

func (c *RestAdapter) setupRoutes() {
	c.router.Use(gin.Logger())

	c.router.GET("/", c.health)

	v1 := c.router.Group("/api/v1")
	{
		v1.GET("/notes", c.getNotes)
		v1.GET("/notes/:id", c.getNote)
		v1.POST("/notes", c.newNote)
	}

}

func (c *RestAdapter) Run() error {
	err := c.router.Run(config.SERVER_PORT)
	if err != nil {
		return &shared.ServerError{Message: "Dogodila se greška prilikom pokretanja servera: " + err.Error()}
	}
	return nil
}

func (c *RestAdapter) health(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "OK",
	})
}

func (c *RestAdapter) getNotes(ctx *gin.Context) {
	notes, err := c.service.GetAllNotes()
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, notes)
}

func (c *RestAdapter) getNote(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "Neispravan ID",
		})
		return
	}
	note, err := c.service.GetNote(id)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, note)
}

func (c *RestAdapter) newNote(ctx *gin.Context) {
	var note NoteDTO

	if err := ctx.ShouldBindJSON(&note); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.service.NewNote(note.Title, note.Description)

	if err != nil {
		_ = c.service.NewNote(note.Title, note.Description)
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
