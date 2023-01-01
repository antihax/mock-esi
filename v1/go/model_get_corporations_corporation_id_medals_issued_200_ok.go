package esiv1

import "time"

/*
200 ok object
*/
type GetCorporationsCorporationIdMedalsIssued200Ok struct {
	/*
	 ID of the character who was rewarded this medal */
	CharacterId int32 `json:"character_id,omitempty"`
	/*
	 issued_at string */
	IssuedAt time.Time `json:"issued_at,omitempty"`
	/*
	 ID of the character who issued the medal */
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
}
