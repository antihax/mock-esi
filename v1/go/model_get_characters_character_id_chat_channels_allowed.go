package esiv1

/*
allowed object */
type GetCharactersCharacterIdChatChannelsAllowed struct {
	/*
	 ID of an allowed channel member */
	AccessorId int32 `json:"accessor_id,omitempty"`
	/*
	 accessor_type string */
	AccessorType string `json:"accessor_type,omitempty"`
}
