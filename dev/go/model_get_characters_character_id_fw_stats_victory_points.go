package esidev

/*
Summary of victory points gained by the given character for the enlisted faction */
type GetCharactersCharacterIdFwStatsVictoryPoints struct {
	/*
	 Last week's victory points gained by the given character */
	LastWeek int32 `json:"last_week,omitempty"`
	/*
	 Total victory points gained since the given character enlisted */
	Total int32 `json:"total,omitempty"`
	/*
	 Yesterday's victory points gained by the given character */
	Yesterday int32 `json:"yesterday,omitempty"`
}
