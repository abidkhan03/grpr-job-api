package models

type Rank_ struct {
	ClientRankingUrl      string
	ClientRankingPosition int
	ClientUrlRankingCount int
	CurrentTraffic        float64
	CurrentValue          float64
}

var Rank = Rank_{
	ClientRankingUrl:      "",
	ClientRankingPosition: 101,
	ClientUrlRankingCount: 0,
	CurrentTraffic:        0.0,
	CurrentValue:          0.0,
}
