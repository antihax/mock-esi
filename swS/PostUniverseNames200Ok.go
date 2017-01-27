package swS

import "time"
var _ time.Time

/* 
200 ok object */
type PostUniverseNames200Ok struct {
/*
	 category string */
	category string `json:"category,omitempty"`
/*
	 id integer */
	id int32 `json:"id,omitempty"`
/*
	 name string */
	name string `json:"name,omitempty"`
}
