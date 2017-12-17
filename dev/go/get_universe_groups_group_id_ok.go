package esidev

/*
200 ok object */
type GetUniverseGroupsGroupIdOk struct {
	/*
	 group_id integer */
	GroupId int32 `json:"group_id,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 published boolean */
	Published bool `json:"published,omitempty"`
	/*
	 category_id integer */
	CategoryId int32 `json:"category_id,omitempty"`
	/*
	 types array */
	Types []int32 `json:"types,omitempty"`
}
