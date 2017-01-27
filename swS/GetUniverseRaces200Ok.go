package swS

import "time"
var _ time.Time

/* 
200 ok object */
type GetUniverseRaces200Ok struct {
/*
	 The alliance generally associated with this race */
	alliance_id int32 `json:"alliance_id,omitempty"`
/*
	 description string */
	description string `json:"description,omitempty"`
/*
	 name string */
	name string `json:"name,omitempty"`
/*
	 race_id integer */
	race_id int32 `json:"race_id,omitempty"`
}
