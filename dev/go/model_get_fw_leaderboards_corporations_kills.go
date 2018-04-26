package esidev

/*
Top 10 rankings of corporations by number of kills from yesterday, last week and in total. */
type GetFwLeaderboardsCorporationsKills struct {
	/*
	 Top 10 ranking of corporations active in faction warfare by total kills. A corporation is considered \"active\" if they have participated in faction warfare in the past 14 days. */
	ActiveTotal []GetFwLeaderboardsCorporationsActiveTotal `json:"active_total,omitempty"`
	/*
	 Top 10 ranking of corporations by kills in the past week */
	LastWeek []GetFwLeaderboardsCorporationsLastWeek `json:"last_week,omitempty"`
	/*
	 Top 10 ranking of corporations by kills in the past day */
	Yesterday []GetFwLeaderboardsCorporationsYesterday `json:"yesterday,omitempty"`
}
