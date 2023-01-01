package esiv5

/*
200 ok object
*/
type GetCorporationsCorporationIdAssets200Ok struct {
	/*
	 is_blueprint_copy boolean */
	IsBlueprintCopy bool `json:"is_blueprint_copy,omitempty"`
	/*
	 is_singleton boolean */
	IsSingleton bool `json:"is_singleton,omitempty"`
	/*
	 item_id integer */
	ItemId int64 `json:"item_id,omitempty"`
	/*
	 location_flag string */
	LocationFlag string `json:"location_flag,omitempty"`
	/*
	 location_id integer */
	LocationId int64 `json:"location_id,omitempty"`
	/*
	 location_type string */
	LocationType string `json:"location_type,omitempty"`
	/*
	 quantity integer */
	Quantity int32 `json:"quantity,omitempty"`
	/*
	 type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
}
