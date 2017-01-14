package swaggerServer

import "time"
var _ time.Time

/* 
contents object */
type PutCharactersCharacterIdMailMailIdContents struct {
/*
	 Labels to assign to the mail. Pre-existing labels are unassigned. */
	labels []int64 `json:"labels,omitempty"`
/*
	 Whether the mail is flagged as read */
	read bool `json:"read,omitempty"`
}
