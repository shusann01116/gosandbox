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

func (c *Controller) GetTodos(ctx *gin.Context) {
	todos, err := c.todoService.GetTodos(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, todos)
}
