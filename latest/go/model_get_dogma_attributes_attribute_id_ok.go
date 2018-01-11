package esilatest

/*
200 ok object */
type GetDogmaAttributesAttributeIdOk struct {
	/*
	 attribute_id integer */
	AttributeId int32 `json:"attribute_id,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 description string */
	Description string `json:"description,omitempty"`
	/*
	 icon_id integer */
	IconId int32 `json:"icon_id,omitempty"`
	/*
	 default_value number */
	DefaultValue float32 `json:"default_value,omitempty"`
	/*
	 published boolean */
	Published bool `json:"published,omitempty"`
	/*
	 display_name string */
	DisplayName string `json:"display_name,omitempty"`
	/*
	 unit_id integer */
	UnitId int32 `json:"unit_id,omitempty"`
	/*
	 stackable boolean */
	Stackable bool `json:"stackable,omitempty"`
	/*
	 high_is_good boolean */
	HighIsGood bool `json:"high_is_good,omitempty"`
}
