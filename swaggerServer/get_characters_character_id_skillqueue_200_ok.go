package swaggerServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersCharacterIdSkillqueue200Ok struct {
/*
	 finish_date string */
	finish_date time.Time `json:"finish_date,omitempty"`
/*
	 finished_level integer */
	finished_level int32 `json:"finished_level,omitempty"`
/*
	 level_end_sp integer */
	level_end_sp int32 `json:"level_end_sp,omitempty"`
/*
	 Amount of SP that was in the skill when it started training it's current level. Used to calculate % of current level complete. */
	level_start_sp int32 `json:"level_start_sp,omitempty"`
/*
	 queue_position integer */
	queue_position int32 `json:"queue_position,omitempty"`
/*
	 skill_id integer */
	skill_id int32 `json:"skill_id,omitempty"`
/*
	 start_date string */
	start_date time.Time `json:"start_date,omitempty"`
/*
	 training_start_sp integer */
	training_start_sp int32 `json:"training_start_sp,omitempty"`
}
