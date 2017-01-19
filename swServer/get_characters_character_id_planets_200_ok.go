package swServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersCharacterIdPlanets200Ok struct {
/*
	 last_update string */
	last_update time.Time `json:"last_update,omitempty"`
/*
	 num_pins integer */
	num_pins int32 `json:"num_pins,omitempty"`
/*
	 owner_id integer */
	owner_id int32 `json:"owner_id,omitempty"`
/*
	 planet_id integer */
	planet_id int32 `json:"planet_id,omitempty"`
/*
	 planet_type string */
	planet_type string `json:"planet_type,omitempty"`
/*
	 solar_system_id integer */
	solar_system_id int32 `json:"solar_system_id,omitempty"`
/*
	 upgrade_level integer */
	upgrade_level int32 `json:"upgrade_level,omitempty"`
}
