package esilegacy

/*
200 ok object */
type GetCharactersCharacterIdLoyaltyPoints200Ok struct {
	/*
	 corporation_id integer */
	CorporationId int32 `json:"corporation_id,omitempty"`
	/*
	 loyalty_points integer */
	LoyaltyPoints int32 `json:"loyalty_points,omitempty"`
}
