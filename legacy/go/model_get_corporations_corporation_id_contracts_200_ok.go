package esilegacy



/* 
200 ok object */
type GetCorporationsCorporationIdContracts200Ok struct {
/*
	 Who will accept the contract */
	AcceptorId int32 `json:"acceptor_id,omitempty"`
/*
	 ID to whom the contract is assigned, can be corporation or character ID */
	AssigneeId int32 `json:"assignee_id,omitempty"`
/*
	 To whom the contract is available */
	Availability string `json:"availability,omitempty"`
/*
	 Buyout price (for Auctions only) */
	Buyout float64 `json:"buyout,omitempty"`
/*
	 Collateral price (for Couriers only) */
	Collateral float64 `json:"collateral,omitempty"`
/*
	 contract_id integer */
	ContractId int32 `json:"contract_id,omitempty"`
/*
	 Date of confirmation of contract */
	DateAccepted time.Time `json:"date_accepted,omitempty"`
/*
	 Date of completed of contract */
	DateCompleted time.Time `json:"date_completed,omitempty"`
/*
	 Expiration date of the contract */
	DateExpired time.Time `json:"date_expired,omitempty"`
/*
	 Сreation date of the contract */
	DateIssued time.Time `json:"date_issued,omitempty"`
/*
	 Number of days to perform the contract */
	DaysToComplete int32 `json:"days_to_complete,omitempty"`
/*
	 End location ID (for Couriers contract) */
	EndLocationId int64 `json:"end_location_id,omitempty"`
/*
	 true if the contract was issued on behalf of the issuer's corporation */
	ForCorporation bool `json:"for_corporation,omitempty"`
/*
	 Character's corporation ID for the issuer */
	IssuerCorporationId int32 `json:"issuer_corporation_id,omitempty"`
/*
	 Character ID for the issuer */
	IssuerId int32 `json:"issuer_id,omitempty"`
/*
	 Price of contract (for ItemsExchange and Auctions) */
	Price float64 `json:"price,omitempty"`
/*
	 Remuneration for contract (for Couriers only) */
	Reward float64 `json:"reward,omitempty"`
/*
	 Start location ID (for Couriers contract) */
	StartLocationId int64 `json:"start_location_id,omitempty"`
/*
	 Status of the the contract */
	Status string `json:"status,omitempty"`
/*
	 Title of the contract */
	Title string `json:"title,omitempty"`
/*
	 Type of the contract */
	Type_ string `json:"type,omitempty"`
/*
	 Volume of items in the contract */
	Volume float64 `json:"volume,omitempty"`
}
