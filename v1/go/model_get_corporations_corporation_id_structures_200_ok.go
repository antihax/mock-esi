package esiv1

import "time"

/*
200 ok object */
type GetCorporationsCorporationIdStructures200Ok struct {
	/*
	 ID of the corporation that owns the structure */
	CorporationId int32 `json:"corporation_id,omitempty"`
	/*
	 This week's vulnerability windows, Monday is day 0 */
	CurrentVul []GetCorporationsCorporationIdStructuresCurrentVulCurrentVul `json:"current_vul,omitempty"`
	/*
	 Date on which the structure will run out of fuel */
	FuelExpires time.Time `json:"fuel_expires,omitempty"`
	/*
	 Next week's vulnerability windows, Monday is day 0 */
	NextVul []GetCorporationsCorporationIdStructuresNextVulNextVul `json:"next_vul,omitempty"`
	/*
	 The id of the ACL profile for this citadel */
	ProfileId int32 `json:"profile_id,omitempty"`
	/*
	 Contains a list of service upgrades, and their state */
	Services []GetCorporationsCorporationIdStructuresService `json:"services,omitempty"`
	/*
	 Date at which the structure will move to it's next state */
	StateTimerEnd time.Time `json:"state_timer_end,omitempty"`
	/*
	 Date at which the structure entered it's current state */
	StateTimerStart time.Time `json:"state_timer_start,omitempty"`
	/*
	 The Item ID of the structure */
	StructureId int64 `json:"structure_id,omitempty"`
	/*
	 The solar system the structure is in */
	SystemId int32 `json:"system_id,omitempty"`
	/*
	 The type id of the structure */
	TypeId int32 `json:"type_id,omitempty"`
	/*
	 Date at which the structure will unanchor */
	UnanchorsAt time.Time `json:"unanchors_at,omitempty"`
}
