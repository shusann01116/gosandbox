package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	todoService Service
}

func NewController(svc Service) Controller {
	return Controller{
		todoService: svc,
	}
}

func (tc *Controller) GetTodos(c *gin.Context) {
	todos, err := tc.todoService.GetTodos()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, todos)
}
