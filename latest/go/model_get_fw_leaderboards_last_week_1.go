package esilatest

/*
last_week object */
type GetFwLeaderboardsLastWeek1 struct {
	/*
	 Amount of victory points */
	Amount int32 `json:"amount,omitempty"`
	/*
	 faction_id integer */
	FactionId int32 `json:"faction_id,omitempty"`
}
