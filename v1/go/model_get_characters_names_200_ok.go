package esiv1

/*
200 ok object */
type GetCharactersNames200Ok struct {
	/*
	 character_id integer */
	CharacterId int64 `json:"character_id,omitempty"`
	/*
	 character_name string */
	CharacterName string `json:"character_name,omitempty"`
}
