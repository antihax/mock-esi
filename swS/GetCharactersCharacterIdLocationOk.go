package swS

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersCharacterIdLocationOk struct {
/*
	 solar_system_id integer */
	solar_system_id int32 `json:"solar_system_id,omitempty"`
/*
	 station_id integer */
	station_id int32 `json:"station_id,omitempty"`
/*
	 structure_id integer */
	structure_id int64 `json:"structure_id,omitempty"`
}
