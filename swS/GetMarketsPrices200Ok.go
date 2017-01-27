package swS

import "time"
var _ time.Time

/* 
200 ok object */
type GetMarketsPrices200Ok struct {
/*
	 adjusted_price number */
	adjusted_price float32 `json:"adjusted_price,omitempty"`
/*
	 average_price number */
	average_price float32 `json:"average_price,omitempty"`
/*
	 type_id integer */
	type_id int32 `json:"type_id,omitempty"`
}
