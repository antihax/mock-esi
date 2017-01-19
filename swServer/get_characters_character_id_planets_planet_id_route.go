package swServer

import "time"
var _ time.Time

/* 
route object */
type GetCharactersCharacterIdPlanetsPlanetIdRoute struct {
/*
	 content_type_id integer */
	content_type_id int32 `json:"content_type_id,omitempty"`
/*
	 destination_pin_id integer */
	destination_pin_id int64 `json:"destination_pin_id,omitempty"`
/*
	 quantity integer */
	quantity int64 `json:"quantity,omitempty"`
/*
	 route_id integer */
	route_id int64 `json:"route_id,omitempty"`
/*
	 source_pin_id integer */
	source_pin_id int64 `json:"source_pin_id,omitempty"`
/*
	 waypoints array */
	waypoints []GetCharactersCharacterIdPlanetsPlanetIdWaypoint `json:"waypoints,omitempty"`
}
