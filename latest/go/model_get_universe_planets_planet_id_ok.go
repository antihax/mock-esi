package esilatest

/*
200 ok object */
type GetUniversePlanetsPlanetIdOk struct {
	/*
	 planet_id integer */
	PlanetId int32 `json:"planet_id,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
	/* */
	Position GetUniversePlanetsPlanetIdPosition `json:"position,omitempty"`
	/*
	 The solar system this planet is in */
	SystemId int32 `json:"system_id,omitempty"`
}
