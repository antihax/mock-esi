package esilatest

/*
fitting object */
type PostCharactersCharacterIdFittingsFitting struct {
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
	Items []PostCharactersCharacterIdFittingsItem `json:"items,omitempty"`
}
