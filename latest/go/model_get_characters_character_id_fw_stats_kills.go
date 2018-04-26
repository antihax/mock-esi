package esilatest

/*
Summary of kills done by the given character against enemy factions */
type GetCharactersCharacterIdFwStatsKills struct {
	/*
	 Last week's total number of kills by a given character against enemy factions */
	LastWeek int32 `json:"last_week,omitempty"`
	/*
	 Total number of kills by a given character against enemy factions since the character enlisted */
	Total int32 `json:"total,omitempty"`
	/*
	 Yesterday's total number of kills by a given character against enemy factions */
	Yesterday int32 `json:"yesterday,omitempty"`
}
