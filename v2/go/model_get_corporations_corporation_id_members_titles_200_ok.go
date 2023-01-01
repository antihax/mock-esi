package esiv2

/*
200 ok object
*/
type GetCorporationsCorporationIdMembersTitles200Ok struct {
	/*
	 character_id integer */
	CharacterId int32 `json:"character_id,omitempty"`
	/*
	 A list of title_id */
	Titles []int32 `json:"titles,omitempty"`
}
