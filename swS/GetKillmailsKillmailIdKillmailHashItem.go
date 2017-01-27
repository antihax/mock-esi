package swS

import "time"
var _ time.Time

/* 
item object */
type GetKillmailsKillmailIdKillmailHashItem struct {
/*
	 flag integer */
	flag int32 `json:"flag,omitempty"`
/*
	 item_type_id integer */
	item_type_id int32 `json:"item_type_id,omitempty"`
/*
	 quantity_destroyed integer */
	quantity_destroyed int64 `json:"quantity_destroyed,omitempty"`
/*
	 quantity_dropped integer */
	quantity_dropped int64 `json:"quantity_dropped,omitempty"`
/*
	 singleton integer */
	singleton int32 `json:"singleton,omitempty"`
}
