package esilegacy



/* 
wallet transaction */
type GetCharactersCharacterIdWalletTransactions200Ok struct {
/*
	 client_id integer */
	ClientId int32 `json:"client_id,omitempty"`
/*
	 Date and time of transaction */
	Date time.Time `json:"date,omitempty"`
/*
	 is_buy boolean */
	IsBuy bool `json:"is_buy,omitempty"`
/*
	 is_personal boolean */
	IsPersonal bool `json:"is_personal,omitempty"`
/*
	 journal_ref_id integer */
	JournalRefId int64 `json:"journal_ref_id,omitempty"`
/*
	 location_id integer */
	LocationId int64 `json:"location_id,omitempty"`
/*
	 quantity integer */
	Quantity int32 `json:"quantity,omitempty"`
/*
	 Unique transaction ID */
	TransactionId int64 `json:"transaction_id,omitempty"`
/*
	 type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
/*
	 Amount paid per unit */
	UnitPrice float64 `json:"unit_price,omitempty"`
}
