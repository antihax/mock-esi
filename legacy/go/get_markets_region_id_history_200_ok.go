package esilegacy

/*
200 ok object */
type GetMarketsRegionIdHistory200Ok struct {
	/*
	 The date of this historical statistic entry */
	Date string `json:"date,omitempty"`
	/*
	 Total number of orders happened that day */
	OrderCount int64 `json:"order_count,omitempty"`
	/*
	 Total */
	Volume int64 `json:"volume,omitempty"`
	/*
	 highest number */
	Highest float64 `json:"highest,omitempty"`
	/*
	 average number */
	Average float64 `json:"average,omitempty"`
	/*
	 lowest number */
	Lowest float64 `json:"lowest,omitempty"`
}
