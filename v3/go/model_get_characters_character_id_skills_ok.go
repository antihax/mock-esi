package esiv3

/*
200 ok object
*/
type GetCharactersCharacterIdSkillsOk struct {
	/*
	 skills array */
	Skills []GetCharactersCharacterIdSkillsSkill `json:"skills,omitempty"`
	/*
	 total_sp integer */
	TotalSp int64 `json:"total_sp,omitempty"`
}
