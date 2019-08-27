package esiv2

/*
Summary of kills done by the given corporation against enemy factions */
type GetCorporationsCorporationIdFwStatsKills struct {
	/*
	 Last week's total number of kills by members of the given corporation against enemy factions */
	LastWeek int32 `json:"last_week,omitempty"`
	/*
	 Total number of kills by members of the given corporation against enemy factions since the corporation enlisted */
	Total int32 `json:"total,omitempty"`
	/*
	 Yesterday's total number of kills by members of the given corporation against enemy factions */
	Yesterday int32 `json:"yesterday,omitempty"`
}
