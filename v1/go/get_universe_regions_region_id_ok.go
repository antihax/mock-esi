package esiv1

/*
200 ok object */
type GetUniverseRegionsRegionIdOk struct {
	/*
	 region_id integer */
	RegionId int32 `json:"region_id,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 description string */
	Description string `json:"description,omitempty"`
	/*
	 constellations array */
	Constellations []int32 `json:"constellations,omitempty"`
}
