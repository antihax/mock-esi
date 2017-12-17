package esidev

/*
item object */
type GetKillmailsKillmailIdKillmailHashItem struct {
	/*
	 item_type_id integer */
	ItemTypeId int32 `json:"item_type_id,omitempty"`
	/*
	 quantity_destroyed integer */
	QuantityDestroyed int64 `json:"quantity_destroyed,omitempty"`
	/*
	 quantity_dropped integer */
	QuantityDropped int64 `json:"quantity_dropped,omitempty"`
	/*
	 singleton integer */
	Singleton int32 `json:"singleton,omitempty"`
	/*
	 flag integer */
	Flag int32 `json:"flag,omitempty"`
}
