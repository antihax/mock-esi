package swaggerServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetCorporationsCorporationIdOk struct {
/*
	 id of alliance that corporation is a member of, if any */
	alliance_id int32 `json:"alliance_id,omitempty"`
/*
	 ceo_id integer */
	ceo_id int32 `json:"ceo_id,omitempty"`
/*
	 the full name of the corporation */
	corporation_name string `json:"corporation_name,omitempty"`
/*
	 member_count integer */
	member_count int32 `json:"member_count,omitempty"`
/*
	 the short name of the corporation */
	ticker string `json:"ticker,omitempty"`
}
