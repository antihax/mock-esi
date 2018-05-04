package esilegacy



/* 
200 ok object */
type GetIndustrySystems200Ok struct {
/*
	 cost_indices array */
	CostIndices []GetIndustrySystemsCostIndice `json:"cost_indices,omitempty"`
/*
	 solar_system_id integer */
	SolarSystemId int32 `json:"solar_system_id,omitempty"`
}
