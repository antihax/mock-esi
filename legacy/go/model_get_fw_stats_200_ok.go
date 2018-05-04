package esilegacy



/* 
200 ok object */
type GetFwStats200Ok struct {
/*
	 faction_id integer */
	FactionId int32 `json:"faction_id,omitempty"`
/* */
	Kills GetFwStatsKills `json:"kills,omitempty"`
/*
	 How many pilots fight for the given faction */
	Pilots int32 `json:"pilots,omitempty"`
/*
	 The number of solar systems controlled by the given faction */
	SystemsControlled int32 `json:"systems_controlled,omitempty"`
/* */
	VictoryPoints GetFwStatsVictoryPoints `json:"victory_points,omitempty"`
}
