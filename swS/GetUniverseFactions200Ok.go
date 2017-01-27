package swS

import "time"
var _ time.Time

/* 
200 ok object */
type GetUniverseFactions200Ok struct {
/*
	 corporation_id integer */
	corporation_id int32 `json:"corporation_id,omitempty"`
/*
	 description string */
	description string `json:"description,omitempty"`
/*
	 faction_id integer */
	faction_id int32 `json:"faction_id,omitempty"`
/*
	 is_unique boolean */
	is_unique bool `json:"is_unique,omitempty"`
/*
	 militia_corporation_id integer */
	militia_corporation_id int32 `json:"militia_corporation_id,omitempty"`
/*
	 name string */
	name string `json:"name,omitempty"`
/*
	 size_factor number */
	size_factor float32 `json:"size_factor,omitempty"`
/*
	 solar_system_id integer */
	solar_system_id int32 `json:"solar_system_id,omitempty"`
/*
	 station_count integer */
	station_count int32 `json:"station_count,omitempty"`
/*
	 station_system_count integer */
	station_system_count int32 `json:"station_system_count,omitempty"`
}
