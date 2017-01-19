package swServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetUniverseSchematicsSchematicIdOk struct {
/*
	 Time in seconds to process a run */
	cycle_time int32 `json:"cycle_time,omitempty"`
/*
	 schematic_name string */
	schematic_name string `json:"schematic_name,omitempty"`
}
