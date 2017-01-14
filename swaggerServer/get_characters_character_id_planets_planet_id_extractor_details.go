package swaggerServer

import "time"
var _ time.Time

/* 
extractor_details object */
type GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails struct {
/*
	 in seconds */
	cycle_time int32 `json:"cycle_time,omitempty"`
/*
	 head_radius number */
	head_radius float32 `json:"head_radius,omitempty"`
/*
	 heads array */
	heads []GetCharactersCharacterIdPlanetsPlanetIdHead `json:"heads,omitempty"`
/*
	 product_type_id integer */
	product_type_id int32 `json:"product_type_id,omitempty"`
/*
	 qty_per_cycle integer */
	qty_per_cycle int32 `json:"qty_per_cycle,omitempty"`
}
