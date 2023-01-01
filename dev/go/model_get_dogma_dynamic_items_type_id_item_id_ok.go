package esidev

/*
200 ok object
*/
type GetDogmaDynamicItemsTypeIdItemIdOk struct {
	/*
	 The ID of the character who created the item */
	CreatedBy int32 `json:"created_by,omitempty"`
	/*
	 dogma_attributes array */
	DogmaAttributes []GetDogmaDynamicItemsTypeIdItemIdDogmaAttribute `json:"dogma_attributes,omitempty"`
	/*
	 dogma_effects array */
	DogmaEffects []GetDogmaDynamicItemsTypeIdItemIdDogmaEffect `json:"dogma_effects,omitempty"`
	/*
	 The type ID of the mutator used to generate the dynamic item. */
	MutatorTypeId int32 `json:"mutator_type_id,omitempty"`
	/*
	 The type ID of the source item the mutator was applied to create the dynamic item. */
	SourceTypeId int32 `json:"source_type_id,omitempty"`
}
