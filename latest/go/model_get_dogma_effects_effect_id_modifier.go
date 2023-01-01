package esilatest

/*
modifier object
*/
type GetDogmaEffectsEffectIdModifier struct {
	/*
	 domain string */
	Domain string `json:"domain,omitempty"`
	/*
	 effect_id integer */
	EffectId int32 `json:"effect_id,omitempty"`
	/*
	 func string */
	Func_ string `json:"func,omitempty"`
	/*
	 modified_attribute_id integer */
	ModifiedAttributeId int32 `json:"modified_attribute_id,omitempty"`
	/*
	 modifying_attribute_id integer */
	ModifyingAttributeId int32 `json:"modifying_attribute_id,omitempty"`
	/*
	 operator integer */
	Operator int32 `json:"operator,omitempty"`
}
