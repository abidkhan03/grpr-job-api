package response

import "time"

type CombinedJob struct {
	Id                             uint      `json:"id"`
	Created                        time.Time `json:"created"`
	ReportTitle                    string    `json:"report_title"`
	GroupingMethod                 string    `json:"grouping_method"`
	MainKeywordGroupingAccuracy    int       `json:"main_keyword_grouping_accuracy"`
	VariantKeywordGroupingAccuracy int       `json:"variant_keyword_group"`
	SearchEngine                   string    `json:"search_engine"`
	Location                       string    `json:"location"`
	Language                       string    `json:"language"`
}
