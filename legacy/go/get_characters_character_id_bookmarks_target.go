package esilegacy

/*
target object */
type GetCharactersCharacterIdBookmarksTarget struct {
	/*
	 location_id integer */
	LocationId int64 `json:"location_id,omitempty"`
	/* */
	Item GetCharactersCharacterIdBookmarksItem `json:"item,omitempty"`
	/* */
	Coordinates GetCharactersCharacterIdBookmarksCoordinates `json:"coordinates,omitempty"`
}
