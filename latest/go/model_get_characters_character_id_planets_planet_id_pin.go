package esilatest

import "time"

/*
pin object
*/
type GetCharactersCharacterIdPlanetsPlanetIdPin struct {
	/*
	 contents array */
	Contents []GetCharactersCharacterIdPlanetsPlanetIdContent `json:"contents,omitempty"`
	/*
	 expiry_time string */
	ExpiryTime time.Time `json:"expiry_time,omitempty"`
	/* */
	ExtractorDetails GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails `json:"extractor_details,omitempty"`
	/* */
	FactoryDetails GetCharactersCharacterIdPlanetsPlanetIdFactoryDetails `json:"factory_details,omitempty"`
	/*
	 install_time string */
	InstallTime time.Time `json:"install_time,omitempty"`
	/*
	 last_cycle_start string */
	LastCycleStart time.Time `json:"last_cycle_start,omitempty"`
	/*
	 latitude number */
	Latitude float32 `json:"latitude,omitempty"`
	/*
	 longitude number */
	Longitude float32 `json:"longitude,omitempty"`
	/*
	 pin_id integer */
	PinId int64 `json:"pin_id,omitempty"`
	/*
	 schematic_id integer */
	SchematicId int32 `json:"schematic_id,omitempty"`
	/*
	 type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
}
