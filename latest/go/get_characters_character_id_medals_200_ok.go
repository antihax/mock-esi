package esilatest

import "time"

/*
200 ok object */
type GetCharactersCharacterIdMedals200Ok struct {
	/*
	 medal_id integer */
	MedalId int32 `json:"medal_id,omitempty"`
	/*
	 title string */
	Title string `json:"title,omitempty"`
	/*
	 description string */
	Description string `json:"description,omitempty"`
	/*
	 corporation_id integer */
	CorporationId int32 `json:"corporation_id,omitempty"`
	/*
	 issuer_id integer */
	IssuerId int32 `json:"issuer_id,omitempty"`
	/*
	 date string */
	Date time.Time `json:"date,omitempty"`
	/*
	 reason string */
	Reason string `json:"reason,omitempty"`
	/*
	 status string */
	Status string `json:"status,omitempty"`
	/*
	 graphics array */
	Graphics []GetCharactersCharacterIdMedalsGraphic `json:"graphics,omitempty"`
}
