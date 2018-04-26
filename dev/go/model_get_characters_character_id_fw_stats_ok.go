package esidev

import "time"

/*
200 ok object */
type GetCharactersCharacterIdFwStatsOk struct {
	/*
	 The given character's current faction rank */
	CurrentRank int32 `json:"current_rank,omitempty"`
	/*
	 The enlistment date of the given character into faction warfare. Will not be included if character is not enlisted in faction warfare */
	EnlistedOn time.Time `json:"enlisted_on,omitempty"`
	/*
	 The faction the given character is enlisted to fight for. Will not be included if character is not enlisted in faction warfare */
	FactionId int32 `json:"faction_id,omitempty"`
	/*
	 The given character's highest faction rank achieved */
	HighestRank int32 `json:"highest_rank,omitempty"`
	/* */
	Kills GetCharactersCharacterIdFwStatsKills `json:"kills,omitempty"`
	/* */
	VictoryPoints GetCharactersCharacterIdFwStatsVictoryPoints `json:"victory_points,omitempty"`
}
