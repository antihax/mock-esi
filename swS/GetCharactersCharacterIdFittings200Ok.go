package swS

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersCharacterIdFittings200Ok struct {
/*
	 description string */
	description string `json:"description,omitempty"`
/*
	 fitting_id integer */
	fitting_id int32 `json:"fitting_id,omitempty"`
/*
	 items array */
	items []GetCharactersCharacterIdFittingsItem `json:"items,omitempty"`
/*
	 name string */
	name string `json:"name,omitempty"`
/*
	 ship_type_id integer */
	ship_type_id int32 `json:"ship_type_id,omitempty"`
}
