package esilatest

import "time"

/*
200 ok object */
type GetCorporationsCorporationIdOk struct {
	/*
	 ID of the alliance that corporation is a member of, if any */
	AllianceId int32 `json:"alliance_id,omitempty"`
	/*
	 ceo_id integer */
	CeoId int32 `json:"ceo_id,omitempty"`
	/*
	 creator_id integer */
	CreatorId int32 `json:"creator_id,omitempty"`
	/*
	 date_founded string */
	DateFounded time.Time `json:"date_founded,omitempty"`
	/*
	 description string */
	Description string `json:"description,omitempty"`
	/*
	 faction_id integer */
	FactionId int32 `json:"faction_id,omitempty"`
	/*
	 home_station_id integer */
	HomeStationId int32 `json:"home_station_id,omitempty"`
	/*
	 member_count integer */
	MemberCount int32 `json:"member_count,omitempty"`
	/*
	 the full name of the corporation */
	Name string `json:"name,omitempty"`
	/*
	 shares integer */
	Shares int64 `json:"shares,omitempty"`
	/*
	 tax_rate number */
	TaxRate float32 `json:"tax_rate,omitempty"`
	/*
	 the short name of the corporation */
	Ticker string `json:"ticker,omitempty"`
	/*
	 url string */
	Url string `json:"url,omitempty"`
	/*
	 war_eligible boolean */
	WarEligible bool `json:"war_eligible,omitempty"`
}
