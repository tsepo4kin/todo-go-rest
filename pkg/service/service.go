package service

import (
	"github.com/tsepo4kin/todo-go-rest/pkg/repository"
	"github.com/tsepo4kin/todo-go-rest/types"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}