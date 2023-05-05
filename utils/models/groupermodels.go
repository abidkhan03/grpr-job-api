package models

type Group struct {
	Number                   int      `default:"0"`
	CommonLinks              []string `default:"[]"`
	LinksInCommon            int      `default:"0"`
	MainKeyword              string   `default:""`
	HeighestVolume           int      `default:"0"`
	HeighestVolumeKeyword    string   `default:""`
	AverageKwDifficulty      float64  `default:"0.0"`
	AverageRank              int      `default:"101"`
	SumOfCurrentValues       float64  `default:"0.0"`
	RankPercentage           float64  `default:"0.0"`
	TopicVolume              int      `default:"0"`
	QuartileVolume           float64  `default:"0.0"`
	AverageRankQuartile      float64  `default:"0.0"`
	SumValueOpporunity       float64  `default:"0.0"`
	SumVolumeOpporunity      float64  `default:"0.0"`
	VariantCount             int      `default:"0"`
	Relevancy                float64  `default:"1.0"`
	Cluster                  int      `default:"0"`
	PotentialContentGap      bool     `default:"false"`
	TotalContentGap          bool     `default:"true"`
	KeywordGap               bool     `default:"false"`
	PotentialCannibalization bool     `default:"false"`
	AutoMappedUrl            string   `default:""`
}

type GLink struct {
	URL                 string `default:""`
	Position            int    `default:"0"`
	RelatedResultsCount int    `default:"0"`
}

type GNode struct {
	Keyword                string   `default:""`
	Links                  []Link   `default:"[]"`
	SearchVolume           int      `default:"0"`
	Group                  []Group  `default:"None"`
	SubGroup               []Group  `default:"None"`
	PrimarySearchIntents   []string `default:"[]"`
	SecondarySearchIntents []string `default:"[]"`
	Rank                   Rank
	CPC                    float64 `default:"-1.0"`
	CPS                    float64 `default:"-1.0"`
	Difficulty             float64 `default:"0.0"`
	CurrentTraffic         float64 `default:"0.0"`
	PotentialTraffic       float64 `default:"0.0"`
	CurrentValue           float64 `default:"0.0"`
	PotentialValue         float64 `default:"0.0"`
	FibonacciHelper        int     `default:"0"`
	VolumePercent          float64 `default:"0.0"`
	PriorityScore          float64 `default:"0.0"`
	ValueOpportunity       float64 `default:"0.0"`
	ColumeOpportunity      float64 `default:"0.0"`
	CompetitorRanks        []Rank  `default:"[]"`
	CompetitorScore        int     `default:"1"`
	CompetitorRankingCount int     `default:"0"`
}
