//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitializeTodoController() (TodoController, error) {
	wire.Build(NewTodoController, NewTodoService, NewInMemoryTodoRepo)
	return TodoController{}, nil
}
