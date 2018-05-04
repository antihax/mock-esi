package esiv1



/* 
muted object */
type GetCharactersCharacterIdChatChannelsMuted struct {
/*
	 ID of a muted channel member */
	AccessorId int32 `json:"accessor_id,omitempty"`
/*
	 accessor_type string */
	AccessorType string `json:"accessor_type,omitempty"`
/*
	 Time at which this accessor will no longer be muted */
	EndAt time.Time `json:"end_at,omitempty"`
/*
	 Reason this accessor is muted */
	Reason string `json:"reason,omitempty"`
}
