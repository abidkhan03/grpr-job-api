package dao

import (
	"time"

	"github.com/grpr-job-api/internal/errors"
)

type CombinedJob struct {
	Id                             uint      `db:"id"`
	Created                        time.Time `db:"created"`
	ReportTitle                    string    `db:"report_title"`
	GroupingMethod                 string    `db:"grouping_method"`
	MainKeywordGroupingAccuracy    int       `db:"main_keyword_grouping_accuracy"`
	VariantKeywordGroupingAccuracy int       `db:"variant_keyword_group"`
	SearchEngine                   string    `db:"search_engine"`
	Location                       string    `db:"location"`
	Language                       string    `db:"language"`
}

func (f *CombinedJob) GetID() uint {
	return f.Id
}

func (*CombinedJob) GetTable() string {
	return "combinedjob"
}

func (f *CombinedJob) Validate() error {
	if f.ReportTitle == "" {
		return errors.BadRequest("Report title is required")
	}
	return nil
}
