package esilatest

/*
yesterday object */
type GetFwLeaderboardsCorporationsYesterday1 struct {
	/*
	 corporation_id integer */
	CorporationId int32 `json:"corporation_id,omitempty"`
	/*
	 Amount of victory points */
	Amount int32 `json:"amount,omitempty"`
}