package service

import (
	"github.com/grpr-job-api/internal/repo"
)

type Service struct {
	repo     *repo.Repo
}

func New(r *repo.Repo) *Service {
	return &Service{r}
}
