package esidev

/*
active_total object */
type GetFwLeaderboardsActiveTotal struct {
	/*
	 Amount of kills */
	Amount int32 `json:"amount,omitempty"`
	/*
	 faction_id integer */
	FactionId int32 `json:"faction_id,omitempty"`
}
