package swS

import "time"
var _ time.Time

/* 
dogma_effect object */
type GetUniverseTypesTypeIdDogmaEffect struct {
/*
	 effect_id integer */
	effect_id int32 `json:"effect_id,omitempty"`
/*
	 is_default boolean */
	is_default bool `json:"is_default,omitempty"`
}
