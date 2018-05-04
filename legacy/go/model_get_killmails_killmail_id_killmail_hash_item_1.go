package esilegacy



/* 
item object */
type GetKillmailsKillmailIdKillmailHashItem1 struct {
/*
	 Flag for the location of the item  */
	Flag int32 `json:"flag,omitempty"`
/*
	 item_type_id integer */
	ItemTypeId int32 `json:"item_type_id,omitempty"`
/*
	 items array */
	Items []GetKillmailsKillmailIdKillmailHashItem `json:"items,omitempty"`
/*
	 How many of the item were destroyed if any  */
	QuantityDestroyed int64 `json:"quantity_destroyed,omitempty"`
/*
	 How many of the item were dropped if any  */
	QuantityDropped int64 `json:"quantity_dropped,omitempty"`
/*
	 singleton integer */
	Singleton int32 `json:"singleton,omitempty"`
}
