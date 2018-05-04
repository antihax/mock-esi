package esilegacy

/*
Optional object that is returned if a bookmark was made on a particular item. */
type GetCorporationsCorporationIdBookmarksItem struct {
	/*
	 item_id integer */
	ItemId int64 `json:"item_id,omitempty"`
	/*
	 type_id integer */
	TypeId int32 `json:"type_id,omitempty"`
}
