package esidev

/*
yesterday object */
type GetFwLeaderboardsYesterday struct {
	/*
	 faction_id integer */
	FactionId int32 `json:"faction_id,omitempty"`
	/*
	 Amount of kills */
	Amount int32 `json:"amount,omitempty"`
}
