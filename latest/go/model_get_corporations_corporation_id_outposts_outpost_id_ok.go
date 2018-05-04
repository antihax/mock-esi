package esilatest



/* 
200 ok object */
type GetCorporationsCorporationIdOutpostsOutpostIdOk struct {
/* */
	Coordinates GetCorporationsCorporationIdOutpostsOutpostIdCoordinates `json:"coordinates,omitempty"`
/*
	 docking_cost_per_ship_volume number */
	DockingCostPerShipVolume float32 `json:"docking_cost_per_ship_volume,omitempty"`
/*
	 office_rental_cost integer */
	OfficeRentalCost int64 `json:"office_rental_cost,omitempty"`
/*
	 The entity that owns the station (e.g. the entity whose logo is on the station services bar) */
	OwnerId int32 `json:"owner_id,omitempty"`
/*
	 reprocessing_efficiency number */
	ReprocessingEfficiency float32 `json:"reprocessing_efficiency,omitempty"`
/*
	 reprocessing_station_take number */
	ReprocessingStationTake float32 `json:"reprocessing_station_take,omitempty"`
/*
	 A list of services the given outpost provides */
	Services []GetCorporationsCorporationIdOutpostsOutpostIdService `json:"services,omitempty"`
/*
	 The owner ID that sets the ability for someone to dock based on standings. */
	StandingOwnerId int32 `json:"standing_owner_id,omitempty"`
/*
	 The ID of the solar system the outpost rests in */
	SystemId int32 `json:"system_id,omitempty"`
/*
	 The type ID of the given outpost */
	TypeId int32 `json:"type_id,omitempty"`
}
