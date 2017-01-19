package swServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetUniverseStationsStationIdOk struct {
/*
	 solar_system_id integer */
	solar_system_id int32 `json:"solar_system_id,omitempty"`
/*
	 the full name of the station */
	station_name string `json:"station_name,omitempty"`
}
