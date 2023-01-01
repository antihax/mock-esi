package esidev

import "time"

/*
200 ok object
*/
type GetCorporationsCorporationIdStructures200Ok struct {
	/*
	 ID of the corporation that owns the structure */
	CorporationId int32 `json:"corporation_id,omitempty"`
	/*
	 Date on which the structure will run out of fuel */
	FuelExpires time.Time `json:"fuel_expires,omitempty"`
	/*
	 The structure name */
	Name string `json:"name,omitempty"`
	/*
	 The date and time when the structure's newly requested reinforcement times (e.g. next_reinforce_hour and next_reinforce_day) will take effect */
	NextReinforceApply time.Time `json:"next_reinforce_apply,omitempty"`
	/*
	 The requested change to reinforce_hour that will take effect at the time shown by next_reinforce_apply */
	NextReinforceHour int32 `json:"next_reinforce_hour,omitempty"`
	/*
	 The id of the ACL profile for this citadel */
	ProfileId int32 `json:"profile_id,omitempty"`
	/*
	 The hour of day that determines the four hour window when the structure will randomly exit its reinforcement periods and become vulnerable to attack against its armor and/or hull. The structure will become vulnerable at a random time that is +/- 2 hours centered on the value of this property */
	ReinforceHour int32 `json:"reinforce_hour,omitempty"`
	/*
	 Contains a list of service upgrades, and their state */
	Services []GetCorporationsCorporationIdStructuresService `json:"services,omitempty"`
	/*
	 state string */
	State string `json:"state,omitempty"`
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
