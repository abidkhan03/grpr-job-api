package response

import "time"

type GrouperJob struct {
	Id                             uint      `json:"id"`
	Created                        time.Time `json:"created"`
	ReportTitle                    string    `json:"report_title"`
	GroupingMethod                 string    `json:"grouping_method"`
	MainKeywordGroupingAccuracy    int       `json:"main_keyword_group"`
	VariantKeywordGroupingAccuracy int       `json:"variant_keyword_group"`
}
