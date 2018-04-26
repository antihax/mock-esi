package esidev

import "time"

/*
200 ok object */
type GetCharactersCharacterIdOpportunities200Ok struct {
	/*
	 completed_at string */
	CompletedAt time.Time `json:"completed_at,omitempty"`
	/*
	 task_id integer */
	TaskId int32 `json:"task_id,omitempty"`
}
