package swaggerServer

import "time"
var _ time.Time

/* 
link object */
type GetCharactersCharacterIdPlanetsPlanetIdLink struct {
/*
	 destination_pin_id integer */
	destination_pin_id int64 `json:"destination_pin_id,omitempty"`
/*
	 link_level integer */
	link_level int32 `json:"link_level,omitempty"`
/*
	 source_pin_id integer */
	source_pin_id int64 `json:"source_pin_id,omitempty"`
}
