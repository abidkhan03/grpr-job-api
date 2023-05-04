package dao

import (
	"time"

	"github.com/grpr-job-api/internal/errors"
)

type FetcherJob struct {
	Id           uint      `db:"id"`
	Created      time.Time `db:"created"`
	ReportTitle  string    `db:"report_title"`
	SearchEngine string    `db:"search_engine"`
	Location     string    `db:"location"`
	Language     string    `db:"language"`
}

func (f *FetcherJob) GetID() uint {
	return f.Id
}

func (*FetcherJob) GetTable() string {
	return "fetcherjob"
}

func (f *FetcherJob) Validate() error {
	if f.Created.IsZero() {
		return errors.BadRequest("Created date is required")
	}
	if f.ReportTitle == "" {
		return errors.BadRequest("Report title is required")
	}
	return nil
}
