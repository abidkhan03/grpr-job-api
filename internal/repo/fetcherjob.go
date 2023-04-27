package repo

import (
	"context"

	"github.com/grpr-job-api/internal/dao"
)

func (r *Repo) GetFetcherJobs(ctx context.Context) ([]*dao.FetcherJob, error) {
	var job []*dao.FetcherJob
	err := r.db(ctx).Find(&job).Error
	return job, wrap(err)
}
