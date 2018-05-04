package esiv2

/*
200 ok object */
type GetUniverseTypesTypeIdOk struct {
	/*
	 capacity number */
	Capacity float32 `json:"capacity,omitempty"`
	/*
	 description string */
	Description string `json:"description,omitempty"`
	/*
	 dogma_attributes array */
	DogmaAttributes []GetUniverseTypesTypeIdDogmaAttribute `json:"dogma_attributes,omitempty"`
	/*
	 dogma_effects array */
	DogmaEffects []GetUniverseTypesTypeIdDogmaEffect `json:"dogma_effects,omitempty"`
	/*
	 graphic_id integer */
	GraphicId int32 `json:"graphic_id,omitempty"`
	/*
	 group_id integer */
	GroupId int32 `json:"group_id,omitempty"`
	/*
	 icon_id integer */
	IconId int32 `json:"icon_id,omitempty"`
	/*
	 mass number */
	Mass float32 `json:"mass,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 portion_size integer */
	PortionSize int32 `json:"portion_size,omitempty"`
	/*
	 published boolean */
	Published bool `json:"published,omitempty"`
	/*
	 radius number */
	Radius float32 `json:"radius,omitempty"`
	/*
	 type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
	/*
	 volume number */
	Volume float32 `json:"volume,omitempty"`
}
