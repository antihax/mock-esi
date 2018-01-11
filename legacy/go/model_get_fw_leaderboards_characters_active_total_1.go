package esilegacy

/*
active_total object */
type GetFwLeaderboardsCharactersActiveTotal1 struct {
	/*
	 character_id integer */
	CharacterId int32 `json:"character_id,omitempty"`
	/*
	 Amount of victory points */
	Amount int32 `json:"amount,omitempty"`
}
