package esiv2

/*
planet object */
type GetUniverseSystemsSystemIdPlanet struct {
	/*
	 moons array */
	Moons []int32 `json:"moons,omitempty"`
	/*
	 planet_id integer */
	PlanetId int32 `json:"planet_id,omitempty"`
}
