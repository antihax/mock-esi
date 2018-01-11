package esilatest

/*
200 ok object */
type GetUniverseSystemsSystemIdOk struct {
	/*
	 star_id integer */
	StarId int32 `json:"star_id,omitempty"`
	/*
	 system_id integer */
	SystemId int32 `json:"system_id,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/* */
	Position GetUniverseSystemsSystemIdPosition `json:"position,omitempty"`
	/*
	 security_status number */
	SecurityStatus float32 `json:"security_status,omitempty"`
	/*
	 security_class string */
	SecurityClass string `json:"security_class,omitempty"`
	/*
	 The constellation this solar system is in */
	ConstellationId int32 `json:"constellation_id,omitempty"`
	/*
	 planets array */
	Planets []GetUniverseSystemsSystemIdPlanet `json:"planets,omitempty"`
	/*
	 stargates array */
	Stargates []int32 `json:"stargates,omitempty"`
	/*
	 stations array */
	Stations []int32 `json:"stations,omitempty"`
}
