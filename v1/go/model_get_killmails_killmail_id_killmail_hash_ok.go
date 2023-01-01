package esiv1

import "time"

/*
200 ok object
*/
type GetKillmailsKillmailIdKillmailHashOk struct {
	/*
	 attackers array */
	Attackers []GetKillmailsKillmailIdKillmailHashAttacker `json:"attackers,omitempty"`
	/*
	 ID of the killmail */
	KillmailId int32 `json:"killmail_id,omitempty"`
	/*
	 Time that the victim was killed and the killmail generated  */
	KillmailTime time.Time `json:"killmail_time,omitempty"`
	/*
	 Moon if the kill took place at one */
	MoonId int32 `json:"moon_id,omitempty"`
	/*
	 Solar system that the kill took place in  */
	SolarSystemId int32 `json:"solar_system_id,omitempty"`
	/* */
	Victim GetKillmailsKillmailIdKillmailHashVictim `json:"victim,omitempty"`
	/*
	 War if the killmail is generated in relation to an official war  */
	WarId int32 `json:"war_id,omitempty"`
}
