package esidev

/*
level object */
type GetInsurancePricesLevel struct {
	/*
	 cost number */
	Cost float32 `json:"cost,omitempty"`
	/*
	 payout number */
	Payout float32 `json:"payout,omitempty"`
	/*
	 Localized insurance level */
	Name string `json:"name,omitempty"`
}
