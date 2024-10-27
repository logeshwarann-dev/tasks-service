package handlers

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/tasks", CreateTask)
	router.GET("/tasks", GetTasks)
	router.DELETE("/tasks/:id", DeleteTask)
}
