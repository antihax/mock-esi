package swServer

import "time"
var _ time.Time

/* 
ally object */
type GetWarsWarIdAlly struct {
/*
	 Alliance ID if and only if this ally is an alliance */
	alliance_id int32 `json:"alliance_id,omitempty"`
/*
	 Corporation ID if and only if this ally is a corporation */
	corporation_id int32 `json:"corporation_id,omitempty"`
}
