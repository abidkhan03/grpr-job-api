package request

import (
	"time"

	"github.com/grpr-job-api/internal/errors"
)

type GrouperJob struct {
	Id                             uint      `json:"id"`
	Created                        time.Time `json:"created"`
	ReportTitle                    string    `json:"report_title"`
	GroupingMethod                 string    `json:"grouping_method"`
	MainKeywordGroupingAccuracy    int       `json:"main_keyword_group"`
	VariantKeywordGroupingAccuracy int       `json:"variant_keyword_group"`
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
