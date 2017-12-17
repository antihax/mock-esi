package esiv1

import "time"

/*
200 ok object */
type GetCorporationsCorporationIdBookmarks200Ok struct {
	/*
	 bookmark_id integer */
	BookmarkId int32 `json:"bookmark_id,omitempty"`
	/*
	 creator_id integer */
	CreatorId int32 `json:"creator_id,omitempty"`
	/*
	 folder_id integer */
	FolderId int32 `json:"folder_id,omitempty"`
	/*
	 created string */
	Created time.Time `json:"created,omitempty"`
	/*
	 label string */
	Label string `json:"label,omitempty"`
	/*
	 notes string */
	Notes string `json:"notes,omitempty"`
	/*
	 location_id integer */
	LocationId int32 `json:"location_id,omitempty"`
	/* */
	Item GetCorporationsCorporationIdBookmarksItem `json:"item,omitempty"`
	/* */
	Coordinates GetCorporationsCorporationIdBookmarksCoordinates `json:"coordinates,omitempty"`
}
