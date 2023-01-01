package esilatest

/*
200 ok object
*/
type GetFwLeaderboardsOk struct {
	/* */
	Kills GetFwLeaderboardsKills `json:"kills,omitempty"`
	/* */
	VictoryPoints GetFwLeaderboardsVictoryPoints `json:"victory_points,omitempty"`
}
