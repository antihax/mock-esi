package swS

import "time"
var _ time.Time

/* 
200 ok object */
type GetIncursions200Ok struct {
/*
	 The constellation id in which this incursion takes place */
	constellation_id int32 `json:"constellation_id,omitempty"`
/*
	 The attacking faction's id */
	faction_id int32 `json:"faction_id,omitempty"`
/*
	 Whether the final encounter has boss or not */
	has_boss bool `json:"has_boss,omitempty"`
/*
	 A list of infested solar system ids that are a part of this incursion */
	infested_solar_systems []int32 `json:"infested_solar_systems,omitempty"`
/*
	 Influence of this incursion as a float from 0 to 1 */
	influence float32 `json:"influence,omitempty"`
/*
	 Staging solar system for this incursion */
	staging_solar_system_id int32 `json:"staging_solar_system_id,omitempty"`
/*
	 The state of this incursion */
	state string `json:"state,omitempty"`
/*
	 The type of this incursion */
	_type string `json:"type,omitempty"`
}
