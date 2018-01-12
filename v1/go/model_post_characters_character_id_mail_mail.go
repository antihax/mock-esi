package esiv1

/*
mail object */
type PostCharactersCharacterIdMailMail struct {
	/*
	 recipients array */
	Recipients []PostCharactersCharacterIdMailRecipient `json:"recipients,omitempty"`
	/*
	 subject string */
	Subject string `json:"subject,omitempty"`
	/*
	 body string */
	Body string `json:"body,omitempty"`
	/*
	 approved_cost integer */
	ApprovedCost int64 `json:"approved_cost,omitempty"`
}
