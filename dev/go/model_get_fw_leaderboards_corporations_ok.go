package esidev

/*
200 ok object
*/
type GetFwLeaderboardsCorporationsOk struct {
	/* */
	Kills GetFwLeaderboardsCorporationsKills `json:"kills,omitempty"`
	/* */
	VictoryPoints GetFwLeaderboardsCorporationsVictoryPoints `json:"victory_points,omitempty"`
}
