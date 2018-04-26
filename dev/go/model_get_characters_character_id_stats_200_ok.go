package esidev

/*
Aggregate stats for a year */
type GetCharactersCharacterIdStats200Ok struct {
	/* */
	Character GetCharactersCharacterIdStatsCharacter `json:"character,omitempty"`
	/* */
	Combat GetCharactersCharacterIdStatsCombat `json:"combat,omitempty"`
	/* */
	Industry GetCharactersCharacterIdStatsIndustry `json:"industry,omitempty"`
	/* */
	Inventory GetCharactersCharacterIdStatsInventory `json:"inventory,omitempty"`
	/* */
	Isk GetCharactersCharacterIdStatsIsk `json:"isk,omitempty"`
	/* */
	Market GetCharactersCharacterIdStatsMarket `json:"market,omitempty"`
	/* */
	Mining GetCharactersCharacterIdStatsMining `json:"mining,omitempty"`
	/* */
	Module GetCharactersCharacterIdStatsModule `json:"module,omitempty"`
	/* */
	Orbital GetCharactersCharacterIdStatsOrbital `json:"orbital,omitempty"`
	/* */
	Pve GetCharactersCharacterIdStatsPve `json:"pve,omitempty"`
	/* */
	Social GetCharactersCharacterIdStatsSocial `json:"social,omitempty"`
	/* */
	Travel GetCharactersCharacterIdStatsTravel `json:"travel,omitempty"`
	/*
	 Gregorian year for this set of aggregates */
	Year int32 `json:"year,omitempty"`
}
