package esilatest

/*
200 ok object */
type GetLoyaltyStoresCorporationIdOffers200Ok struct {
	/*
	 offer_id integer */
	OfferId int32 `json:"offer_id,omitempty"`
	/*
	 type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
	/*
	 quantity integer */
	Quantity int32 `json:"quantity,omitempty"`
	/*
	 lp_cost integer */
	LpCost int32 `json:"lp_cost,omitempty"`
	/*
	 isk_cost number */
	IskCost float32 `json:"isk_cost,omitempty"`
	/*
	 required_items array */
	RequiredItems []GetLoyaltyStoresCorporationIdOffersRequiredItem `json:"required_items,omitempty"`
}
