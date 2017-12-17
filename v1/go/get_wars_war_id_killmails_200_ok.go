package esiv1

/*
200 ok object */
type GetWarsWarIdKillmails200Ok struct {
	/*
	 ID of this killmail */
	KillmailId int32 `json:"killmail_id,omitempty"`
	/*
	 A hash of this killmail */
	KillmailHash string `json:"killmail_hash,omitempty"`
}
