package repository

import "github.com/gitSanje/go-taskManiac/internal/server"

type Repositories struct{}

func NewRepositories(s *server.Server) *Repositories {
	return &Repositories{}
}
