package repo

import (
	"context"

	"github.com/grpr-job-api/internal/dao"
)

// GetGrouperJobs returns all grouper jobs
func (r *Repo) GetGrouperJobs(ctx context.Context) ([]*dao.GrouperJob, error) {
	var job []*dao.GrouperJob
	err := r.db(ctx).Find(&job).Error
	return job, wrap(err)
}
