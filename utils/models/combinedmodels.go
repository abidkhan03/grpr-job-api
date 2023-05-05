package models

type Rank struct {
	ClientRankingUrl      string  `default:""`
	ClientRankingPosition int     `default:"101"`
	ClientUrlRankingCount int     `default:"0"`
	CurrentTraffic        float64 `default:"0.0"`
	CurrentValue          float64 `default:"0.0"`
}
