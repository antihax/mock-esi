package swServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersCharacterIdKillmailsRecent200Ok struct {
/*
	 A hash of this killmail */
	killmail_hash string `json:"killmail_hash,omitempty"`
/*
	 ID of this killmail */
	killmail_id int32 `json:"killmail_id,omitempty"`
}
