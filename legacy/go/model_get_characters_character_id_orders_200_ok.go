package esilegacy



/* 
200 ok object */
type GetCharactersCharacterIdOrders200Ok struct {
/*
	 Wallet division for the buyer or seller of this order. Always 1000 for characters. Currently 1000 through 1006 for corporations */
	AccountId int32 `json:"account_id,omitempty"`
/*
	 Number of days the order is valid for (starting from the issued date). An order expires at time issued + duration */
	Duration int32 `json:"duration,omitempty"`
/*
	 For buy orders, the amount of ISK in escrow */
	Escrow float64 `json:"escrow,omitempty"`
/*
	 True for a bid (buy) order. False for an offer (sell) order */
	IsBuyOrder bool `json:"is_buy_order,omitempty"`
/*
	 is_corp boolean */
	IsCorp bool `json:"is_corp,omitempty"`
/*
	 Date and time when this order was issued */
	Issued time.Time `json:"issued,omitempty"`
/*
	 ID of the location where order was placed */
	LocationId int64 `json:"location_id,omitempty"`
/*
	 For bids (buy orders), the minimum quantity that will be accepted in a matching offer (sell order) */
	MinVolume int32 `json:"min_volume,omitempty"`
/*
	 Unique order ID */
	OrderId int64 `json:"order_id,omitempty"`
/*
	 Cost per unit for this order */
	Price float64 `json:"price,omitempty"`
/*
	 Valid order range, numbers are ranges in jumps */
	Range_ string `json:"range,omitempty"`
/*
	 ID of the region where order was placed */
	RegionId int32 `json:"region_id,omitempty"`
/*
	 Current order state */
	State string `json:"state,omitempty"`
/*
	 The type ID of the item transacted in this order */
	TypeId int32 `json:"type_id,omitempty"`
/*
	 Quantity of items still required or offered */
	VolumeRemain int32 `json:"volume_remain,omitempty"`
/*
	 Quantity of items required or offered at time order was placed */
	VolumeTotal int32 `json:"volume_total,omitempty"`
}
