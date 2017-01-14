package swaggerServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetUniverseStructuresStructureIdOk struct {
/*
	 The full name of the structure */
	name string `json:"name,omitempty"`
/* */
	position GetUniverseStructuresStructureIdPosition `json:"position,omitempty"`
/*
	 solar_system_id integer */
	solar_system_id int32 `json:"solar_system_id,omitempty"`
/*
	 type_id integer */
	type_id int32 `json:"type_id,omitempty"`
}
