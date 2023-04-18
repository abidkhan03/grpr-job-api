package request

import (
	"time"

	"github.com/grpr-job-api/internal/errors"
)

type FetcherJob struct {
	Id           uint      `json:"id"`
	Created      time.Time `json:"created"`
	ReportTitle  string    `json:"report_title"`
	SearchEngine string    `json:"search_engine"`
	Location     string    `json:"location"`
	Language     string    `json:"language"`
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
