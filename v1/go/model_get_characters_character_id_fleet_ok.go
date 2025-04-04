package esiv1

/*
200 ok object
*/
type GetCharactersCharacterIdFleetOk struct {
	/*
	 Character ID of the current fleet boss */
	FleetBossId int64 `json:"fleet_boss_id,omitempty"`
	/*
	 The character's current fleet ID */
	FleetId int64 `json:"fleet_id,omitempty"`
	/*
	 Member’s role in fleet */
	Role string `json:"role,omitempty"`
	/*
	 ID of the squad the member is in. If not applicable, will be set to -1 */
	SquadId int64 `json:"squad_id,omitempty"`
	/*
	 ID of the wing the member is in. If not applicable, will be set to -1 */
	WingId int64 `json:"wing_id,omitempty"`
}
