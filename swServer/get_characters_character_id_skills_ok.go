package swServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersCharacterIdSkillsOk struct {
/*
	 skills array */
	skills []GetCharactersCharacterIdSkillsSkill `json:"skills,omitempty"`
/*
	 total_sp integer */
	total_sp int64 `json:"total_sp,omitempty"`
}
