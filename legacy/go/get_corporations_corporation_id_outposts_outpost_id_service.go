package esilegacy

/*
service object */
type GetCorporationsCorporationIdOutpostsOutpostIdService struct {
	/*
	 service_name string */
	ServiceName string `json:"service_name,omitempty"`
	/*
	 minimum_standing number */
	MinimumStanding float64 `json:"minimum_standing,omitempty"`
	/*
	 surcharge_per_bad_standing number */
	SurchargePerBadStanding float64 `json:"surcharge_per_bad_standing,omitempty"`
	/*
	 discount_per_good_standing number */
	DiscountPerGoodStanding float64 `json:"discount_per_good_standing,omitempty"`
}
