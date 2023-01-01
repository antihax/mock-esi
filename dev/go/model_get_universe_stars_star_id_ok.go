package esidev

/*
200 ok object
*/
type GetUniverseStarsStarIdOk struct {
	/*
	 Age of star in years */
	Age int64 `json:"age,omitempty"`
	/*
	 luminosity number */
	Luminosity float32 `json:"luminosity,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 radius integer */
	Radius int64 `json:"radius,omitempty"`
	/*
	 solar_system_id integer */
	SolarSystemId int32 `json:"solar_system_id,omitempty"`
	/*
	 spectral_class string */
	SpectralClass string `json:"spectral_class,omitempty"`
	/*
	 temperature integer */
	Temperature int32 `json:"temperature,omitempty"`
	/*
	 type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
}
