package esiv1

/*
200 ok object */
type GetCharactersCharacterIdFittings200Ok struct {
	/*
	 description string */
	Description string `json:"description,omitempty"`
	/*
	 fitting_id integer */
	FittingId int32 `json:"fitting_id,omitempty"`
	/*
	 items array */
	Items []GetCharactersCharacterIdFittingsItem `json:"items,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 ship_type_id integer */
	ShipTypeId int32 `json:"ship_type_id,omitempty"`
}
