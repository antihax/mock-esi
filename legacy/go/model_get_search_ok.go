package esilegacy

/*
200 ok object */
type GetSearchOk struct {
	/*
	 agent array */
	Agent []int32 `json:"agent,omitempty"`
	/*
	 alliance array */
	Alliance []int32 `json:"alliance,omitempty"`
	/*
	 character array */
	Character []int32 `json:"character,omitempty"`
	/*
	 constellation array */
	Constellation []int32 `json:"constellation,omitempty"`
	/*
	 corporation array */
	Corporation []int32 `json:"corporation,omitempty"`
	/*
	 faction array */
	Faction []int32 `json:"faction,omitempty"`
	/*
	 inventory_type array */
	InventoryType []int32 `json:"inventory_type,omitempty"`
	/*
	 region array */
	Region []int32 `json:"region,omitempty"`
	/*
	 solar_system array */
	SolarSystem []int32 `json:"solar_system,omitempty"`
	/*
	 station array */
	Station []int32 `json:"station,omitempty"`
}
