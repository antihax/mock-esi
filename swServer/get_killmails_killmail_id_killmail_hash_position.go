package swServer

import "time"
var _ time.Time

/* 
Coordinates of the victim in Cartesian space relative to the Sun
 */
type GetKillmailsKillmailIdKillmailHashPosition struct {
/*
	 x number */
	x float32 `json:"x,omitempty"`
/*
	 y number */
	y float32 `json:"y,omitempty"`
/*
	 z number */
	z float32 `json:"z,omitempty"`
}
