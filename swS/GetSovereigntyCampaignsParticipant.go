package swS

import "time"
var _ time.Time

/* 
participant object */
type GetSovereigntyCampaignsParticipant struct {
/*
	 alliance_id integer */
	alliance_id int32 `json:"alliance_id,omitempty"`
/*
	 score number */
	score float32 `json:"score,omitempty"`
}
