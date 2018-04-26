package esidev

/*
Summary of victory points gained for the given faction */
type GetFwStatsVictoryPoints struct {
	/*
	 Last week's victory points gained */
	LastWeek int32 `json:"last_week,omitempty"`
	/*
	 Total victory points gained since faction warfare began */
	Total int32 `json:"total,omitempty"`
	/*
	 Yesterday's victory points gained */
	Yesterday int32 `json:"yesterday,omitempty"`
}
