package esilegacy

/*
200 ok object */
type GetUniverseStargatesStargateIdOk struct {
	/*
	 stargate_id integer */
	StargateId int32 `json:"stargate_id,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
	/* */
	Position GetUniverseStargatesStargateIdPosition `json:"position,omitempty"`
	/*
	 The solar system this stargate is in */
	SystemId int32 `json:"system_id,omitempty"`
	/* */
	Destination GetUniverseStargatesStargateIdDestination `json:"destination,omitempty"`
}
