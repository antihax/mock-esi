package esilegacy

/*
modifier object
*/
type GetDogmaEffectsEffectIdModifier struct {
	/*
	 domain string */
	Domain string `json:"domain,omitempty"`
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
