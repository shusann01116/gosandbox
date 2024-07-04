package todo

import "fmt"

type Service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s *Service) AddTodo(todo Item) error {
	if err := todo.IsValid(); err != nil {
		return fmt.Errorf("failed to add todo: %w", err)
	}

	s.repo.AddTodos(todo)
	return nil
}

func (s *Service) GetTodos() ([]Item, error) {
	todos, err := s.repo.ListTodos()
	if err != nil {
		return nil, fmt.Errorf("failed to get Todos: %w", err)
	}
	return todos, nil
}
