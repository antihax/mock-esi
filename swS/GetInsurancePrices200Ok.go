package swS

import "time"
var _ time.Time

/* 
200 ok object */
type GetInsurancePrices200Ok struct {
/*
	 A list of a available insurance levels for this ship type */
	levels []GetInsurancePricesLevel `json:"levels,omitempty"`
/*
	 type_id integer */
	type_id int32 `json:"type_id,omitempty"`
}
