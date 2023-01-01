package esiv5

import "time"

/*
200 ok object
*/
type GetCharactersCharacterIdWalletJournal200Ok struct {
	/*
	 The amount of ISK given or taken from the wallet as a result of the given transaction. Positive when ISK is deposited into the wallet and negative when ISK is withdrawn */
	Amount float64 `json:"amount,omitempty"`
	/*
	 Wallet balance after transaction occurred */
	Balance float64 `json:"balance,omitempty"`
	/*
	 An ID that gives extra context to the particular transaction. Because of legacy reasons the context is completely different per ref_type and means different things. It is also possible to not have a context_id */
	ContextId int64 `json:"context_id,omitempty"`
	/*
	 The type of the given context_id if present */
	ContextIdType string `json:"context_id_type,omitempty"`
	/*
	 Date and time of transaction */
	Date time.Time `json:"date,omitempty"`
	/*
	 The reason for the transaction, mirrors what is seen in the client */
	Description string `json:"description,omitempty"`
	/*
	 The id of the first party involved in the transaction. This attribute has no consistency and is different or non existant for particular ref_types. The description attribute will help make sense of what this attribute means. For more info about the given ID it can be dropped into the /universe/names/ ESI route to determine its type and name */
	FirstPartyId int32 `json:"first_party_id,omitempty"`
	/*
	 Unique journal reference ID */
	Id int64 `json:"id,omitempty"`
	/*
	 The user stated reason for the transaction. Only applies to some ref_types */
	Reason string `json:"reason,omitempty"`
	/*
	 \"The transaction type for the given. transaction. Different transaction types will populate different attributes.\" */
	RefType string `json:"ref_type,omitempty"`
	/*
	 The id of the second party involved in the transaction. This attribute has no consistency and is different or non existant for particular ref_types. The description attribute will help make sense of what this attribute means. For more info about the given ID it can be dropped into the /universe/names/ ESI route to determine its type and name */
	SecondPartyId int32 `json:"second_party_id,omitempty"`
	/*
	 Tax amount received. Only applies to tax related transactions */
	Tax float64 `json:"tax,omitempty"`
	/*
	 The corporation ID receiving any tax paid. Only applies to tax related transactions */
	TaxReceiverId int32 `json:"tax_receiver_id,omitempty"`
}
