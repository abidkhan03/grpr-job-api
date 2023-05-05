package models

import (
	"github.com/grpr-job-api/utils/constants"
)

type Feature struct {
	AnswerBox                 bool
	OrganicResultCount        int
	AdResultTopCount          int
	AdResultBottomCount       int
	AdResultRightCount        int
	FeaturedSnippet           bool
	SitelinksSearchBox        bool
	SitelinksExpanded         bool
	SitelinksInline           bool
	EventsResults             bool
	InlineImages              bool
	InlinePeopleAlsoSearchFor bool
	ShoppingResults           bool
	InlineVideos              bool
	InlineVideoCarousels      bool
	KnowledgeGraph            bool
	LocalResults              bool
	NewsResults               bool
	TopStories                bool
	InlineProducts            bool
	RecipesResults            bool
	RelatedQuestions          bool
	TwitterResults            bool
}

type Link struct {
	URL                 string
	Position            int
	Title               string
	Snippet             string
	RelatedResultsCount int
}

type Node struct {
	Keyword                string
	Volume                 int
	Links                  []Link
	PrimarySearchIntents   []string
	SecondarySearchIntents []string
	// Rank                   []Rank
	// competitorRanks        []Rank
	Difficulty             float64
	CPC                    float64
	CPS                    float64
	CurrentTraffic         float64
	PotentialTraffic       float64
	CurrentValue           float64
	PotentialValue         float64
	FibonacciHelper        int
}

type Arguments struct {
	InputFilePath string
	APIKey        string
	SearchEngine  string
	Region        string
	GL            string
}

var Args = Arguments{
	InputFilePath: "",
	APIKey:        "",
	SearchEngine:  constants.SEARCH_ENGINE,
	Region:        constants.REGION,
	GL:            constants.GL,
}
