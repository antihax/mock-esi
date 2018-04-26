package esilatest

import "time"

/*
200 ok object */
type GetCharactersCharacterIdAttributesOk struct {
	/*
	 Neural remapping cooldown after a character uses remap accrued over time */
	AccruedRemapCooldownDate time.Time `json:"accrued_remap_cooldown_date,omitempty"`
	/*
	 Number of available bonus character neural remaps */
	BonusRemaps int32 `json:"bonus_remaps,omitempty"`
	/*
	 charisma integer */
	Charisma int32 `json:"charisma,omitempty"`
	/*
	 intelligence integer */
	Intelligence int32 `json:"intelligence,omitempty"`
	/*
	 Datetime of last neural remap, including usage of bonus remaps */
	LastRemapDate time.Time `json:"last_remap_date,omitempty"`
	/*
	 memory integer */
	Memory int32 `json:"memory,omitempty"`
	/*
	 perception integer */
	Perception int32 `json:"perception,omitempty"`
	/*
	 willpower integer */
	Willpower int32 `json:"willpower,omitempty"`
}
