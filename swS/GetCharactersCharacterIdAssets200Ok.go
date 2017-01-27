package swS

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersCharacterIdAssets200Ok struct {
/*
	 is_singleton boolean */
	is_singleton bool `json:"is_singleton,omitempty"`
/*
	 item_id integer */
	item_id int64 `json:"item_id,omitempty"`
/*
	 location_flag string */
	location_flag string `json:"location_flag,omitempty"`
/*
	 location_id integer */
	location_id int64 `json:"location_id,omitempty"`
/*
	 location_type string */
	location_type string `json:"location_type,omitempty"`
/*
	 quantity integer */
	quantity int32 `json:"quantity,omitempty"`
/*
	 type_id integer */
	type_id int32 `json:"type_id,omitempty"`
}
