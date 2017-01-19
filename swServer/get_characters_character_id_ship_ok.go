package swServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersCharacterIdShipOk struct {
/*
	 Item id's are unique to a ship and persist until it is repackaged. This value can be used to track repeated uses of a ship, or detect when a pilot changes into a different instance of the same ship type. */
	ship_item_id int64 `json:"ship_item_id,omitempty"`
/*
	 ship_name string */
	ship_name string `json:"ship_name,omitempty"`
/*
	 ship_type_id integer */
	ship_type_id int32 `json:"ship_type_id,omitempty"`
}
