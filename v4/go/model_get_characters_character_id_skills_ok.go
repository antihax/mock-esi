package esiv4



/* 
200 ok object */
type GetCharactersCharacterIdSkillsOk struct {
/*
	 skills array */
	Skills []GetCharactersCharacterIdSkillsSkill `json:"skills,omitempty"`
/*
	 total_sp integer */
	TotalSp int64 `json:"total_sp,omitempty"`
/*
	 Skill points available to be assigned */
	UnallocatedSp int32 `json:"unallocated_sp,omitempty"`
}
