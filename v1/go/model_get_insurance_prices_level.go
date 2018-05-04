package esiv1



/* 
level object */
type GetInsurancePricesLevel struct {
/*
	 cost number */
	Cost float32 `json:"cost,omitempty"`
/*
	 Localized insurance level */
	Name string `json:"name,omitempty"`
/*
	 payout number */
	Payout float32 `json:"payout,omitempty"`
}
