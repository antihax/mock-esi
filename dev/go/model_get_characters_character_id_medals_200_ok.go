package esidev

import "time"

/*
200 ok object
*/
type GetCharactersCharacterIdMedals200Ok struct {
	/*
	 corporation_id integer */
	CorporationId int32 `json:"corporation_id,omitempty"`
	/*
	 date string */
	Date time.Time `json:"date,omitempty"`
	/*
	 description string */
	Description string `json:"description,omitempty"`
	/*
	 graphics array */
	Graphics []GetCharactersCharacterIdMedalsGraphic `json:"graphics,omitempty"`
	/*
	 issuer_id integer */
	IssuerId int32 `json:"issuer_id,omitempty"`
	/*
	 medal_id integer */
	MedalId int32 `json:"medal_id,omitempty"`
	/*
	 reason string */
	Reason string `json:"reason,omitempty"`
	/*
	 status string */
	Status string `json:"status,omitempty"`
	/*
	 title string */
	Title string `json:"title,omitempty"`
}
