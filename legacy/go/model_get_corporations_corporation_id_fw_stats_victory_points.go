package esilegacy



/* 
Summary of victory points gained by the given corporation for the enlisted faction */
type GetCorporationsCorporationIdFwStatsVictoryPoints struct {
/*
	 Last week's victory points gained by members of the given corporation */
	LastWeek int32 `json:"last_week,omitempty"`
/*
	 Total victory points gained since the given corporation enlisted */
	Total int32 `json:"total,omitempty"`
/*
	 Yesterday's victory points gained by members of the given corporation */
	Yesterday int32 `json:"yesterday,omitempty"`
}
