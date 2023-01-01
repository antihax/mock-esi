package esiv2

/*
200 ok object
*/
type GetUniverseAsteroidBeltsAsteroidBeltIdOk struct {
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/* */
	Position GetUniverseAsteroidBeltsAsteroidBeltIdPosition `json:"position,omitempty"`
	/*
	 The solar system this asteroid belt is in */
	SystemId int32 `json:"system_id,omitempty"`
}
