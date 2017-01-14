package swaggerServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetUniverseTypesTypeIdOk struct {
/*
	 category_id integer */
	category_id int32 `json:"category_id,omitempty"`
/*
	 graphic_id integer */
	graphic_id int32 `json:"graphic_id,omitempty"`
/*
	 group_id integer */
	group_id int32 `json:"group_id,omitempty"`
/*
	 icon_id integer */
	icon_id int32 `json:"icon_id,omitempty"`
/*
	 type_description string */
	type_description string `json:"type_description,omitempty"`
/*
	 type_name string */
	type_name string `json:"type_name,omitempty"`
}
