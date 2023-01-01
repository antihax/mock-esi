package esidev

/*
200 ok object
*/
type GetMarketsPrices200Ok struct {
	/*
	 adjusted_price number */
	AdjustedPrice float64 `json:"adjusted_price,omitempty"`
	/*
	 average_price number */
	AveragePrice float64 `json:"average_price,omitempty"`
	/*
	 type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
}
