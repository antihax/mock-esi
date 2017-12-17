package esidev

/*
invitation object */
type PostFleetsFleetIdMembersInvitation struct {
	/*
	 The character you want to invite */
	CharacterId int32 `json:"character_id,omitempty"`
	/*
	 If a character is invited with the `fleet_commander` role, neither `wing_id` or `squad_id` should be specified. If a character is invited with the `wing_commander` role, only `wing_id` should be specified. If a character is invited with the `squad_commander` role, both `wing_id` and `squad_id` should be specified. If a character is invited with the `squad_member` role, `wing_id` and `squad_id` should either both be specified or not specified at all. If they aren’t specified, the invited character will join any squad with available positions. */
	Role string `json:"role,omitempty"`
	/*
	 wing_id integer */
	WingId int64 `json:"wing_id,omitempty"`
	/*
	 squad_id integer */
	SquadId int64 `json:"squad_id,omitempty"`
}
