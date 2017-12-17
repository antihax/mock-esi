package esilatest

/*
route object */
type GetCharactersCharacterIdPlanetsPlanetIdRoute struct {
	/*
	 route_id integer */
	RouteId int64 `json:"route_id,omitempty"`
	/*
	 source_pin_id integer */
	SourcePinId int64 `json:"source_pin_id,omitempty"`
	/*
	 destination_pin_id integer */
	DestinationPinId int64 `json:"destination_pin_id,omitempty"`
	/*
	 content_type_id integer */
	ContentTypeId int32 `json:"content_type_id,omitempty"`
	/*
	 quantity number */
	Quantity float32 `json:"quantity,omitempty"`
	/*
	 list of pin ID waypoints */
	Waypoints []int64 `json:"waypoints,omitempty"`
}
