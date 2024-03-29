package esilatest

import "time"

/*
200 ok object
*/
type GetCharactersCharacterIdPlanets200Ok struct {
	/*
	 last_update string */
	LastUpdate time.Time `json:"last_update,omitempty"`
	/*
	 num_pins integer */
	NumPins int32 `json:"num_pins,omitempty"`
	/*
	 owner_id integer */
	OwnerId int32 `json:"owner_id,omitempty"`
	/*
	 planet_id integer */
	PlanetId int32 `json:"planet_id,omitempty"`
	/*
	 planet_type string */
	PlanetType string `json:"planet_type,omitempty"`
	/*
	 solar_system_id integer */
	SolarSystemId int32 `json:"solar_system_id,omitempty"`
	/*
	 upgrade_level integer */
	UpgradeLevel int32 `json:"upgrade_level,omitempty"`
}
