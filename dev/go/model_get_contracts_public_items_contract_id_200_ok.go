package esidev

/*
200 ok object */
type GetContractsPublicItemsContractId200Ok struct {
	/*
	 is_blueprint_copy boolean */
	IsBlueprintCopy bool `json:"is_blueprint_copy,omitempty"`
	/*
	 true if the contract issuer has submitted this item with the contract, false if the isser is asking for this item in the contract */
	IsIncluded bool `json:"is_included,omitempty"`
	/*
	 Unique ID for the item being sold. Not present if item is being requested by contract rather than sold with contract */
	ItemId int64 `json:"item_id,omitempty"`
	/*
	 Material Efficiency Level of the blueprint */
	MaterialEfficiency int32 `json:"material_efficiency,omitempty"`
	/*
	 Number of items in the stack */
	Quantity int32 `json:"quantity,omitempty"`
	/*
	 Unique ID for the item, used by the contract system */
	RecordId int64 `json:"record_id,omitempty"`
	/*
	 Number of runs remaining if the blueprint is a copy, -1 if it is an original */
	Runs int32 `json:"runs,omitempty"`
	/*
	 Time Efficiency Level of the blueprint */
	TimeEfficiency int32 `json:"time_efficiency,omitempty"`
	/*
	 Type ID for item */
	TypeId int32 `json:"type_id,omitempty"`
}
