package swServer

import "time"
var _ time.Time

/* 
waypoint object */
type GetCharactersCharacterIdPlanetsPlanetIdWaypoint struct {
/*
	 order integer */
	order int32 `json:"order,omitempty"`
/*
	 pin_id integer */
	pin_id int64 `json:"pin_id,omitempty"`
}
