package esilegacy

import "time"

/*
200 ok object
*/
type GetWarsWarIdOk struct {
	/* */
	Aggressor GetWarsWarIdAggressor `json:"aggressor,omitempty"`
	/*
	 allied corporations or alliances, each object contains either corporation_id or alliance_id */
	Allies []GetWarsWarIdAlly `json:"allies,omitempty"`
	/*
	 Time that the war was declared */
	Declared time.Time `json:"declared,omitempty"`
	/* */
	Defender GetWarsWarIdDefender `json:"defender,omitempty"`
	/*
	 Time the war ended and shooting was no longer allowed */
	Finished time.Time `json:"finished,omitempty"`
	/*
	 ID of the specified war */
	Id int32 `json:"id,omitempty"`
	/*
	 Was the war declared mutual by both parties */
	Mutual bool `json:"mutual,omitempty"`
	/*
	 Is the war currently open for allies or not */
	OpenForAllies bool `json:"open_for_allies,omitempty"`
	/*
	 Time the war was retracted but both sides could still shoot each other */
	Retracted time.Time `json:"retracted,omitempty"`
	/*
	 Time when the war started and both sides could shoot each other */
	Started time.Time `json:"started,omitempty"`
}
