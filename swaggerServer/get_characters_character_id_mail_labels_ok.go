package swaggerServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersCharacterIdMailLabelsOk struct {
/*
	 labels array */
	labels []GetCharactersCharacterIdMailLabelsLabel `json:"labels,omitempty"`
/*
	 total_unread_count integer */
	total_unread_count int32 `json:"total_unread_count,omitempty"`
}
