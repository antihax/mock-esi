package esilatest

import "time"

/*
200 ok object */
type GetCharactersCharacterIdOpportunities200Ok struct {
	/*
	 task_id integer */
	TaskId int32 `json:"task_id,omitempty"`
	/*
	 completed_at string */
	CompletedAt time.Time `json:"completed_at,omitempty"`
}
