package esilatest

/*
200 ok object */
type GetMarketsGroupsMarketGroupIdOk struct {
	/*
	 market_group_id integer */
	MarketGroupId int32 `json:"market_group_id,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 description string */
	Description string `json:"description,omitempty"`
	/*
	 types array */
	Types []int32 `json:"types,omitempty"`
	/*
	 parent_group_id integer */
	ParentGroupId int32 `json:"parent_group_id,omitempty"`
}
