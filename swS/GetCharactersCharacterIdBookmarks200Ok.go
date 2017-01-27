package swS

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersCharacterIdBookmarks200Ok struct {
/*
	 bookmark_id integer */
	bookmark_id int64 `json:"bookmark_id,omitempty"`
/*
	 create_date string */
	create_date time.Time `json:"create_date,omitempty"`
/*
	 creator_id integer */
	creator_id int32 `json:"creator_id,omitempty"`
/*
	 folder_id integer */
	folder_id int32 `json:"folder_id,omitempty"`
/*
	 memo string */
	memo string `json:"memo,omitempty"`
/*
	 note string */
	note string `json:"note,omitempty"`
/*
	 owner_id integer */
	owner_id int32 `json:"owner_id,omitempty"`
/* */
	target GetCharactersCharacterIdBookmarksTarget `json:"target,omitempty"`
}
