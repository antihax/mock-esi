package swS

import "time"
var _ time.Time

/* 
mail schema */
type PostCharactersCharacterIdMailMail struct {
/*
	 approved_cost integer */
	approved_cost int64 `json:"approved_cost,omitempty"`
/*
	 body string */
	body string `json:"body,omitempty"`
/*
	 recipients array */
	recipients []PostCharactersCharacterIdMailRecipient `json:"recipients,omitempty"`
/*
	 subject string */
	subject string `json:"subject,omitempty"`
}
