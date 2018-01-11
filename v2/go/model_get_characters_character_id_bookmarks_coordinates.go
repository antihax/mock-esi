package esiv2

/*
Optional object that is returned if a bookmark was made on a planet or a random location in space. */
type GetCharactersCharacterIdBookmarksCoordinates struct {
	/*
	 x number */
	X float64 `json:"x,omitempty"`
	/*
	 y number */
	Y float64 `json:"y,omitempty"`
	/*
	 z number */
	Z float64 `json:"z,omitempty"`
}
