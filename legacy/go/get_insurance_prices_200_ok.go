package esilegacy

/*
200 ok object */
type GetInsurancePrices200Ok struct {
	/*
	 type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
	/*
	 A list of a available insurance levels for this ship type */
	Levels []GetInsurancePricesLevel `json:"levels,omitempty"`
}