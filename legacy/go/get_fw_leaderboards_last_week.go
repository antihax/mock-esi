package esilegacy

/*
last_week object */
type GetFwLeaderboardsLastWeek struct {
	/*
	 faction_id integer */
	FactionId int32 `json:"faction_id,omitempty"`
	/*
	 Amount of kills */
	Amount int32 `json:"amount,omitempty"`
}
