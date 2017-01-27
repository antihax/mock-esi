package swS

import "time"
var _ time.Time

/* 
200 ok object */
type GetMarketsStructuresStructureId200Ok struct {
/*
	 duration integer */
	duration int32 `json:"duration,omitempty"`
/*
	 is_buy_order boolean */
	is_buy_order bool `json:"is_buy_order,omitempty"`
/*
	 issued string */
	issued time.Time `json:"issued,omitempty"`
/*
	 location_id integer */
	location_id int64 `json:"location_id,omitempty"`
/*
	 min_volume integer */
	min_volume int32 `json:"min_volume,omitempty"`
/*
	 order_id integer */
	order_id int64 `json:"order_id,omitempty"`
/*
	 price number */
	price float32 `json:"price,omitempty"`
/*
	 range string */
	_range string `json:"range,omitempty"`
/*
	 type_id integer */
	type_id int32 `json:"type_id,omitempty"`
/*
	 volume_remain integer */
	volume_remain int32 `json:"volume_remain,omitempty"`
/*
	 volume_total integer */
	volume_total int32 `json:"volume_total,omitempty"`
}
