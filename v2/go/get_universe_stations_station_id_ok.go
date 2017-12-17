package esiv2

/*
200 ok object */
type GetUniverseStationsStationIdOk struct {
	/*
	 station_id integer */
	StationId int32 `json:"station_id,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 ID of the corporation that controls this station */
	Owner int32 `json:"owner,omitempty"`
	/*
	 type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
	/*
	 race_id integer */
	RaceId int32 `json:"race_id,omitempty"`
	/* */
	Position GetUniverseStationsStationIdPosition `json:"position,omitempty"`
	/*
	 The solar system this station is in */
	SystemId int32 `json:"system_id,omitempty"`
	/*
	 reprocessing_efficiency number */
	ReprocessingEfficiency float32 `json:"reprocessing_efficiency,omitempty"`
	/*
	 reprocessing_stations_take number */
	ReprocessingStationsTake float32 `json:"reprocessing_stations_take,omitempty"`
	/*
	 max_dockable_ship_volume number */
	MaxDockableShipVolume float32 `json:"max_dockable_ship_volume,omitempty"`
	/*
	 office_rental_cost number */
	OfficeRentalCost float32 `json:"office_rental_cost,omitempty"`
	/*
	 services array */
	Services []string `json:"services,omitempty"`
}
