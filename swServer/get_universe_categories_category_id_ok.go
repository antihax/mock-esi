package swServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetUniverseCategoriesCategoryIdOk struct {
/*
	 category_id integer */
	category_id int32 `json:"category_id,omitempty"`
/*
	 groups array */
	groups []int32 `json:"groups,omitempty"`
/*
	 name string */
	name string `json:"name,omitempty"`
/*
	 published boolean */
	published bool `json:"published,omitempty"`
}
