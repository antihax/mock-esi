package swServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetIndustryFacilities200Ok struct {
/*
	 ID of the facility */
	facility_id int64 `json:"facility_id,omitempty"`
/*
	 Owner of the facility */
	owner_id int32 `json:"owner_id,omitempty"`
/*
	 Region ID where the facility is */
	region_id int32 `json:"region_id,omitempty"`
/*
	 Solar system ID where the facility is */
	solar_system_id int32 `json:"solar_system_id,omitempty"`
/*
	 Tax imposed by the facility */
	tax float32 `json:"tax,omitempty"`
/*
	 Type ID of the facility */
	type_id int32 `json:"type_id,omitempty"`
}
