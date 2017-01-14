package swaggerServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersCharacterIdPlanetsPlanetIdOk struct {
/*
	 links array */
	links []GetCharactersCharacterIdPlanetsPlanetIdLink `json:"links,omitempty"`
/*
	 pins array */
	pins []GetCharactersCharacterIdPlanetsPlanetIdPin `json:"pins,omitempty"`
/*
	 routes array */
	routes []GetCharactersCharacterIdPlanetsPlanetIdRoute `json:"routes,omitempty"`
}
