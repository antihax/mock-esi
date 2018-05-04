package esilatest



/* 
200 ok object */
type GetUniverseCategoriesCategoryIdOk struct {
/*
	 category_id integer */
	CategoryId int32 `json:"category_id,omitempty"`
/*
	 groups array */
	Groups []int32 `json:"groups,omitempty"`
/*
	 name string */
	Name string `json:"name,omitempty"`
/*
	 published boolean */
	Published bool `json:"published,omitempty"`
}
