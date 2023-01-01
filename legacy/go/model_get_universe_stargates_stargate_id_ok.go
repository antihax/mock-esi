package esilegacy

/*
200 ok object
*/
type GetUniverseStargatesStargateIdOk struct {
	/* */
	Destination GetUniverseStargatesStargateIdDestination `json:"destination,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/* */
	Position GetUniverseStargatesStargateIdPosition `json:"position,omitempty"`
	/*
	 stargate_id integer */
	StargateId int32 `json:"stargate_id,omitempty"`
	/*
	 The solar system this stargate is in */
	SystemId int32 `json:"system_id,omitempty"`
	/*
	 type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
}
