package swServer

import "time"
var _ time.Time

/* 
label object */
type PostCharactersCharacterIdMailLabelsLabel struct {
/*
	 Hexadecimal string representing label color,
in RGB format
 */
	color string `json:"color,omitempty"`
/*
	 name string */
	name string `json:"name,omitempty"`
}
