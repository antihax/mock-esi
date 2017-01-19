package swServer

import "time"
var _ time.Time

/* 
item object */
type GetCharactersCharacterIdFittingsItem struct {
/*
	 flag integer */
	flag int32 `json:"flag,omitempty"`
/*
	 quantity integer */
	quantity int32 `json:"quantity,omitempty"`
/*
	 type_id integer */
	type_id int32 `json:"type_id,omitempty"`
}
