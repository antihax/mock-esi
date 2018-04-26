package esiv1

/*
last_week object */
type GetFwLeaderboardsCharactersLastWeek struct {
	/*
	 Amount of kills */
	Amount int32 `json:"amount,omitempty"`
	/*
	 character_id integer */
	CharacterId int32 `json:"character_id,omitempty"`
}
