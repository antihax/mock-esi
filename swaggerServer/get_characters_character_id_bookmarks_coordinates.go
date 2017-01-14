package swaggerServer

import "time"
var _ time.Time

/* 
coordinates object */
type GetCharactersCharacterIdBookmarksCoordinates struct {
/*
	 x number */
	x float64 `json:"x,omitempty"`
/*
	 y number */
	y float64 `json:"y,omitempty"`
/*
	 z number */
	z float64 `json:"z,omitempty"`
}
