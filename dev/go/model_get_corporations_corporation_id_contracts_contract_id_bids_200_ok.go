package esidev

import "time"

/*
200 ok object */
type GetCorporationsCorporationIdContractsContractIdBids200Ok struct {
	/*
	 The amount bid, in ISK */
	Amount float32 `json:"amount,omitempty"`
	/*
	 Unique ID for the bid */
	BidId int32 `json:"bid_id,omitempty"`
	/*
	 Character ID of the bidder */
	BidderId int32 `json:"bidder_id,omitempty"`
	/*
	 Datetime when the bid was placed */
	DateBid time.Time `json:"date_bid,omitempty"`
}
