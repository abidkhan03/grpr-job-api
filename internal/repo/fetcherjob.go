package repo

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/grpr-job-api/internal/dao"
)

func (r *Repo) GetFetcherJobs(ctx context.Context) ([]*dao.FetcherJob, error) {
	var fetcher []*dao.FetcherJob
	err := pgxscan.Select(ctx, r.Conn.Get(ctx), &fetcher, `
	SELECT * FROM fetcherjob ORDER BY id ASC`)
	return fetcher, wrap(err)
}

