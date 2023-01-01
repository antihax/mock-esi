package esilegacy

/*
Coordinates of the structure in Cartesian space relative to the Sun, in metres.
*/
type GetUniverseStructuresStructureIdPosition struct {
	/*
	 x number */
	X float64 `json:"x,omitempty"`
	/*
	 y number */
	Y float64 `json:"y,omitempty"`
	/*
	 z number */
	Z float64 `json:"z,omitempty"`
}
