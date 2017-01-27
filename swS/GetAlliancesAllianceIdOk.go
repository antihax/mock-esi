package swS

import "time"
var _ time.Time

/* 
200 ok object */
type GetAlliancesAllianceIdOk struct {
/*
	 the full name of the alliance */
	alliance_name string `json:"alliance_name,omitempty"`
/*
	 date_founded string */
	date_founded time.Time `json:"date_founded,omitempty"`
/*
	 the executor corporation ID, if this alliance is not closed */
	executor_corp int32 `json:"executor_corp,omitempty"`
/*
	 the short name of the alliance */
	ticker string `json:"ticker,omitempty"`
}
