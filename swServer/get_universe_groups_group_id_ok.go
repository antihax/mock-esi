package swServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetUniverseGroupsGroupIdOk struct {
/*
	 category_id number */
	category_id float32 `json:"category_id,omitempty"`
/*
	 group_id integer */
	group_id int32 `json:"group_id,omitempty"`
/*
	 name string */
	name string `json:"name,omitempty"`
/*
	 published boolean */
	published bool `json:"published,omitempty"`
/*
	 types array */
	types []int32 `json:"types,omitempty"`
}
