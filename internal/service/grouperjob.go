package service

import (
	"context"

	"github.com/grpr-job-api/internal/dao"
)

// GetGrouperJobs returns all grouper jobs
func (s *Service) GetGrouperJobs(ctx context.Context) ([]*dao.GrouperJob, error) {
	return s.repo.GetGrouperJobs(ctx)
}
