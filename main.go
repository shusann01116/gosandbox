package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func main() {
	ac, err := InitializeTodoController()
	if err != nil {
		panic("failed to init TodoController")
	}
	router := gin.Default()
	router.GET("/todos", ac.GetTodos)

	router.Run("localhost:8080")
}

type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func NewTodo(title string) Todo {
	return Todo{
		ID:    uuid.New().String(),
		Title: title,
	}
}

func (t *Todo) IsValid() error {
	if err := uuid.Validate(t.ID); err != nil {
		return fmt.Errorf("invalid todo ID: %w", err)
	}
	return nil
}

type TodoController struct {
	todoService TodoService
}

func NewTodoController(svc TodoService) TodoController {
	return TodoController{
		todoService: svc,
	}
}

func (tc *TodoController) GetTodos(c *gin.Context) {
	todos, err := tc.todoService.GetTodos()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, todos)
}

type TodoService struct {
	repo TodoRepository
}

func NewTodoService(repo TodoRepository) TodoService {
	return TodoService{
		repo: repo,
	}
}

func (s *TodoService) AddTodo(todo Todo) error {
	if err := todo.IsValid(); err != nil {
		return fmt.Errorf("failed to add todo: %w", err)
	}

	s.repo.AddTodos(todo)
	return nil
}

func (s *TodoService) GetTodos() ([]Todo, error) {
	todos, err := s.repo.ListTodos()
	if err != nil {
		return nil, fmt.Errorf("failed to get Todos: %w", err)
	}
	return todos, nil
}

type TodoRepository interface {
	ListTodos() ([]Todo, error)
	AddTodos(todos ...Todo) error
}

type inMemoryTodoRepo struct {
	todos []Todo
}

func NewInMemoryTodoRepo() (TodoRepository, error) {
	todos := initTodo()
	return &inMemoryTodoRepo{
		todos: todos,
	}, nil
}

func initTodo() []Todo {
	var todos []Todo
	for i := 0; i < 20; i++ {
		todos = append(todos, NewTodo(fmt.Sprintf("My Todo %v", i)))
	}
	return todos
}

func (i *inMemoryTodoRepo) ListTodos() ([]Todo, error) {
	return i.todos, nil
}

func (i *inMemoryTodoRepo) AddTodos(todos ...Todo) error {
	i.todos = append(i.todos, todos...)
	return nil
}
