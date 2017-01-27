package swS

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersCharacterIdBookmarksFolders200Ok struct {
/*
	 folder_id integer */
	folder_id int32 `json:"folder_id,omitempty"`
/*
	 name string */
	name string `json:"name,omitempty"`
/*
	 owner_id integer */
	owner_id int32 `json:"owner_id,omitempty"`
}
