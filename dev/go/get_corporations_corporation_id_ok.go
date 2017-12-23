package esidev

import "time"

/*
200 ok object */
type GetCorporationsCorporationIdOk struct {
	/*
	 the full name of the corporation */
	Name string `json:"name,omitempty"`
	/*
	 the short name of the corporation */
	Ticker string `json:"ticker,omitempty"`
	/*
	 member_count integer */
	MemberCount int32 `json:"member_count,omitempty"`
	/*
	 ceo_id integer */
	CeoId int32 `json:"ceo_id,omitempty"`
	/*
	 ID of the alliance that corporation is a member of, if any */
	AllianceId int32 `json:"alliance_id,omitempty"`
	/*
	 description string */
	Description string `json:"description,omitempty"`
	/*
	 tax_rate number */
	TaxRate float32 `json:"tax_rate,omitempty"`
	/*
	 date_founded string */
	DateFounded time.Time `json:"date_founded,omitempty"`
	/*
	 creator_id integer */
	CreatorId int32 `json:"creator_id,omitempty"`
	/*
	 url string */
	Url string `json:"url,omitempty"`
	/*
	 faction_id integer */
	FactionId int32 `json:"faction_id,omitempty"`
	/*
	 home_station_id integer */
	HomeStationId int32 `json:"home_station_id,omitempty"`
	/*
	 shares integer */
	Shares int64 `json:"shares,omitempty"`
}