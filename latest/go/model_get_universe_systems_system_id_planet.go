package esilatest

/*
planet object */
type GetUniverseSystemsSystemIdPlanet struct {
	/*
	 planet_id integer */
	PlanetId int32 `json:"planet_id,omitempty"`
	/*
	 moons array */
	Moons []int32 `json:"moons,omitempty"`
	/*
	 asteroid_belts array */
	AsteroidBelts []int32 `json:"asteroid_belts,omitempty"`
}
