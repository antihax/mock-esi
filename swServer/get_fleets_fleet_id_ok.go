package swServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetFleetsFleetIdOk struct {
/*
	 Is free-move enabled */
	is_free_move bool `json:"is_free_move,omitempty"`
/*
	 Does the fleet have an active fleet advertisement */
	is_registered bool `json:"is_registered,omitempty"`
/*
	 Is EVE Voice enabled */
	is_voice_enabled bool `json:"is_voice_enabled,omitempty"`
/*
	 Fleet MOTD in CCP flavoured HTML */
	motd string `json:"motd,omitempty"`
}
