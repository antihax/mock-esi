package esilatest

/*
item object */
type PostCharactersCharacterIdFittingsItem struct {
	/*
	 Fitting location for the item. Entries placed in 'Invalid' will be discarded. If this leaves the fitting with nothing, it will cause an error. */
	Flag string `json:"flag,omitempty"`
	/*
	 quantity integer */
	Quantity int32 `json:"quantity,omitempty"`
	/*
	 type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
}
