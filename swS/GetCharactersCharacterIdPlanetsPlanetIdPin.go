package swS

import "time"
var _ time.Time

/* 
pin object */
type GetCharactersCharacterIdPlanetsPlanetIdPin struct {
/*
	 expiry_time string */
	expiry_time time.Time `json:"expiry_time,omitempty"`
/* */
	extractor_details GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails `json:"extractor_details,omitempty"`
/* */
	factory_details GetCharactersCharacterIdPlanetsPlanetIdFactoryDetails `json:"factory_details,omitempty"`
/*
	 install_time string */
	install_time time.Time `json:"install_time,omitempty"`
/*
	 last_cycle_start string */
	last_cycle_start time.Time `json:"last_cycle_start,omitempty"`
/*
	 latitude number */
	latitude float32 `json:"latitude,omitempty"`
/*
	 longitude number */
	longitude float32 `json:"longitude,omitempty"`
/*
	 pin_id integer */
	pin_id int64 `json:"pin_id,omitempty"`
/*
	 schematic_id integer */
	schematic_id int32 `json:"schematic_id,omitempty"`
/*
	 type_id integer */
	type_id int32 `json:"type_id,omitempty"`
}
