package esilegacy

/*
target object
*/
type GetCharactersCharacterIdBookmarksTarget struct {
	/* */
	Coordinates GetCharactersCharacterIdBookmarksCoordinates `json:"coordinates,omitempty"`
	/* */
	Item GetCharactersCharacterIdBookmarksItem `json:"item,omitempty"`
	/*
	 location_id integer */
	LocationId int64 `json:"location_id,omitempty"`
}
