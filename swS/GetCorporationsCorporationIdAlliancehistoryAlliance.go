package swS

import "time"
var _ time.Time

/* 
alliance object */
type GetCorporationsCorporationIdAlliancehistoryAlliance struct {
/*
	 alliance_id integer */
	alliance_id int32 `json:"alliance_id,omitempty"`
/*
	 True if the alliance has been deleted */
	is_deleted bool `json:"is_deleted,omitempty"`
}
