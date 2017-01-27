package swS

import "time"
var _ time.Time

/* 
level object */
type GetInsurancePricesLevel struct {
/*
	 cost number */
	cost float32 `json:"cost,omitempty"`
/*
	 Localized insurance level */
	name string `json:"name,omitempty"`
/*
	 payout number */
	payout float32 `json:"payout,omitempty"`
}
