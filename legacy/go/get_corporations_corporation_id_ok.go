package esilegacy

/*
200 ok object */
type GetCorporationsCorporationIdOk struct {
	/*
	 the full name of the corporation */
	CorporationName string `json:"corporation_name,omitempty"`
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
	 id of alliance that corporation is a member of, if any */
	AllianceId int32 `json:"alliance_id,omitempty"`
}
