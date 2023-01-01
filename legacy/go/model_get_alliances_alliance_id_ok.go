package esilegacy

import "time"

/*
200 ok object
*/
type GetAlliancesAllianceIdOk struct {
	/*
	 ID of the corporation that created the alliance */
	CreatorCorporationId int32 `json:"creator_corporation_id,omitempty"`
	/*
	 ID of the character that created the alliance */
	CreatorId int32 `json:"creator_id,omitempty"`
	/*
	 date_founded string */
	DateFounded time.Time `json:"date_founded,omitempty"`
	/*
	 the executor corporation ID, if this alliance is not closed */
	ExecutorCorporationId int32 `json:"executor_corporation_id,omitempty"`
	/*
	 Faction ID this alliance is fighting for, if this alliance is enlisted in factional warfare */
	FactionId int32 `json:"faction_id,omitempty"`
	/*
	 the full name of the alliance */
	Name string `json:"name,omitempty"`
	/*
	 the short name of the alliance */
	Ticker string `json:"ticker,omitempty"`
}
