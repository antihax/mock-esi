package esiv1

/*
active_total object */
type GetFwLeaderboardsCharactersActiveTotal struct {
	/*
	 character_id integer */
	CharacterId int32 `json:"character_id,omitempty"`
	/*
	 Amount of kills */
	Amount int32 `json:"amount,omitempty"`
}
