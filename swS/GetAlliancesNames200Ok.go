package swS

import "time"
var _ time.Time

/* 
200 ok object */
type GetAlliancesNames200Ok struct {
/*
	 alliance_id integer */
	alliance_id int32 `json:"alliance_id,omitempty"`
/*
	 alliance_name string */
	alliance_name string `json:"alliance_name,omitempty"`
}
