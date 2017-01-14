package swaggerServer

import "time"
var _ time.Time

/* 
head object */
type GetCharactersCharacterIdPlanetsPlanetIdHead struct {
/*
	 head_id integer */
	head_id int32 `json:"head_id,omitempty"`
/*
	 latitude number */
	latitude float32 `json:"latitude,omitempty"`
/*
	 longitude number */
	longitude float32 `json:"longitude,omitempty"`
}
