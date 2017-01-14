package swaggerServer

import "time"
var _ time.Time

/* 
Coordinates of the structure in Cartesian space relative to the Sun, in metres.
 */
type GetUniverseStructuresStructureIdPosition struct {
/*
	 x number */
	x float32 `json:"x,omitempty"`
/*
	 y number */
	y float32 `json:"y,omitempty"`
/*
	 z number */
	z float32 `json:"z,omitempty"`
}
