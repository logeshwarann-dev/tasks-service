package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateTask handles creating a new task
func CreateTask(c *gin.Context) {
	var task Task
	if err := c.BindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := GetTaskCollection()

	// Insert the new task into MongoDB
	result, err := collection.InsertOne(context.Background(), task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"task_id": result.InsertedID})
}

// GetTasks fetches all tasks for the current user
func GetTasks(c *gin.Context) {
	var tasks []Task
	collection := GetTaskCollection()

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var task Task
		if err := cursor.Decode(&task); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		tasks = append(tasks, task)
	}

	c.JSON(http.StatusOK, gin.H{"tasks": tasks})
}

// DeleteTask handles deleting a task by ID
func DeleteTask(c *gin.Context) {
	taskID := c.Param("id")
	objID, _ := primitive.ObjectIDFromHex(taskID)

	collection := GetTaskCollection()

	_, err := collection.DeleteOne(context.Background(), bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
}
