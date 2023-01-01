package esiv1

import "time"

/*
200 ok object
*/
type GetMarketsStructuresStructureId200Ok struct {
	/*
	 duration integer */
	Duration int32 `json:"duration,omitempty"`
	/*
	 is_buy_order boolean */
	IsBuyOrder bool `json:"is_buy_order,omitempty"`
	/*
	 issued string */
	Issued time.Time `json:"issued,omitempty"`
	/*
	 location_id integer */
	LocationId int64 `json:"location_id,omitempty"`
	/*
	 min_volume integer */
	MinVolume int32 `json:"min_volume,omitempty"`
	/*
	 order_id integer */
	OrderId int64 `json:"order_id,omitempty"`
	/*
	 price number */
	Price float64 `json:"price,omitempty"`
	/*
	 range string */
	Range_ string `json:"range,omitempty"`
	/*
	 type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
	/*
	 volume_remain integer */
	VolumeRemain int32 `json:"volume_remain,omitempty"`
	/*
	 volume_total integer */
	VolumeTotal int32 `json:"volume_total,omitempty"`
}
