package esilatest

/*
200 ok object
*/
type GetAlliancesAllianceIdContacts200Ok struct {
	/*
	 contact_id integer */
	ContactId int32 `json:"contact_id,omitempty"`
	/*
	 contact_type string */
	ContactType string `json:"contact_type,omitempty"`
	/*
	 label_ids array */
	LabelIds []int64 `json:"label_ids,omitempty"`
	/*
	 Standing of the contact */
	Standing float32 `json:"standing,omitempty"`
}
