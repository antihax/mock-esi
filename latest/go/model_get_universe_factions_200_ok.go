package esilatest

/*
200 ok object */
type GetUniverseFactions200Ok struct {
	/*
	 corporation_id integer */
	CorporationId int32 `json:"corporation_id,omitempty"`
	/*
	 description string */
	Description string `json:"description,omitempty"`
	/*
	 faction_id integer */
	FactionId int32 `json:"faction_id,omitempty"`
	/*
	 is_unique boolean */
	IsUnique bool `json:"is_unique,omitempty"`
	/*
	 militia_corporation_id integer */
	MilitiaCorporationId int32 `json:"militia_corporation_id,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 size_factor number */
	SizeFactor float32 `json:"size_factor,omitempty"`
	/*
	 solar_system_id integer */
	SolarSystemId int32 `json:"solar_system_id,omitempty"`
	/*
	 station_count integer */
	StationCount int32 `json:"station_count,omitempty"`
	/*
	 station_system_count integer */
	StationSystemCount int32 `json:"station_system_count,omitempty"`
}
