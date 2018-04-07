package esidev

import "time"

/*
200 ok object */
type GetMarketsRegionIdOrders200Ok struct {
	/*
	 order_id integer */
	OrderId int64 `json:"order_id,omitempty"`
	/*
	 type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
	/*
	 location_id integer */
	LocationId int64 `json:"location_id,omitempty"`
	/*
	 The solar system this order was placed */
	SystemId int32 `json:"system_id,omitempty"`
	/*
	 volume_total integer */
	VolumeTotal int32 `json:"volume_total,omitempty"`
	/*
	 volume_remain integer */
	VolumeRemain int32 `json:"volume_remain,omitempty"`
	/*
	 min_volume integer */
	MinVolume int32 `json:"min_volume,omitempty"`
	/*
	 price number */
	Price float64 `json:"price,omitempty"`
	/*
	 is_buy_order boolean */
	IsBuyOrder bool `json:"is_buy_order,omitempty"`
	/*
	 duration integer */
	Duration int32 `json:"duration,omitempty"`
	/*
	 issued string */
	Issued time.Time `json:"issued,omitempty"`
	/*
	 range string */
	Range_ string `json:"range,omitempty"`
}
