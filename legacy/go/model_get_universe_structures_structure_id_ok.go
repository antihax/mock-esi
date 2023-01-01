package esilegacy

/*
200 ok object
*/
type GetUniverseStructuresStructureIdOk struct {
	/*
	 The full name of the structure */
	Name string `json:"name,omitempty"`
	/* */
	Position GetUniverseStructuresStructureIdPosition `json:"position,omitempty"`
	/*
	 solar_system_id integer */
	SolarSystemId int32 `json:"solar_system_id,omitempty"`
	/*
	 type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
}
