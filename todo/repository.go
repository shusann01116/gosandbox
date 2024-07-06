package todo

import (
	"context"
	"fmt"
	"sync"
)

type Repository interface {
	ListTodos(ctx context.Context) ([]Item, error)
	AddTodos(todos ...Item) error
}

type inMemoryTodoRepo struct {
	mu    sync.Mutex
	todos []Item
}

func NewInMemoryRepo() (Repository, error) {
	todos := initTodo()
	return &inMemoryTodoRepo{
		todos: todos,
	}, nil
}

func initTodo() []Item {
	var todos []Item
	for i := 0; i < 20; i++ {
		todos = append(todos, New(fmt.Sprintf("My Todo %v", i)))
	}
	return todos
}

func (i *inMemoryTodoRepo) ListTodos(ctx context.Context) ([]Item, error) {
	return i.todos, nil
}

func (i *inMemoryTodoRepo) AddTodos(todos ...Item) error {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.todos = append(i.todos, todos...)
	return nil
}
