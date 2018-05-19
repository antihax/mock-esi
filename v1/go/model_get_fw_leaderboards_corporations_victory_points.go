package esiv1

/*
Top 10 rankings of corporations by victory points from yesterday, last week and in total */
type GetFwLeaderboardsCorporationsVictoryPoints struct {
	/*
	 Top 10 ranking of corporations active in faction warfare by total victory points. A corporation is considered \"active\" if they have participated in faction warfare in the past 14 days. */
	ActiveTotal []GetFwLeaderboardsCorporationsActiveTotalActiveTotal1 `json:"active_total,omitempty"`
	/*
	 Top 10 ranking of corporations by victory points in the past week */
	LastWeek []GetFwLeaderboardsCorporationsLastWeekLastWeek1 `json:"last_week,omitempty"`
	/*
	 Top 10 ranking of corporations by victory points in the past day */
	Yesterday []GetFwLeaderboardsCorporationsYesterdayYesterday1 `json:"yesterday,omitempty"`
}
