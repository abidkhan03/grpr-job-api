package service

import (
	"context"

	"github.com/grpr-job-api/internal/dao"
)

func (s *Service) GetAllJobs(ctx context.Context) ([]*dao.FetcherJob, error) {
	return s.repo.GetAllJobs(ctx)
}
