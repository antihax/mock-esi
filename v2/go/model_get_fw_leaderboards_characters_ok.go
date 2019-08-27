package esiv2

/*
200 ok object */
type GetFwLeaderboardsCharactersOk struct {
	/* */
	Kills GetFwLeaderboardsCharactersKills `json:"kills,omitempty"`
	/* */
	VictoryPoints GetFwLeaderboardsCharactersVictoryPoints `json:"victory_points,omitempty"`
}
