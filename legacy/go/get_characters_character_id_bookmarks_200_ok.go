package esilegacy

import "time"

/*
200 ok object */
type GetCharactersCharacterIdBookmarks200Ok struct {
	/*
	 bookmark_id integer */
	BookmarkId int64 `json:"bookmark_id,omitempty"`
	/*
	 creator_id integer */
	CreatorId int32 `json:"creator_id,omitempty"`
	/*
	 owner_id integer */
	OwnerId int32 `json:"owner_id,omitempty"`
	/*
	 folder_id integer */
	FolderId int32 `json:"folder_id,omitempty"`
	/*
	 create_date string */
	CreateDate time.Time `json:"create_date,omitempty"`
	/*
	 memo string */
	Memo string `json:"memo,omitempty"`
	/*
	 note string */
	Note string `json:"note,omitempty"`
	/* */
	Target GetCharactersCharacterIdBookmarksTarget `json:"target,omitempty"`
}
