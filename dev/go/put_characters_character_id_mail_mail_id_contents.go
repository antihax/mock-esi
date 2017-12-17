package esidev

/*
contents object */
type PutCharactersCharacterIdMailMailIdContents struct {
	/*
	 Whether the mail is flagged as read */
	Read bool `json:"read,omitempty"`
	/*
	 Labels to assign to the mail. Pre-existing labels are unassigned. */
	Labels []int64 `json:"labels,omitempty"`
}
