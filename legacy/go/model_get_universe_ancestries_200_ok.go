package esilegacy

/*
200 ok object */
type GetUniverseAncestries200Ok struct {
	/*
	 id integer */
	Id int32 `json:"id,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 The bloodline associated with this ancestry */
	BloodlineId int32 `json:"bloodline_id,omitempty"`
	/*
	 description string */
	Description string `json:"description,omitempty"`
	/*
	 short_description string */
	ShortDescription string `json:"short_description,omitempty"`
	/*
	 icon_id integer */
	IconId int32 `json:"icon_id,omitempty"`
}
