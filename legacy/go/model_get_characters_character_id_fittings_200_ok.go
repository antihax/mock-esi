package esilegacy

/*
200 ok object */
type GetCharactersCharacterIdFittings200Ok struct {
	/*
	 fitting_id integer */
	FittingId int32 `json:"fitting_id,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 description string */
	Description string `json:"description,omitempty"`
	/*
	 ship_type_id integer */
	ShipTypeId int32 `json:"ship_type_id,omitempty"`
	/*
	 items array */
	Items []GetCharactersCharacterIdFittingsItem `json:"items,omitempty"`
}
