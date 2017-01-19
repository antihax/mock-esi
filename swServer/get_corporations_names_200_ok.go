package swServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetCorporationsNames200Ok struct {
/*
	 corporation_id integer */
	corporation_id int32 `json:"corporation_id,omitempty"`
/*
	 corporation_name string */
	corporation_name string `json:"corporation_name,omitempty"`
}
