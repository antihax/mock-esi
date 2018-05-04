package esiv2

/*
orbital object */
type GetCharactersCharacterIdStatsOrbital struct {
	/*
	 strike_characters_killed integer */
	StrikeCharactersKilled int64 `json:"strike_characters_killed,omitempty"`
	/*
	 strike_damage_to_players_armor_amount integer */
	StrikeDamageToPlayersArmorAmount int64 `json:"strike_damage_to_players_armor_amount,omitempty"`
	/*
	 strike_damage_to_players_shield_amount integer */
	StrikeDamageToPlayersShieldAmount int64 `json:"strike_damage_to_players_shield_amount,omitempty"`
}
