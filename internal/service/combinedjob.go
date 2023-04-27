package service

import (
	"context"

	"github.com/grpr-job-api/internal/dao"
)

func (s *Service) GetCombinedJobs(ctx context.Context) ([]*dao.CombinedJob, error) {
	return s.repo.GetCombinedJobs(ctx)
}