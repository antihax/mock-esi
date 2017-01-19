package swServer

import "time"
var _ time.Time

/* 
dogma_attribute object */
type GetUniverseTypesTypeIdDogmaAttribute struct {
/*
	 attribute_id integer */
	attribute_id int32 `json:"attribute_id,omitempty"`
/*
	 value number */
	value float32 `json:"value,omitempty"`
}
