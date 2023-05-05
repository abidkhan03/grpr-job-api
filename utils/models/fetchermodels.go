package models

import (
	"github.com/grpr-job-api/utils/constants"
)

type Feature struct {
	AnswerBox                 bool `default:"false"`
	OrganicResultCount        int  `default:"0"`
	AdResultTopCount          int  `default:"0"`
	AdResultBottomCount       int  `default:"0"`
	AdResultRightCount        int  `default:"0"`
	FeaturedSnippet           bool `default:"false"`
	SitelinksSearchBox        bool `default:"false"`
	SitelinksExpanded         bool `default:"false"`
	SitelinksInline           bool `default:"false"`
	EventsResults             bool `default:"false"`
	InlineImages              bool `default:"false"`
	InlinePeopleAlsoSearchFor bool `default:"false"`
	ShoppingResults           bool `default:"false"`
	InlineVideos              bool `default:"false"`
	InlineVideoCarousels      bool `default:"false"`
	KnowledgeGraph            bool `default:"false"`
	LocalResults              bool `default:"false"`
	NewsResults               bool `default:"false"`
	TopStories                bool `default:"false"`
	InlineProducts            bool `default:"false"`
	RecipesResults            bool `default:"false"`
	RelatedQuestions          bool `default:"false"`
	TwitterResults            bool `default:"false"`
}

type Link struct {
	URL                 string `default:""`
	Position            int    `default:"0"`
	Title               string `default:""`
	Snippet             string `default:""`
	RelatedResultsCount int    `default:"0"`
}

type Node struct {
	Keyword                string   `default:""`
	Volume                 int      `default:"5"`
	Links                  []Link   `default:"[]"`
	PrimarySearchIntents   []string `default:"[]"`
	SecondarySearchIntents []string `default:"[]"`
	Rank                   Rank
	competitorRanks        []Rank  `default:"[]"`
	Difficulty             float64 `default:"0.0"`
	CPC                    float64 `default:"-1.0"`
	CPS                    float64 `default:"-1.0"`
	CurrentTraffic         float64 `default:"0.0"`
	PotentialTraffic       float64 `default:"0.0"`
	CurrentValue           float64 `default:"0.0"`
	PotentialValue         float64 `default:"0.0"`
	FibonacciHelper        int     `default:"0"`
}

type Arguments struct {
	InputFilePath string `default:""`
	APIKey        string `default:""`
	SearchEngine  string `default:"google"`
	Region        string `default:"United States"`
	GL            string `default:"us"`
}

var Args = Arguments{
	InputFilePath: "",
	APIKey:        "",
	SearchEngine:  constants.SEARCH_ENGINE,
	Region:        constants.REGION,
	GL:            constants.GL,
}
