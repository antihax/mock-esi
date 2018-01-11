package esilatest

/*
last_week object */
type GetFwLeaderboardsCorporationsLastWeek1 struct {
	/*
	 corporation_id integer */
	CorporationId int32 `json:"corporation_id,omitempty"`
	/*
	 Amount of victory points */
	Amount int32 `json:"amount,omitempty"`
}
