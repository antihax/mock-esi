package esilatest



/* 
movement object */
type PutFleetsFleetIdMembersMemberIdMovement struct {
/*
	 If a character is moved to the `fleet_commander` role, neither `wing_id` or `squad_id` should be specified. If a character is moved to the `wing_commander` role, only `wing_id` should be specified. If a character is moved to the `squad_commander` role, both `wing_id` and `squad_id` should be specified. If a character is moved to the `squad_member` role, both `wing_id` and `squad_id` should be specified. */
	Role string `json:"role,omitempty"`
/*
	 squad_id integer */
	SquadId int64 `json:"squad_id,omitempty"`
/*
	 wing_id integer */
	WingId int64 `json:"wing_id,omitempty"`
}
