package esiv1

/*
200 ok object */
type GetCharactersCharacterIdAssets200Ok struct {
	/*
	 type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
	/*
	 quantity integer */
	Quantity int32 `json:"quantity,omitempty"`
	/*
	 location_id integer */
	LocationId int64 `json:"location_id,omitempty"`
	/*
	 location_type string */
	LocationType string `json:"location_type,omitempty"`
	/*
	 item_id integer */
	ItemId int64 `json:"item_id,omitempty"`
	/*
	 location_flag string */
	LocationFlag string `json:"location_flag,omitempty"`
	/*
	 is_singleton boolean */
	IsSingleton bool `json:"is_singleton,omitempty"`
}
