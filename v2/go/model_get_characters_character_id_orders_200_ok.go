package esiv2

import "time"

/*
200 ok object */
type GetCharactersCharacterIdOrders200Ok struct {
	/*
	 Unique order ID */
	OrderId int64 `json:"order_id,omitempty"`
	/*
	 The type ID of the item transacted in this order */
	TypeId int32 `json:"type_id,omitempty"`
	/*
	 ID of the region where order was placed */
	RegionId int32 `json:"region_id,omitempty"`
	/*
	 ID of the location where order was placed */
	LocationId int64 `json:"location_id,omitempty"`
	/*
	 Valid order range, numbers are ranges in jumps */
	Range_ string `json:"range,omitempty"`
	/*
	 Cost per unit for this order */
	Price float64 `json:"price,omitempty"`
	/*
	 Quantity of items required or offered at time order was placed */
	VolumeTotal int32 `json:"volume_total,omitempty"`
	/*
	 Quantity of items still required or offered */
	VolumeRemain int32 `json:"volume_remain,omitempty"`
	/*
	 Date and time when this order was issued */
	Issued time.Time `json:"issued,omitempty"`
	/*
	 True if the order is a bid (buy) order */
	IsBuyOrder bool `json:"is_buy_order,omitempty"`
	/*
	 For buy orders, the minimum quantity that will be accepted in a matching sell order */
	MinVolume int32 `json:"min_volume,omitempty"`
	/*
	 For buy orders, the amount of ISK in escrow */
	Escrow float64 `json:"escrow,omitempty"`
	/*
	 Number of days for which order is valid (starting from the issued date). An order expires at time issued + duration */
	Duration int32 `json:"duration,omitempty"`
	/*
	 Signifies whether the buy/sell order was placed on behalf of a corporation. */
	IsCorporation bool `json:"is_corporation,omitempty"`
}
