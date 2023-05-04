package repo

import (
	"context"

	"github.com/grpr-job-api/internal/dao"
)

func (r *Repo) GetCombinedJobs(ctx context.Context) ([]*dao.CombinedJob, error) {
	var job []*dao.CombinedJob
	err := r.db(ctx).Table("combinedjob").Find(&job).Error
	return job, wrap(err)
}
