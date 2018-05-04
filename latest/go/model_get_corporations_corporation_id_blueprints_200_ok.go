package esilatest



/* 
200 ok object */
type GetCorporationsCorporationIdBlueprints200Ok struct {
/*
	 Unique ID for this item. */
	ItemId int64 `json:"item_id,omitempty"`
/*
	 Type of the location_id */
	LocationFlag string `json:"location_flag,omitempty"`
/*
	 References a solar system, station or item_id if this blueprint is located within a container. */
	LocationId int64 `json:"location_id,omitempty"`
/*
	 Material Efficiency Level of the blueprint. */
	MaterialEfficiency int32 `json:"material_efficiency,omitempty"`
/*
	 A range of numbers with a minimum of -2 and no maximum value where -1 is an original and -2 is a copy. It can be a positive integer if it is a stack of blueprint originals fresh from the market (e.g. no activities performed on them yet). */
	Quantity int32 `json:"quantity,omitempty"`
/*
	 Number of runs remaining if the blueprint is a copy, -1 if it is an original. */
	Runs int32 `json:"runs,omitempty"`
/*
	 Time Efficiency Level of the blueprint. */
	TimeEfficiency int32 `json:"time_efficiency,omitempty"`
/*
	 type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
}
