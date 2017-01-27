package swS

import "time"
var _ time.Time

/* 
squad object */
type GetFleetsFleetIdWingsSquad struct {
/*
	 id integer */
	id int64 `json:"id,omitempty"`
/*
	 name string */
	name string `json:"name,omitempty"`
}
