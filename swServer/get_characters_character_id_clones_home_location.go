package swServer

import "time"
var _ time.Time

/* 
home_location object */
type GetCharactersCharacterIdClonesHomeLocation struct {
/*
	 location_id integer */
	location_id int64 `json:"location_id,omitempty"`
/*
	 location_type string */
	location_type string `json:"location_type,omitempty"`
}
