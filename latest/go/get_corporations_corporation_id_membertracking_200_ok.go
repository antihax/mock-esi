package esilatest

import "time"

/*
200 ok object */
type GetCorporationsCorporationIdMembertracking200Ok struct {
	/*
	 character_id integer */
	CharacterId int32 `json:"character_id,omitempty"`
	/*
	 start_date string */
	StartDate time.Time `json:"start_date,omitempty"`
	/*
	 base_id integer */
	BaseId int32 `json:"base_id,omitempty"`
	/*
	 logon_date string */
	LogonDate time.Time `json:"logon_date,omitempty"`
	/*
	 logoff_date string */
	LogoffDate time.Time `json:"logoff_date,omitempty"`
	/*
	 location_id integer */
	LocationId int64 `json:"location_id,omitempty"`
	/*
	 ship_type_id integer */
	ShipTypeId int32 `json:"ship_type_id,omitempty"`
}
