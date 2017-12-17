package esidev

/*
Summary of kills against an enemy faction for the given faction */
type GetFwStatsKills struct {
	/*
	 Yesterday's total number of kills against enemy factions */
	Yesterday int32 `json:"yesterday,omitempty"`
	/*
	 Last week's total number of kills against enemy factions */
	LastWeek int32 `json:"last_week,omitempty"`
	/*
	 Total number of kills against enemy factions since faction warfare began */
	Total int32 `json:"total,omitempty"`
}
