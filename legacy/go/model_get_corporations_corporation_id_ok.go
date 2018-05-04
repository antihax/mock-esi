package esilegacy



/* 
200 ok object */
type GetCorporationsCorporationIdOk struct {
/*
	 id of alliance that corporation is a member of, if any */
	AllianceId int32 `json:"alliance_id,omitempty"`
/*
	 ceo_id integer */
	CeoId int32 `json:"ceo_id,omitempty"`
/*
	 corporation_description string */
	CorporationDescription string `json:"corporation_description,omitempty"`
/*
	 the full name of the corporation */
	CorporationName string `json:"corporation_name,omitempty"`
/*
	 creation_date string */
	CreationDate time.Time `json:"creation_date,omitempty"`
/*
	 creator_id integer */
	CreatorId int32 `json:"creator_id,omitempty"`
/*
	 faction string */
	Faction string `json:"faction,omitempty"`
/*
	 member_count integer */
	MemberCount int32 `json:"member_count,omitempty"`
/*
	 tax_rate number */
	TaxRate float32 `json:"tax_rate,omitempty"`
/*
	 the short name of the corporation */
	Ticker string `json:"ticker,omitempty"`
/*
	 url string */
	Url string `json:"url,omitempty"`
}
