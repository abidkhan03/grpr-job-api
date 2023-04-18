package response

import "time"

type FetcherJob struct {
	Id           uint      `json:"id"`
	Created      time.Time `json:"created"`
	ReportTitle  string    `json:"report_title"`
	SearchEngine string    `json:"search_engine"`
	Location     string    `json:"location"`
	Language     string    `json:"language"`
}
