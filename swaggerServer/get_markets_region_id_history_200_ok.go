package swaggerServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetMarketsRegionIdHistory200Ok struct {
/*
	 average number */
	average float32 `json:"average,omitempty"`
/*
	 The date of this historical statistic entry */
	date time.Time `json:"date,omitempty"`
/*
	 highest number */
	highest float32 `json:"highest,omitempty"`
/*
	 lowest number */
	lowest float32 `json:"lowest,omitempty"`
/*
	 Total number of orders happened that day */
	order_count int64 `json:"order_count,omitempty"`
/*
	 Total */
	volume int64 `json:"volume,omitempty"`
}
