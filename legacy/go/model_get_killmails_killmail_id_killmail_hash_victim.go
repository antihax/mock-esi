package esilegacy



/* 
victim object */
type GetKillmailsKillmailIdKillmailHashVictim struct {
/*
	 alliance_id integer */
	AllianceId int32 `json:"alliance_id,omitempty"`
/*
	 character_id integer */
	CharacterId int32 `json:"character_id,omitempty"`
/*
	 corporation_id integer */
	CorporationId int32 `json:"corporation_id,omitempty"`
/*
	 How much total damage was taken by the victim  */
	DamageTaken int32 `json:"damage_taken,omitempty"`
/*
	 faction_id integer */
	FactionId int32 `json:"faction_id,omitempty"`
/*
	 items array */
	Items []GetKillmailsKillmailIdKillmailHashItem1 `json:"items,omitempty"`
/* */
	Position GetKillmailsKillmailIdKillmailHashPosition `json:"position,omitempty"`
/*
	 The ship that the victim was piloting and was destroyed  */
	ShipTypeId int32 `json:"ship_type_id,omitempty"`
}
