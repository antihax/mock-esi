package esilegacy

/*
200 ok object */
type GetUniverseSystemJumps200Ok struct {
	/*
	 ship_jumps integer */
	ShipJumps int32 `json:"ship_jumps,omitempty"`
	/*
	 system_id integer */
	SystemId int32 `json:"system_id,omitempty"`
}
