package swS

import "time"
var _ time.Time

/* 
skill object */
type GetCharactersCharacterIdSkillsSkill struct {
/*
	 current_skill_level integer */
	current_skill_level int32 `json:"current_skill_level,omitempty"`
/*
	 skill_id integer */
	skill_id int32 `json:"skill_id,omitempty"`
/*
	 skillpoints_in_skill integer */
	skillpoints_in_skill int64 `json:"skillpoints_in_skill,omitempty"`
}
