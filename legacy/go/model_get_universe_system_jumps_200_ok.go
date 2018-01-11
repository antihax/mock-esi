package esilegacy

/*
200 ok object */
type GetUniverseSystemJumps200Ok struct {
	/*
	 system_id integer */
	SystemId int32 `json:"system_id,omitempty"`
	/*
	 ship_jumps integer */
	ShipJumps int32 `json:"ship_jumps,omitempty"`
}
