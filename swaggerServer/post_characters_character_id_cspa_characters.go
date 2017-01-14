package swaggerServer

import "time"
var _ time.Time

/* 
characters schema */
type PostCharactersCharacterIdCspaCharacters struct {
/*
	 characters array */
	characters []int32 `json:"characters,omitempty"`
}
