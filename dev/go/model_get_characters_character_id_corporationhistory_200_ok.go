package esidev

import "time"

/*
200 ok object */
type GetCharactersCharacterIdCorporationhistory200Ok struct {
	/*
	 corporation_id integer */
	CorporationId int32 `json:"corporation_id,omitempty"`
	/*
	 True if the corporation has been deleted */
	IsDeleted bool `json:"is_deleted,omitempty"`
	/*
	 An incrementing ID that can be used to canonically establish order of records in cases where dates may be ambiguous */
	RecordId int32 `json:"record_id,omitempty"`
	/*
	 start_date string */
	StartDate time.Time `json:"start_date,omitempty"`
}
