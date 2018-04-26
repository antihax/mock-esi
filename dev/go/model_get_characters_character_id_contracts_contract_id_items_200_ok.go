package esidev

/*
200 ok object */
type GetCharactersCharacterIdContractsContractIdItems200Ok struct {
	/*
	 true if the contract issuer has submitted this item with the contract, false if the isser is asking for this item in the contract. */
	IsIncluded bool `json:"is_included,omitempty"`
	/*
	 is_singleton boolean */
	IsSingleton bool `json:"is_singleton,omitempty"`
	/*
	 Number of items in the stack */
	Quantity int32 `json:"quantity,omitempty"`
	/*
	 -1 indicates that the item is a singleton (non-stackable). If the item happens to be a Blueprint, -1 is an Original and -2 is a Blueprint Copy */
	RawQuantity int32 `json:"raw_quantity,omitempty"`
	/*
	 Unique ID for the item */
	RecordId int64 `json:"record_id,omitempty"`
	/*
	 Type ID for item */
	TypeId int32 `json:"type_id,omitempty"`
}
