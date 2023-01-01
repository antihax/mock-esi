package esiv2

/*
200 ok object
*/
type GetUniverseMoonsMoonIdOk struct {
	/*
	 moon_id integer */
	MoonId int32 `json:"moon_id,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/* */
	Position GetUniverseMoonsMoonIdPosition `json:"position,omitempty"`
	/*
	 The solar system this moon is in */
	SystemId int32 `json:"system_id,omitempty"`
}
