package esilatest

import "time"

/*
200 ok object */
type GetCorporationsCorporationIdFwStatsOk struct {
	/*
	 The faction the given corporation is enlisted to fight for. Will not be included if corporation is not enlisted in faction warfare */
	FactionId int32 `json:"faction_id,omitempty"`
	/*
	 The enlistment date of the given corporation into faction warfare. Will not be included if corporation is not enlisted in faction warfare */
	EnlistedOn time.Time `json:"enlisted_on,omitempty"`
	/*
	 How many pilots the enlisted corporation has. Will not be included if corporation is not enlisted in faction warfare */
	Pilots int32 `json:"pilots,omitempty"`
	/* */
	Kills GetCorporationsCorporationIdFwStatsKills `json:"kills,omitempty"`
	/* */
	VictoryPoints GetCorporationsCorporationIdFwStatsVictoryPoints `json:"victory_points,omitempty"`
}
