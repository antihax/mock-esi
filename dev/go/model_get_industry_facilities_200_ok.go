package esidev

/*
200 ok object
*/
type GetIndustryFacilities200Ok struct {
	/*
	 ID of the facility */
	FacilityId int64 `json:"facility_id,omitempty"`
	/*
	 Owner of the facility */
	OwnerId int32 `json:"owner_id,omitempty"`
	/*
	 Region ID where the facility is */
	RegionId int32 `json:"region_id,omitempty"`
	/*
	 Solar system ID where the facility is */
	SolarSystemId int32 `json:"solar_system_id,omitempty"`
	/*
	 Tax imposed by the facility */
	Tax float32 `json:"tax,omitempty"`
	/*
	 Type ID of the facility */
	TypeId int32 `json:"type_id,omitempty"`
}
