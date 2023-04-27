package dao

import (
	"time"

	"github.com/grpr-job-api/internal/errors"
)

type GrouperJob struct {
	Id                             uint      `db:"id"`
	Created                        time.Time `db:"created"`
	ReportTitle                    string    `db:"report_title"`
	GroupingMethod                 string    `db:"grouping_method"`
	MainKeywordGroupingAccuracy    int       `db:"main_keyword_group"`
	VariantKeywordGroupingAccuracy int       `db:"variant_keyword_group"`
}

func (g *GrouperJob) GetID() uint {
	return g.Id
}

func (*GrouperJob) GetTable() string {
	return "grouperjob"
}

func (g *GrouperJob) Validate() error {
	if g.ReportTitle == "" {
		return errors.BadRequest("Report title is required")
	}
	if g.MainKeywordGroupingAccuracy == 0 {
		return errors.BadRequest("Main keyword grouping accuracy is required")
	}
	if g.VariantKeywordGroupingAccuracy == 0 {
		return errors.BadRequest("Variant keyword grouping accuracy is required")
	}
	return nil
}
