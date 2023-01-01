package esilegacy

/*
200 ok object
*/
type GetUniverseStationsStationIdOk struct {
	/*
	 solar_system_id integer */
	SolarSystemId int32 `json:"solar_system_id,omitempty"`
	/*
	 the full name of the station */
	StationName string `json:"station_name,omitempty"`
}
