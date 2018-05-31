package esilegacy

/*
Top 100 rankings of pilots by number of kills from yesterday, last week and in total. */
type GetFwLeaderboardsCharactersKills struct {
	/*
	 Top 100 ranking of pilots active in faction warfare by total kills. A pilot is considered \"active\" if they have participated in faction warfare in the past 14 days. */
	ActiveTotal []GetFwLeaderboardsCharactersActiveTotalActiveTotal `json:"active_total,omitempty"`
	/*
	 Top 100 ranking of pilots by kills in the past week */
	LastWeek []GetFwLeaderboardsCharactersLastWeekLastWeek `json:"last_week,omitempty"`
	/*
	 Top 100 ranking of pilots by kills in the past day */
	Yesterday []GetFwLeaderboardsCharactersYesterdayYesterday `json:"yesterday,omitempty"`
}
