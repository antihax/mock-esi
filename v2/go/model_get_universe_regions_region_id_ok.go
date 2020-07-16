package esiv2

/*
200 ok object */
type GetUniverseRegionsRegionIdOk struct {
	/*
	 constellations array */
	Constellations []int32 `json:"constellations,omitempty"`
	/*
	 description string */
	Description string `json:"description,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 region_id integer */
	RegionId int32 `json:"region_id,omitempty"`
}
