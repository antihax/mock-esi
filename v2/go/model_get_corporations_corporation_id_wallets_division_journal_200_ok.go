package esiv2



/* 
200 ok object */
type GetCorporationsCorporationIdWalletsDivisionJournal200Ok struct {
/*
	 Transaction amount. Positive when value transferred to the first party. Negative otherwise */
	Amount float64 `json:"amount,omitempty"`
/*
	 Wallet balance after transaction occurred */
	Balance float64 `json:"balance,omitempty"`
/*
	 Date and time of transaction */
	Date time.Time `json:"date,omitempty"`
/* */
	ExtraInfo GetCorporationsCorporationIdWalletsDivisionJournalExtraInfo `json:"extra_info,omitempty"`
/*
	 first_party_id integer */
	FirstPartyId int32 `json:"first_party_id,omitempty"`
/*
	 first_party_type string */
	FirstPartyType string `json:"first_party_type,omitempty"`
/*
	 reason string */
	Reason string `json:"reason,omitempty"`
/*
	 Unique journal reference ID */
	RefId int64 `json:"ref_id,omitempty"`
/*
	 Transaction type, different type of transaction will populate different fields in `extra_info` Note: If you have an existing XML API application that is using ref_types, you will need to know which string ESI ref_type maps to which integer. You can use the following gist to see string->int mappings: https://gist.github.com/ccp-zoetrope/c03db66d90c2148724c06171bc52e0ec */
	RefType string `json:"ref_type,omitempty"`
/*
	 second_party_id integer */
	SecondPartyId int32 `json:"second_party_id,omitempty"`
/*
	 second_party_type string */
	SecondPartyType string `json:"second_party_type,omitempty"`
/*
	 Tax amount received for tax related transactions */
	Tax float64 `json:"tax,omitempty"`
/*
	 the corporation ID receiving any tax paid */
	TaxReceiverId int32 `json:"tax_receiver_id,omitempty"`
}
