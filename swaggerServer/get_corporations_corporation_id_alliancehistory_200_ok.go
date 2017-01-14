package swaggerServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetCorporationsCorporationIdAlliancehistory200Ok struct {
/* */
	alliance GetCorporationsCorporationIdAlliancehistoryAlliance `json:"alliance,omitempty"`
/*
	 An incrementing ID that can be used to canonically establish order of records in cases where dates may be ambiguous */
	record_id int32 `json:"record_id,omitempty"`
/*
	 start_date string */
	start_date time.Time `json:"start_date,omitempty"`
}
