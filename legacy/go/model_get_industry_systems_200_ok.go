package esilegacy

/*
200 ok object */
type GetIndustrySystems200Ok struct {
	/*
	 solar_system_id integer */
	SolarSystemId int32 `json:"solar_system_id,omitempty"`
	/*
	 cost_indices array */
	CostIndices []GetIndustrySystemsCostIndice `json:"cost_indices,omitempty"`
}
