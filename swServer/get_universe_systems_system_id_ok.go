package swServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetUniverseSystemsSystemIdOk struct {
/*
	 the full name of the system */
	solar_system_name string `json:"solar_system_name,omitempty"`
}
