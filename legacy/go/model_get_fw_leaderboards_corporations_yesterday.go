package esilegacy

/*
yesterday object */
type GetFwLeaderboardsCorporationsYesterday struct {
	/*
	 corporation_id integer */
	CorporationId int32 `json:"corporation_id,omitempty"`
	/*
	 Amount of kills */
	Amount int32 `json:"amount,omitempty"`
}
