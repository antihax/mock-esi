package esilegacy

/*
200 ok object */
type GetCharactersCharacterIdMailLabels200Ok struct {
	/*
	 label_id integer */
	LabelId int32 `json:"label_id,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 color string */
	Color string `json:"color,omitempty"`
}