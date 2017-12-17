package esidev

/*
market object */
type GetCharactersCharacterIdStatsMarket struct {
	/*
	 accept_contracts_courier integer */
	AcceptContractsCourier int64 `json:"accept_contracts_courier,omitempty"`
	/*
	 accept_contracts_item_exchange integer */
	AcceptContractsItemExchange int64 `json:"accept_contracts_item_exchange,omitempty"`
	/*
	 buy_orders_placed integer */
	BuyOrdersPlaced int64 `json:"buy_orders_placed,omitempty"`
	/*
	 cancel_market_order integer */
	CancelMarketOrder int64 `json:"cancel_market_order,omitempty"`
	/*
	 create_contracts_auction integer */
	CreateContractsAuction int64 `json:"create_contracts_auction,omitempty"`
	/*
	 create_contracts_courier integer */
	CreateContractsCourier int64 `json:"create_contracts_courier,omitempty"`
	/*
	 create_contracts_item_exchange integer */
	CreateContractsItemExchange int64 `json:"create_contracts_item_exchange,omitempty"`
	/*
	 deliver_courier_contract integer */
	DeliverCourierContract int64 `json:"deliver_courier_contract,omitempty"`
	/*
	 isk_gained integer */
	IskGained int64 `json:"isk_gained,omitempty"`
	/*
	 isk_spent integer */
	IskSpent int64 `json:"isk_spent,omitempty"`
	/*
	 modify_market_order integer */
	ModifyMarketOrder int64 `json:"modify_market_order,omitempty"`
	/*
	 search_contracts integer */
	SearchContracts int64 `json:"search_contracts,omitempty"`
	/*
	 sell_orders_placed integer */
	SellOrdersPlaced int64 `json:"sell_orders_placed,omitempty"`
}
