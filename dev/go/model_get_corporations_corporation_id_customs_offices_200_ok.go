package esidev

/*
200 ok object
*/
type GetCorporationsCorporationIdCustomsOffices200Ok struct {
	/*
	 Only present if alliance access is allowed */
	AllianceTaxRate float32 `json:"alliance_tax_rate,omitempty"`
	/*
	 standing_level and any standing related tax rate only present when this is true */
	AllowAccessWithStandings bool `json:"allow_access_with_standings,omitempty"`
	/*
	 allow_alliance_access boolean */
	AllowAllianceAccess bool `json:"allow_alliance_access,omitempty"`
	/*
	 bad_standing_tax_rate number */
	BadStandingTaxRate float32 `json:"bad_standing_tax_rate,omitempty"`
	/*
	 corporation_tax_rate number */
	CorporationTaxRate float32 `json:"corporation_tax_rate,omitempty"`
	/*
	 Tax rate for entities with excellent level of standing, only present if this level is allowed, same for all other standing related tax rates */
	ExcellentStandingTaxRate float32 `json:"excellent_standing_tax_rate,omitempty"`
	/*
	 good_standing_tax_rate number */
	GoodStandingTaxRate float32 `json:"good_standing_tax_rate,omitempty"`
	/*
	 neutral_standing_tax_rate number */
	NeutralStandingTaxRate float32 `json:"neutral_standing_tax_rate,omitempty"`
	/*
	 unique ID of this customs office */
	OfficeId int64 `json:"office_id,omitempty"`
	/*
	 reinforce_exit_end integer */
	ReinforceExitEnd int32 `json:"reinforce_exit_end,omitempty"`
	/*
	 Together with reinforce_exit_end, marks a 2-hour period where this customs office could exit reinforcement mode during the day after initial attack */
	ReinforceExitStart int32 `json:"reinforce_exit_start,omitempty"`
	/*
	 Access is allowed only for entities with this level of standing or better */
	StandingLevel string `json:"standing_level,omitempty"`
	/*
	 ID of the solar system this customs office is located in */
	SystemId int32 `json:"system_id,omitempty"`
	/*
	 terrible_standing_tax_rate number */
	TerribleStandingTaxRate float32 `json:"terrible_standing_tax_rate,omitempty"`
}
