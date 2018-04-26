package esidev

import "time"

/*
200 ok object */
type GetCharactersCharacterIdOk struct {
	/*
	 The character's alliance ID */
	AllianceId int32 `json:"alliance_id,omitempty"`
	/*
	 ancestry_id integer */
	AncestryId int32 `json:"ancestry_id,omitempty"`
	/*
	 Creation date of the character */
	Birthday time.Time `json:"birthday,omitempty"`
	/*
	 bloodline_id integer */
	BloodlineId int32 `json:"bloodline_id,omitempty"`
	/*
	 The character's corporation ID */
	CorporationId int32 `json:"corporation_id,omitempty"`
	/*
	 description string */
	Description string `json:"description,omitempty"`
	/*
	 ID of the faction the character is fighting for, if the character is enlisted in Factional Warfare */
	FactionId int32 `json:"faction_id,omitempty"`
	/*
	 gender string */
	Gender string `json:"gender,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 race_id integer */
	RaceId int32 `json:"race_id,omitempty"`
	/*
	 security_status number */
	SecurityStatus float32 `json:"security_status,omitempty"`
}
