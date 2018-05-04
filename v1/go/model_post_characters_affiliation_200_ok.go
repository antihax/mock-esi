package esiv1



/* 
200 ok object */
type PostCharactersAffiliation200Ok struct {
/*
	 The character's alliance ID, if their corporation is in an alliance */
	AllianceId int32 `json:"alliance_id,omitempty"`
/*
	 The character's ID */
	CharacterId int32 `json:"character_id,omitempty"`
/*
	 The character's corporation ID */
	CorporationId int32 `json:"corporation_id,omitempty"`
/*
	 The character's faction ID, if their corporation is in a faction */
	FactionId int32 `json:"faction_id,omitempty"`
}
