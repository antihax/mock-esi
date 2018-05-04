package esilegacy

/*
contents object */
type PutCharactersCharacterIdMailMailIdContents struct {
	/*
	 Labels to assign to the mail. Pre-existing labels are unassigned. */
	Labels []int32 `json:"labels,omitempty"`
	/*
	 Whether the mail is flagged as read */
	Read bool `json:"read,omitempty"`
}
