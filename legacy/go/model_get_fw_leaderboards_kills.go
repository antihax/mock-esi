package esilegacy

/*
Top 4 rankings of factions by number of kills from yesterday, last week and in total */
type GetFwLeaderboardsKills struct {
	/*
	 Top 4 ranking of factions active in faction warfare by total kills. A faction is considered \"active\" if they have participated in faction warfare in the past 14 days */
	ActiveTotal []GetFwLeaderboardsActiveTotalActiveTotal `json:"active_total,omitempty"`
	/*
	 Top 4 ranking of factions by kills in the past week */
	LastWeek []GetFwLeaderboardsLastWeekLastWeek `json:"last_week,omitempty"`
	/*
	 Top 4 ranking of factions by kills in the past day */
	Yesterday []GetFwLeaderboardsYesterdayYesterday `json:"yesterday,omitempty"`
}
