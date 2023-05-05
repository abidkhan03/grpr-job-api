package repo

import (
	"context"

	"github.com/grpr-job-api/internal/dao"
)

func (r *Repo) GetFetcherJobs(ctx context.Context) ([]*dao.FetcherJob, error) {
	var jobs []*dao.FetcherJob
	err := r.db(ctx).Table("fetcherjob").Find(&jobs).Error
	return jobs, wrap(err)
}

func (r *Repo) AddFetcherJob(ctx context.Context, job *dao.FetcherJob) error {
	err := r.db(ctx).Table("fetcherjob").Create(job).Error
	return wrap(err)
}
