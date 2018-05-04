package esiv1



/* 
200 ok object */
type GetMarketsGroupsMarketGroupIdOk struct {
/*
	 description string */
	Description string `json:"description,omitempty"`
/*
	 market_group_id integer */
	MarketGroupId int32 `json:"market_group_id,omitempty"`
/*
	 name string */
	Name string `json:"name,omitempty"`
/*
	 parent_group_id integer */
	ParentGroupId int32 `json:"parent_group_id,omitempty"`
/*
	 types array */
	Types []int32 `json:"types,omitempty"`
}
