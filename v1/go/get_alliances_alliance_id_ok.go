package esiv1

import "time"

/*
200 ok object */
type GetAlliancesAllianceIdOk struct {
	/*
	 the full name of the alliance */
	AllianceName string `json:"alliance_name,omitempty"`
	/*
	 the short name of the alliance */
	Ticker string `json:"ticker,omitempty"`
	/*
	 the executor corporation ID, if this alliance is not closed */
	ExecutorCorp int32 `json:"executor_corp,omitempty"`
	/*
	 date_founded string */
	DateFounded time.Time `json:"date_founded,omitempty"`
	/* */
	LogoUrls GetAlliancesAllianceIdLogoUrls `json:"logo_urls,omitempty"`
}
