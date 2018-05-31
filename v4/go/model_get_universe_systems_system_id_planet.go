package esiv4

/*
planet object */
type GetUniverseSystemsSystemIdPlanet struct {
	/*
	 asteroid_belts array */
	AsteroidBelts []int32 `json:"asteroid_belts,omitempty"`
	/*
	 moons array */
	Moons []int32 `json:"moons,omitempty"`
	/*
	 planet_id integer */
	PlanetId int32 `json:"planet_id,omitempty"`
}
