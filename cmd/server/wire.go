//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/shusann01116/gosandbox/todo"
)

func InitializeTodoController() (todo.Controller, error) {
	wire.Build(todo.NewController, todo.NewService, todo.NewInMemoryRepo)
	return todo.Controller{}, nil
}
