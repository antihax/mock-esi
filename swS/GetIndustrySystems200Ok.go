package swS

import "time"
var _ time.Time

/* 
200 ok object */
type GetIndustrySystems200Ok struct {
/*
	 cost_indices array */
	cost_indices []GetIndustrySystemsCostIndice `json:"cost_indices,omitempty"`
/*
	 solar_system_id integer */
	solar_system_id int32 `json:"solar_system_id,omitempty"`
}
