package esilatest

/*
200 ok object */
type GetMarketsRegionIdHistory200Ok struct {
	/*
	 average number */
	Average float64 `json:"average,omitempty"`
	/*
	 The date of this historical statistic entry */
	Date string `json:"date,omitempty"`
	/*
	 highest number */
	Highest float64 `json:"highest,omitempty"`
	/*
	 lowest number */
	Lowest float64 `json:"lowest,omitempty"`
	/*
	 Total number of orders happened that day */
	OrderCount int64 `json:"order_count,omitempty"`
	/*
	 Total */
	Volume int64 `json:"volume,omitempty"`
}
