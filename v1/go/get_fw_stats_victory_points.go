package esiv1

/*
Summary of victory points gained for the given faction */
type GetFwStatsVictoryPoints struct {
	/*
	 Yesterday's victory points gained */
	Yesterday int32 `json:"yesterday,omitempty"`
	/*
	 Last week's victory points gained */
	LastWeek int32 `json:"last_week,omitempty"`
	/*
	 Total victory points gained since faction warfare began */
	Total int32 `json:"total,omitempty"`
}
