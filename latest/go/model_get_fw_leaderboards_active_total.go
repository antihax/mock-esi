package esilatest

/*
active_total object */
type GetFwLeaderboardsActiveTotal struct {
	/*
	 faction_id integer */
	FactionId int32 `json:"faction_id,omitempty"`
	/*
	 Amount of kills */
	Amount int32 `json:"amount,omitempty"`
}
