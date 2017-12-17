package esidev

import "time"

/*
200 ok object */
type GetCorporationsCorporationIdRolesHistory200Ok struct {
	/*
	 The character whose roles are changed */
	CharacterId int32 `json:"character_id,omitempty"`
	/*
	 changed_at string */
	ChangedAt time.Time `json:"changed_at,omitempty"`
	/*
	 ID of the character who issued this change */
	IssuerId int32 `json:"issuer_id,omitempty"`
	/*
	 role_type string */
	RoleType string `json:"role_type,omitempty"`
	/*
	 old_roles array */
	OldRoles []string `json:"old_roles,omitempty"`
	/*
	 new_roles array */
	NewRoles []string `json:"new_roles,omitempty"`
}
