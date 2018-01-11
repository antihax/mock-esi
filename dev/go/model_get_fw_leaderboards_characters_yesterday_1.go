package esidev

/*
yesterday object */
type GetFwLeaderboardsCharactersYesterday1 struct {
	/*
	 character_id integer */
	CharacterId int32 `json:"character_id,omitempty"`
	/*
	 Amount of victory points */
	Amount int32 `json:"amount,omitempty"`
}
