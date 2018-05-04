package esilegacy

import "time"

/*
blocked object */
type GetCharactersCharacterIdChatChannelsBlocked struct {
	/*
	 ID of a blocked channel member */
	AccessorId int32 `json:"accessor_id,omitempty"`
	/*
	 accessor_type string */
	AccessorType string `json:"accessor_type,omitempty"`
	/*
	 Time at which this accessor will no longer be blocked */
	EndAt time.Time `json:"end_at,omitempty"`
	/*
	 Reason this accessor is blocked */
	Reason string `json:"reason,omitempty"`
}
