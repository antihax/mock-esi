package esidev

/*
Top 100 rankings of pilots by victory points from yesterday, last week and in total */
type GetFwLeaderboardsCharactersVictoryPoints struct {
	/*
	 Top 100 ranking of pilots active in faction warfare by total victory points. A pilot is considered \"active\" if they have participated in faction warfare in the past 14 days. */
	ActiveTotal []GetFwLeaderboardsCharactersActiveTotal1 `json:"active_total,omitempty"`
	/*
	 Top 100 ranking of pilots by victory points in the past week */
	LastWeek []GetFwLeaderboardsCharactersLastWeek1 `json:"last_week,omitempty"`
	/*
	 Top 100 ranking of pilots by victory points in the past day */
	Yesterday []GetFwLeaderboardsCharactersYesterday1 `json:"yesterday,omitempty"`
}
