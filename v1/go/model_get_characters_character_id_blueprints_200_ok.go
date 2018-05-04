package esiv1

/*
200 ok object */
type GetCharactersCharacterIdBlueprints200Ok struct {
	/*
	 Unique ID for this item. The ID of an item is stable if that item is not repackaged, stacked, detached from a stack, assembled, or otherwise altered. If an item is changed in one of these ways, then the ID will also change. */
	ItemId int64 `json:"item_id,omitempty"`
	/*
	 Indicates something about this item's storage location. The flag is used to differentiate between hangar divisions, drone bay, fitting location, and similar. */
	LocationFlag string `json:"location_flag,omitempty"`
	/*
	 References a solar system, station or item_id if this blueprint is located within a container. If an item_id the Character - AssetList API must be queried to find the container using the item_id, from which the correct location of the Blueprint can be derived. */
	LocationId int64 `json:"location_id,omitempty"`
	/*
	 Material Efficiency Level of the blueprint, can be any integer between 0 and 10. */
	MaterialEfficiency int32 `json:"material_efficiency,omitempty"`
	/*
	 Typically will be -1 or -2 designating a singleton item, where -1 is an original and -2 is a copy. It can be a positive integer if it is a stack of blueprint originals fresh from the market (no activities performed on them yet). */
	Quantity int32 `json:"quantity,omitempty"`
	/*
	 Number of runs remaining if the blueprint is a copy, -1 if it is an original. */
	Runs int32 `json:"runs,omitempty"`
	/*
	 Time Efficiency Level of the blueprint, can be any even integer between 0 and 20. */
	TimeEfficiency int32 `json:"time_efficiency,omitempty"`
	/*
	 type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
}
