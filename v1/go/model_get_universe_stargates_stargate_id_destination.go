package esiv1

/*
destination object
*/
type GetUniverseStargatesStargateIdDestination struct {
	/*
	 The stargate this stargate connects to */
	StargateId int32 `json:"stargate_id,omitempty"`
	/*
	 The solar system this stargate connects to */
	SystemId int32 `json:"system_id,omitempty"`
}
