package esidev

/*
yesterday object */
type GetFwLeaderboardsYesterday1 struct {
	/*
	 faction_id integer */
	FactionId int32 `json:"faction_id,omitempty"`
	/*
	 Amount of victory points */
	Amount int32 `json:"amount,omitempty"`
}
