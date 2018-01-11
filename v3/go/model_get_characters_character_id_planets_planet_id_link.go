package esiv3

/*
link object */
type GetCharactersCharacterIdPlanetsPlanetIdLink struct {
	/*
	 source_pin_id integer */
	SourcePinId int64 `json:"source_pin_id,omitempty"`
	/*
	 destination_pin_id integer */
	DestinationPinId int64 `json:"destination_pin_id,omitempty"`
	/*
	 link_level integer */
	LinkLevel int32 `json:"link_level,omitempty"`
}
