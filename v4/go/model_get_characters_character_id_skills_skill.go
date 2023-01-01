package esiv4

/*
skill object
*/
type GetCharactersCharacterIdSkillsSkill struct {
	/*
	 active_skill_level integer */
	ActiveSkillLevel int32 `json:"active_skill_level,omitempty"`
	/*
	 skill_id integer */
	SkillId int32 `json:"skill_id,omitempty"`
	/*
	 skillpoints_in_skill integer */
	SkillpointsInSkill int64 `json:"skillpoints_in_skill,omitempty"`
	/*
	 trained_skill_level integer */
	TrainedSkillLevel int32 `json:"trained_skill_level,omitempty"`
}
