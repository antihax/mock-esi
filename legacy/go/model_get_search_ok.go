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
	 inventorytype array */
	Inventorytype []int32 `json:"inventorytype,omitempty"`
	/*
	 region array */
	Region []int32 `json:"region,omitempty"`
	/*
	 solarsystem array */
	Solarsystem []int32 `json:"solarsystem,omitempty"`
	/*
	 station array */
	Station []int32 `json:"station,omitempty"`
	/*
	 wormhole array */
	Wormhole []int32 `json:"wormhole,omitempty"`
}
