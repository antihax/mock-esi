package swServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetUniverseTypesTypeIdOk struct {
/*
	 capacity number */
	capacity float32 `json:"capacity,omitempty"`
/*
	 description string */
	description string `json:"description,omitempty"`
/*
	 dogma_attributes array */
	dogma_attributes []GetUniverseTypesTypeIdDogmaAttribute `json:"dogma_attributes,omitempty"`
/*
	 dogma_effects array */
	dogma_effects []GetUniverseTypesTypeIdDogmaEffect `json:"dogma_effects,omitempty"`
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
	 mass number */
	mass float32 `json:"mass,omitempty"`
/*
	 name string */
	name string `json:"name,omitempty"`
/*
	 portion_size integer */
	portion_size int32 `json:"portion_size,omitempty"`
/*
	 published boolean */
	published bool `json:"published,omitempty"`
/*
	 radius number */
	radius float32 `json:"radius,omitempty"`
/*
	 type_id integer */
	type_id int32 `json:"type_id,omitempty"`
/*
	 volume number */
	volume float32 `json:"volume,omitempty"`
}
