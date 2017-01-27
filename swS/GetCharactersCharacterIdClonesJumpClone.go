package swS

import "time"
var _ time.Time

/* 
jump_clone object */
type GetCharactersCharacterIdClonesJumpClone struct {
/*
	 implants array */
	implants []int32 `json:"implants,omitempty"`
/*
	 location_id integer */
	location_id int64 `json:"location_id,omitempty"`
/*
	 location_type string */
	location_type string `json:"location_type,omitempty"`
}
