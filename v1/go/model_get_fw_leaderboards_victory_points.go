package esiv1

/*
Top 4 rankings of factions by victory points from yesterday, last week and in total */
type GetFwLeaderboardsVictoryPoints struct {
	/*
	 Top 4 ranking of factions active in faction warfare by total victory points. A faction is considered \"active\" if they have participated in faction warfare in the past 14 days. */
	ActiveTotal []GetFwLeaderboardsActiveTotal1 `json:"active_total,omitempty"`
	/*
	 Top 4 ranking of factions by victory points in the past week */
	LastWeek []GetFwLeaderboardsLastWeek1 `json:"last_week,omitempty"`
	/*
	 Top 4 ranking of factions by victory points in the past day */
	Yesterday []GetFwLeaderboardsYesterday1 `json:"yesterday,omitempty"`
}
