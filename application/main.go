package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"main/handler" // ここにhandlerディレクトリのパスを記述します
)

type Todo struct {
	ID   int
	Task string
}

type TodoRepository interface {
	GetTodos() ([]Todo, error)
}

type MockTodoRepository struct{}

func (r MockTodoRepository) GetTodos() ([]Todo, error) {
	todos := []Todo{
		{ID: 1, Task: "Task 1"},
		{ID: 2, Task: "Task 2"},
		{ID: 3, Task: "Task 3"},
	}
	return todos, nil
}

func main() {
	r := gin.Default()

	r.GET("/todos", func(c *gin.Context) {
		repo := MockTodoRepository{}
		todos, err := repo.GetTodos()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, todos)
	})

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
