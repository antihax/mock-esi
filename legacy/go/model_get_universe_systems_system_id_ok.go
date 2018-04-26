package esilegacy

/*
200 ok object */
type GetUniverseSystemsSystemIdOk struct {
	/*
	 The constellation this solar system is in */
	ConstellationId int32 `json:"constellation_id,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 planets array */
	Planets []GetUniverseSystemsSystemIdPlanet `json:"planets,omitempty"`
	/* */
	Position GetUniverseSystemsSystemIdPosition `json:"position,omitempty"`
	/*
	 security_class string */
	SecurityClass string `json:"security_class,omitempty"`
	/*
	 security_status number */
	SecurityStatus float32 `json:"security_status,omitempty"`
	/*
	 star_id integer */
	StarId int32 `json:"star_id,omitempty"`
	/*
	 stargates array */
	Stargates []int32 `json:"stargates,omitempty"`
	/*
	 system_id integer */
	SystemId int32 `json:"system_id,omitempty"`
}
