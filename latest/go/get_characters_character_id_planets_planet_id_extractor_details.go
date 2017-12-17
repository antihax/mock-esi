package esilatest

/*
extractor_details object */
type GetCharactersCharacterIdPlanetsPlanetIdExtractorDetails struct {
	/*
	 heads array */
	Heads []GetCharactersCharacterIdPlanetsPlanetIdHead `json:"heads,omitempty"`
	/*
	 product_type_id integer */
	ProductTypeId int32 `json:"product_type_id,omitempty"`
	/*
	 in seconds */
	CycleTime int32 `json:"cycle_time,omitempty"`
	/*
	 head_radius number */
	HeadRadius float32 `json:"head_radius,omitempty"`
	/*
	 qty_per_cycle integer */
	QtyPerCycle int32 `json:"qty_per_cycle,omitempty"`
}
