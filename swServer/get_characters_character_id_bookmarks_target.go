package swServer

import "time"
var _ time.Time

/* 
target object */
type GetCharactersCharacterIdBookmarksTarget struct {
/* */
	coordinates GetCharactersCharacterIdBookmarksCoordinates `json:"coordinates,omitempty"`
/* */
	item GetCharactersCharacterIdBookmarksItem `json:"item,omitempty"`
/*
	 location_id integer */
	location_id int64 `json:"location_id,omitempty"`
}
