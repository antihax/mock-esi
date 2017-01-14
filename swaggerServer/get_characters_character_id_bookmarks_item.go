package swaggerServer

import "time"
var _ time.Time

/* 
item object */
type GetCharactersCharacterIdBookmarksItem struct {
/*
	 item_id integer */
	item_id int64 `json:"item_id,omitempty"`
/*
	 type_id integer */
	type_id int32 `json:"type_id,omitempty"`
}
