package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"

	models "github.com/Aibier/go-aml-service/internal/pkg/models/tasks"
	"github.com/Aibier/go-aml-service/internal/pkg/persistence"
	requestErr "github.com/Aibier/go-aml-service/internal/util/http-err"
)

// GetTaskByID godoc
func GetTaskByID(c *gin.Context) {
	s := persistence.GetTaskRepository()
	id := c.Param("id")
	if task, err := s.Get(id); err != nil {
		requestErr.NewError(c, http.StatusNotFound, errors.New("task not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, task)
	}
}

// GetTasks godoc
// @Security Authorization Token
func GetTasks(c *gin.Context) {
	s := persistence.GetTaskRepository()
	var q models.Task
	_ = c.Bind(&q)
	if tasks, err := s.Query(&q); err != nil {
		requestErr.NewError(c, http.StatusNotFound, errors.New("tasks not found"))
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, tasks)
	}
}

// CreateTask ...
func CreateTask(c *gin.Context) {
	s := persistence.GetTaskRepository()
	var taskInput models.Task
	_ = c.BindJSON(&taskInput)
	if err := s.Add(&taskInput); err != nil {
		requestErr.NewError(c, http.StatusBadRequest, err)
		log.Println(err)
	} else {
		c.JSON(http.StatusCreated, taskInput)
	}
}

// UpdateTask ...
func UpdateTask(c *gin.Context) {
	s := persistence.GetTaskRepository()
	id := c.Params.ByName("id")
	var taskInput models.Task
	_ = c.BindJSON(&taskInput)
	if _, err := s.Get(id); err != nil {
		requestErr.NewError(c, http.StatusNotFound, errors.New("task not found"))
	} else {
		if err := s.Update(&taskInput); err != nil {
			requestErr.NewError(c, http.StatusNotFound, err)
		} else {
			c.JSON(http.StatusOK, taskInput)
		}
	}
}

// DeleteTask ...
func DeleteTask(c *gin.Context) {
	s := persistence.GetTaskRepository()
	id := c.Params.ByName("id")
	var taskInput models.Task
	_ = c.BindJSON(&taskInput)
	if task, err := s.Get(id); err != nil {
		requestErr.NewError(c, http.StatusNotFound, errors.New("task not found"))
	} else {
		if err := s.Delete(task); err != nil {
			requestErr.NewError(c, http.StatusNotFound, err)
		} else {
			c.JSON(http.StatusNoContent, "")
		}
	}
}
