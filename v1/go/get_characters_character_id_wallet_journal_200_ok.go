package esiv1

import "time"

/*
200 ok object */
type GetCharactersCharacterIdWalletJournal200Ok struct {
	/*
	 Date and time of transaction */
	Date time.Time `json:"date,omitempty"`
	/*
	 Unique journal reference ID */
	RefId int64 `json:"ref_id,omitempty"`
	/*
	 Transaction type, different type of transaction will populate different fields in `extra_info` */
	RefType string `json:"ref_type,omitempty"`
	/*
	 first_party_id integer */
	FirstPartyId int32 `json:"first_party_id,omitempty"`
	/*
	 first_party_type string */
	FirstPartyType string `json:"first_party_type,omitempty"`
	/*
	 second_party_id integer */
	SecondPartyId int32 `json:"second_party_id,omitempty"`
	/*
	 second_party_type string */
	SecondPartyType string `json:"second_party_type,omitempty"`
	/*
	 Transaction amount. Positive when value transferred to the first party. Negative otherwise */
	Amount float64 `json:"amount,omitempty"`
	/*
	 Wallet balance after transaction occurred */
	Balance float64 `json:"balance,omitempty"`
	/*
	 reason string */
	Reason string `json:"reason,omitempty"`
	/*
	 the corporation ID receiving any tax paid */
	TaxRecieverId int32 `json:"tax_reciever_id,omitempty"`
	/*
	 Tax amount received for tax related transactions */
	Tax float64 `json:"tax,omitempty"`
	/* */
	ExtraInfo GetCharactersCharacterIdWalletJournalExtraInfo `json:"extra_info,omitempty"`
}
