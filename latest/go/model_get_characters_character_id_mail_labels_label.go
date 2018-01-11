package esilatest

/*
label object */
type GetCharactersCharacterIdMailLabelsLabel struct {
	/*
	 unread_count integer */
	UnreadCount int32 `json:"unread_count,omitempty"`
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
