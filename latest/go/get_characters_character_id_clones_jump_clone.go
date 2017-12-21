package esilatest

/*
jump_clone object */
type GetCharactersCharacterIdClonesJumpClone struct {
	/*
	 jump_clone_id integer */
	JumpCloneId int32 `json:"jump_clone_id,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 location_id integer */
	LocationId int64 `json:"location_id,omitempty"`
	/*
	 location_type string */
	LocationType string `json:"location_type,omitempty"`
	/*
	 implants array */
	Implants []int32 `json:"implants,omitempty"`
}
