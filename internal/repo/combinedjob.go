package repo

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/grpr-job-api/internal/dao"
)

func (r *Repo) GetCombinedJobs(ctx context.Context) ([]*dao.CombinedJob, error) {
	var job []*dao.CombinedJob
	err := pgxscan.Select(ctx, r.Conn.Get(ctx), &job, `
	SELECT * FROM combinedjob ORDER BY id ASC`)
	return job, wrap(err)
}