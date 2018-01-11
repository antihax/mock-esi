package esilatest

/*
200 ok object */
type GetUniverseTypesTypeIdOk struct {
	/*
	 type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 description string */
	Description string `json:"description,omitempty"`
	/*
	 published boolean */
	Published bool `json:"published,omitempty"`
	/*
	 group_id integer */
	GroupId int32 `json:"group_id,omitempty"`
	/*
	 This only exists for types that can be put on the market */
	MarketGroupId int32 `json:"market_group_id,omitempty"`
	/*
	 radius number */
	Radius float32 `json:"radius,omitempty"`
	/*
	 volume number */
	Volume float32 `json:"volume,omitempty"`
	/*
	 packaged_volume number */
	PackagedVolume float32 `json:"packaged_volume,omitempty"`
	/*
	 icon_id integer */
	IconId int32 `json:"icon_id,omitempty"`
	/*
	 capacity number */
	Capacity float32 `json:"capacity,omitempty"`
	/*
	 portion_size integer */
	PortionSize int32 `json:"portion_size,omitempty"`
	/*
	 mass number */
	Mass float32 `json:"mass,omitempty"`
	/*
	 graphic_id integer */
	GraphicId int32 `json:"graphic_id,omitempty"`
	/*
	 dogma_attributes array */
	DogmaAttributes []GetUniverseTypesTypeIdDogmaAttribute `json:"dogma_attributes,omitempty"`
	/*
	 dogma_effects array */
	DogmaEffects []GetUniverseTypesTypeIdDogmaEffect `json:"dogma_effects,omitempty"`
}
