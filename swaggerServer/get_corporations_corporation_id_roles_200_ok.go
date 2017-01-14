package swaggerServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetCorporationsCorporationIdRoles200Ok struct {
/*
	 character_id integer */
	character_id int32 `json:"character_id,omitempty"`
/*
	 grantable_roles array */
	grantable_roles []string `json:"grantable_roles,omitempty"`
/*
	 grantable_roles_at_base array */
	grantable_roles_at_base []string `json:"grantable_roles_at_base,omitempty"`
/*
	 grantable_roles_at_hq array */
	grantable_roles_at_hq []string `json:"grantable_roles_at_hq,omitempty"`
/*
	 grantable_roles_at_other array */
	grantable_roles_at_other []string `json:"grantable_roles_at_other,omitempty"`
/*
	 roles array */
	roles []string `json:"roles,omitempty"`
/*
	 roles_at_base array */
	roles_at_base []string `json:"roles_at_base,omitempty"`
/*
	 roles_at_hq array */
	roles_at_hq []string `json:"roles_at_hq,omitempty"`
/*
	 roles_at_other array */
	roles_at_other []string `json:"roles_at_other,omitempty"`
}
