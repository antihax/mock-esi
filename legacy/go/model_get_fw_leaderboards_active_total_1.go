package esilegacy

/*
active_total object */
type GetFwLeaderboardsActiveTotal1 struct {
	/*
	 Amount of victory points */
	Amount int32 `json:"amount,omitempty"`
	/*
	 faction_id integer */
	FactionId int32 `json:"faction_id,omitempty"`
}
