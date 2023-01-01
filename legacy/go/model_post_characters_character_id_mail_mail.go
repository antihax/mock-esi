package esilegacy

/*
mail object
*/
type PostCharactersCharacterIdMailMail struct {
	/*
	 approved_cost integer */
	ApprovedCost int64 `json:"approved_cost,omitempty"`
	/*
	 body string */
	Body string `json:"body,omitempty"`
	/*
	 recipients array */
	Recipients []PostCharactersCharacterIdMailRecipient `json:"recipients,omitempty"`
	/*
	 subject string */
	Subject string `json:"subject,omitempty"`
}
