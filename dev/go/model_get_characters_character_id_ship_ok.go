package esidev

/*
200 ok object
*/
type GetCharactersCharacterIdShipOk struct {
	/*
	 Item id's are unique to a ship and persist until it is repackaged. This value can be used to track repeated uses of a ship, or detect when a pilot changes into a different instance of the same ship type. */
	ShipItemId int64 `json:"ship_item_id,omitempty"`
	/*
	 ship_name string */
	ShipName string `json:"ship_name,omitempty"`
	/*
	 ship_type_id integer */
	ShipTypeId int32 `json:"ship_type_id,omitempty"`
}
