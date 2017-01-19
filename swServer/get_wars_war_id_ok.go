package swServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetWarsWarIdOk struct {
/* */
	aggressor GetWarsWarIdAggressor `json:"aggressor,omitempty"`
/*
	 allied corporations or alliances, each object contains either corporation_id or alliance_id */
	allies []GetWarsWarIdAlly `json:"allies,omitempty"`
/*
	 Time that the war was declared */
	declared time.Time `json:"declared,omitempty"`
/* */
	defender GetWarsWarIdDefender `json:"defender,omitempty"`
/*
	 Time the war ended and shooting was no longer allowed */
	finished time.Time `json:"finished,omitempty"`
/*
	 ID of the specified war */
	id int32 `json:"id,omitempty"`
/*
	 Was the war declared mutual by both parties */
	mutual bool `json:"mutual,omitempty"`
/*
	 Is the war currently open for allies or not */
	open_for_allies bool `json:"open_for_allies,omitempty"`
/*
	 Time the war was retracted but both sides could still shoot each other */
	retracted time.Time `json:"retracted,omitempty"`
/*
	 Time when the war started and both sides could shoot each other */
	started time.Time `json:"started,omitempty"`
}
