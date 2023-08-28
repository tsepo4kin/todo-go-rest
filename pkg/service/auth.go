package service

import (
	"crypto/sha1"
	"fmt"

	"github.com/tsepo4kin/todo-go-rest/pkg/repository"
	todo "github.com/tsepo4kin/todo-go-rest/types"
)

const salt = ";lkajsdghui123ijvasduiv98"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
