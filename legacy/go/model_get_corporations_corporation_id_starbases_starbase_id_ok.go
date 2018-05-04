package esilegacy

/*
200 ok object */
type GetCorporationsCorporationIdStarbasesStarbaseIdOk struct {
	/*
	 allow_alliance_members boolean */
	AllowAllianceMembers bool `json:"allow_alliance_members,omitempty"`
	/*
	 allow_corporation_members boolean */
	AllowCorporationMembers bool `json:"allow_corporation_members,omitempty"`
	/*
	 Who can anchor starbase (POS) and its structures */
	Anchor string `json:"anchor,omitempty"`
	/*
	 attack_if_at_war boolean */
	AttackIfAtWar bool `json:"attack_if_at_war,omitempty"`
	/*
	 attack_if_other_security_status_dropping boolean */
	AttackIfOtherSecurityStatusDropping bool `json:"attack_if_other_security_status_dropping,omitempty"`
	/*
	 Starbase (POS) will attack if target's security standing is lower than this value */
	AttackSecurityStatusThreshold float32 `json:"attack_security_status_threshold,omitempty"`
	/*
	 Starbase (POS) will attack if target's standing is lower than this value */
	AttackStandingThreshold float32 `json:"attack_standing_threshold,omitempty"`
	/*
	 Who can take fuel blocks out of the starbase (POS)'s fuel bay */
	FuelBayTake string `json:"fuel_bay_take,omitempty"`
	/*
	 Who can view the starbase (POS)'s fule bay. Characters either need to have required role or belong to the starbase (POS) owner's corporation or alliance, as described by the enum, all other access settings follows the same scheme */
	FuelBayView string `json:"fuel_bay_view,omitempty"`
	/*
	 Fuel blocks and other things that will be consumed when operating a starbase (POS) */
	Fuels []GetCorporationsCorporationIdStarbasesStarbaseIdFuel `json:"fuels,omitempty"`
	/*
	 Who can offline starbase (POS) and its structures */
	Offline string `json:"offline,omitempty"`
	/*
	 Who can online starbase (POS) and its structures */
	Online string `json:"online,omitempty"`
	/*
	 Who can unanchor starbase (POS) and its structures */
	Unanchor string `json:"unanchor,omitempty"`
	/*
	 True if the starbase (POS) is using alliance standings, otherwise using corporation's */
	UseAllianceStandings bool `json:"use_alliance_standings,omitempty"`
}
