package swaggerServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetKillmailsKillmailIdKillmailHashOk struct {
/*
	 attackers array */
	attackers []GetKillmailsKillmailIdKillmailHashAttacker `json:"attackers,omitempty"`
/*
	 ID of the killmail */
	killmail_id int32 `json:"killmail_id,omitempty"`
/*
	 Time that the victim was killed and the killmail generated
 */
	killmail_time time.Time `json:"killmail_time,omitempty"`
/*
	 Moon if the kill took place at one */
	moon_id int32 `json:"moon_id,omitempty"`
/*
	 Solar system that the kill took place in
 */
	solar_system_id int32 `json:"solar_system_id,omitempty"`
/* */
	victim GetKillmailsKillmailIdKillmailHashVictim `json:"victim,omitempty"`
/*
	 War if the killmail is generated in relation to an official war
 */
	war_id int32 `json:"war_id,omitempty"`
}
