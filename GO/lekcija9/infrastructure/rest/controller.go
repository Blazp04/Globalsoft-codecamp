package rest

import (
	"blazperic/lekcija9/config"
	"blazperic/lekcija9/core/port"
	"blazperic/lekcija9/infrastructure/persistence"
	"blazperic/lekcija9/persistence/sqlite"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

	c.router.GET("/tasks/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid id",
			})
			return
		}
		db, _ := sqlite.NewSqliteDatabase()

		taskPersistence := persistence.NewPersistenceAdapter(db.GetDb())

		task, err := taskPersistence.GetTask(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while getting task",
			})
			return
		}

		c.JSON(http.StatusOK, task)
	})

	c.router.POST("/tasks", func(c *gin.Context) {
		c.Header("Content-Type", "application/json")
		db, _ := sqlite.NewSqliteDatabase()

		taskPersistence := persistence.NewPersistenceAdapter(db.GetDb())

		var taskDTO port.TaskDTO
		err := c.BindJSON(&taskDTO)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Error while binding task",
			})
			return
		}

		err = taskPersistence.NewTask(taskDTO.Title, taskDTO.Description, taskDTO.Deadline, false, false)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while creating task",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Task created successfully",
		})
	})
}

func (c *Controller) Run() {
	err := c.router.Run(config.SERVER_PORT)
	if err != nil {
		panic(err)
	}
}
