package esidev

/*
200 ok object
*/
type GetCorporationsCorporationIdRoles200Ok struct {
	/*
	 character_id integer */
	CharacterId int32 `json:"character_id,omitempty"`
	/*
	 grantable_roles array */
	GrantableRoles []string `json:"grantable_roles,omitempty"`
	/*
	 grantable_roles_at_base array */
	GrantableRolesAtBase []string `json:"grantable_roles_at_base,omitempty"`
	/*
	 grantable_roles_at_hq array */
	GrantableRolesAtHq []string `json:"grantable_roles_at_hq,omitempty"`
	/*
	 grantable_roles_at_other array */
	GrantableRolesAtOther []string `json:"grantable_roles_at_other,omitempty"`
	/*
	 roles array */
	Roles []string `json:"roles,omitempty"`
	/*
	 roles_at_base array */
	RolesAtBase []string `json:"roles_at_base,omitempty"`
	/*
	 roles_at_hq array */
	RolesAtHq []string `json:"roles_at_hq,omitempty"`
	/*
	 roles_at_other array */
	RolesAtOther []string `json:"roles_at_other,omitempty"`
}
