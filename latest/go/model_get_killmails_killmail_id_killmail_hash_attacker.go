package esilatest



/* 
attacker object */
type GetKillmailsKillmailIdKillmailHashAttacker struct {
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
	 damage_done integer */
	DamageDone int32 `json:"damage_done,omitempty"`
/*
	 faction_id integer */
	FactionId int32 `json:"faction_id,omitempty"`
/*
	 Was the attacker the one to achieve the final blow  */
	FinalBlow bool `json:"final_blow,omitempty"`
/*
	 Security status for the attacker  */
	SecurityStatus float32 `json:"security_status,omitempty"`
/*
	 What ship was the attacker flying  */
	ShipTypeId int32 `json:"ship_type_id,omitempty"`
/*
	 What weapon was used by the attacker for the kill  */
	WeaponTypeId int32 `json:"weapon_type_id,omitempty"`
}
