package service

import (
	"github.com/gitSanje/go-taskManiac/internal/lib/job"
	"github.com/gitSanje/go-taskManiac/internal/repository"
	"github.com/gitSanje/go-taskManiac/internal/server"
)

type Services struct {
	Auth *AuthService
	Job  *job.JobService
}

func NewServices(s *server.Server, repos *repository.Repositories) (*Services, error) {
	authService := NewAuthService(s)

	return &Services{
		Job:  s.Job,
		Auth: authService,
	}, nil
}
