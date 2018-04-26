package esilegacy

/*
200 ok object */
type GetDogmaAttributesAttributeIdOk struct {
	/*
	 attribute_id integer */
	AttributeId int32 `json:"attribute_id,omitempty"`
	/*
	 default_value number */
	DefaultValue float32 `json:"default_value,omitempty"`
	/*
	 description string */
	Description string `json:"description,omitempty"`
	/*
	 display_name string */
	DisplayName string `json:"display_name,omitempty"`
	/*
	 high_is_good boolean */
	HighIsGood bool `json:"high_is_good,omitempty"`
	/*
	 icon_id integer */
	IconId int32 `json:"icon_id,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 published boolean */
	Published bool `json:"published,omitempty"`
	/*
	 stackable boolean */
	Stackable bool `json:"stackable,omitempty"`
	/*
	 unit_id integer */
	UnitId int32 `json:"unit_id,omitempty"`
}
