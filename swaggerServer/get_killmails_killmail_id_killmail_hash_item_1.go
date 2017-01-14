package swaggerServer

import "time"
var _ time.Time

/* 
item object */
type GetKillmailsKillmailIdKillmailHashItem1 struct {
/*
	 Flag for the location of the item
 */
	flag int32 `json:"flag,omitempty"`
/*
	 item_type_id integer */
	item_type_id int32 `json:"item_type_id,omitempty"`
/*
	 items array */
	items []GetKillmailsKillmailIdKillmailHashItem `json:"items,omitempty"`
/*
	 How many of the item were destroyed if any
 */
	quantity_destroyed int64 `json:"quantity_destroyed,omitempty"`
/*
	 How many of the item were dropped if any
 */
	quantity_dropped int64 `json:"quantity_dropped,omitempty"`
/*
	 singleton integer */
	singleton int32 `json:"singleton,omitempty"`
}
