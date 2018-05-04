package esilatest

import "time"

/*
200 ok object */
type GetCorporationsCorporationIdBookmarks200Ok struct {
	/*
	 bookmark_id integer */
	BookmarkId int32 `json:"bookmark_id,omitempty"`
	/* */
	Coordinates GetCorporationsCorporationIdBookmarksCoordinates `json:"coordinates,omitempty"`
	/*
	 created string */
	Created time.Time `json:"created,omitempty"`
	/*
	 creator_id integer */
	CreatorId int32 `json:"creator_id,omitempty"`
	/*
	 folder_id integer */
	FolderId int32 `json:"folder_id,omitempty"`
	/* */
	Item GetCorporationsCorporationIdBookmarksItem `json:"item,omitempty"`
	/*
	 label string */
	Label string `json:"label,omitempty"`
	/*
	 location_id integer */
	LocationId int32 `json:"location_id,omitempty"`
	/*
	 notes string */
	Notes string `json:"notes,omitempty"`
}
