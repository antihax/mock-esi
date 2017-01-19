package swServer

import "time"
var _ time.Time

/* 
fitting object */
type PostCharactersCharacterIdFittingsFitting struct {
/*
	 description string */
	description string `json:"description,omitempty"`
/*
	 items array */
	items []PostCharactersCharacterIdFittingsItem `json:"items,omitempty"`
/*
	 name string */
	name string `json:"name,omitempty"`
/*
	 ship_type_id integer */
	ship_type_id int32 `json:"ship_type_id,omitempty"`
}
