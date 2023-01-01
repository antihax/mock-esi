package esidev

/*
200 ok object
*/
type GetUniverseSchematicsSchematicIdOk struct {
	/*
	 Time in seconds to process a run */
	CycleTime int32 `json:"cycle_time,omitempty"`
	/*
	 schematic_name string */
	SchematicName string `json:"schematic_name,omitempty"`
}
