package esiv1



/* 
200 ok object */
type PostUniverseIdsOk struct {
/*
	 agents array */
	Agents []PostUniverseIdsAgent `json:"agents,omitempty"`
/*
	 alliances array */
	Alliances []PostUniverseIdsAlliance `json:"alliances,omitempty"`
/*
	 characters array */
	Characters []PostUniverseIdsCharacter `json:"characters,omitempty"`
/*
	 constellations array */
	Constellations []PostUniverseIdsConstellation `json:"constellations,omitempty"`
/*
	 corporations array */
	Corporations []PostUniverseIdsCorporation `json:"corporations,omitempty"`
/*
	 factions array */
	Factions []PostUniverseIdsFaction `json:"factions,omitempty"`
/*
	 inventory_types array */
	InventoryTypes []PostUniverseIdsInventoryType `json:"inventory_types,omitempty"`
/*
	 regions array */
	Regions []PostUniverseIdsRegion `json:"regions,omitempty"`
/*
	 stations array */
	Stations []PostUniverseIdsStation `json:"stations,omitempty"`
/*
	 systems array */
	Systems []PostUniverseIdsSystem `json:"systems,omitempty"`
}
