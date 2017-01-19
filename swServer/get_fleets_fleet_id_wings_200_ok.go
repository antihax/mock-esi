package swServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetFleetsFleetIdWings200Ok struct {
/*
	 id integer */
	id int64 `json:"id,omitempty"`
/*
	 name string */
	name string `json:"name,omitempty"`
/*
	 squads array */
	squads []GetFleetsFleetIdWingsSquad `json:"squads,omitempty"`
}
