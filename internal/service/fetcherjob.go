package service

import (
	"context"

	"github.com/grpr-job-api/internal/dao"
)

func (s *Service) GetFetcherJobs(ctx context.Context) ([]*dao.FetcherJob, error) {
	return s.repo.GetFetcherJobs(ctx)
}

// AddFetcherJob adds a new fetcher job to the database
func (s *Service) AddFetcherJob(ctx context.Context, job *dao.FetcherJob) error {
	return s.repo.AddFetcherJob(ctx, job)
}
