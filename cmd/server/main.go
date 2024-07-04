package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	tc, err := InitializeTodoController()
	if err != nil {
		panic("failed to init TodoController")
	}
	router := gin.Default()
	router.GET("/todos", tc.GetTodos)

	router.Run("localhost:8080")
}
