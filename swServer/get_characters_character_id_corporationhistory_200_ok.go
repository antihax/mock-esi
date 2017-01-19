package swServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersCharacterIdCorporationhistory200Ok struct {
/*
	 corporation_id integer */
	corporation_id int32 `json:"corporation_id,omitempty"`
/*
	 True if the corporation has been deleted */
	is_deleted bool `json:"is_deleted,omitempty"`
/*
	 An incrementing ID that can be used to canonically establish order of records in cases where dates may be ambiguous */
	record_id int32 `json:"record_id,omitempty"`
/*
	 start_date string */
	start_date time.Time `json:"start_date,omitempty"`
}
