package swaggerServer

import "time"
var _ time.Time

/* 
new_settings object */
type PutFleetsFleetIdNewSettings struct {
/*
	 Should free-move be enabled in the fleet */
	is_free_move bool `json:"is_free_move,omitempty"`
/*
	 New fleet MOTD in CCP flavoured HTML */
	motd string `json:"motd,omitempty"`
}
