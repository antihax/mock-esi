package esidev

/*
planet object */
type GetUniverseSystemsSystemIdPlanet struct {
	/*
	 planet_id integer */
	PlanetId int32 `json:"planet_id,omitempty"`
	/*
	 moons array */
	Moons []int32 `json:"moons,omitempty"`
}
