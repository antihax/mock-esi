package esilatest



/* 
combat object */
type GetCharactersCharacterIdStatsCombat struct {
/*
	 cap_drainedby_npc integer */
	CapDrainedbyNpc int64 `json:"cap_drainedby_npc,omitempty"`
/*
	 cap_drainedby_pc integer */
	CapDrainedbyPc int64 `json:"cap_drainedby_pc,omitempty"`
/*
	 cap_draining_pc integer */
	CapDrainingPc int64 `json:"cap_draining_pc,omitempty"`
/*
	 criminal_flag_set integer */
	CriminalFlagSet int64 `json:"criminal_flag_set,omitempty"`
/*
	 damage_from_np_cs_amount integer */
	DamageFromNpCsAmount int64 `json:"damage_from_np_cs_amount,omitempty"`
/*
	 damage_from_np_cs_num_shots integer */
	DamageFromNpCsNumShots int64 `json:"damage_from_np_cs_num_shots,omitempty"`
/*
	 damage_from_players_bomb_amount integer */
	DamageFromPlayersBombAmount int64 `json:"damage_from_players_bomb_amount,omitempty"`
/*
	 damage_from_players_bomb_num_shots integer */
	DamageFromPlayersBombNumShots int64 `json:"damage_from_players_bomb_num_shots,omitempty"`
/*
	 damage_from_players_combat_drone_amount integer */
	DamageFromPlayersCombatDroneAmount int64 `json:"damage_from_players_combat_drone_amount,omitempty"`
/*
	 damage_from_players_combat_drone_num_shots integer */
	DamageFromPlayersCombatDroneNumShots int64 `json:"damage_from_players_combat_drone_num_shots,omitempty"`
/*
	 damage_from_players_energy_amount integer */
	DamageFromPlayersEnergyAmount int64 `json:"damage_from_players_energy_amount,omitempty"`
/*
	 damage_from_players_energy_num_shots integer */
	DamageFromPlayersEnergyNumShots int64 `json:"damage_from_players_energy_num_shots,omitempty"`
/*
	 damage_from_players_fighter_bomber_amount integer */
	DamageFromPlayersFighterBomberAmount int64 `json:"damage_from_players_fighter_bomber_amount,omitempty"`
/*
	 damage_from_players_fighter_bomber_num_shots integer */
	DamageFromPlayersFighterBomberNumShots int64 `json:"damage_from_players_fighter_bomber_num_shots,omitempty"`
/*
	 damage_from_players_fighter_drone_amount integer */
	DamageFromPlayersFighterDroneAmount int64 `json:"damage_from_players_fighter_drone_amount,omitempty"`
/*
	 damage_from_players_fighter_drone_num_shots integer */
	DamageFromPlayersFighterDroneNumShots int64 `json:"damage_from_players_fighter_drone_num_shots,omitempty"`
/*
	 damage_from_players_hybrid_amount integer */
	DamageFromPlayersHybridAmount int64 `json:"damage_from_players_hybrid_amount,omitempty"`
/*
	 damage_from_players_hybrid_num_shots integer */
	DamageFromPlayersHybridNumShots int64 `json:"damage_from_players_hybrid_num_shots,omitempty"`
/*
	 damage_from_players_missile_amount integer */
	DamageFromPlayersMissileAmount int64 `json:"damage_from_players_missile_amount,omitempty"`
/*
	 damage_from_players_missile_num_shots integer */
	DamageFromPlayersMissileNumShots int64 `json:"damage_from_players_missile_num_shots,omitempty"`
/*
	 damage_from_players_projectile_amount integer */
	DamageFromPlayersProjectileAmount int64 `json:"damage_from_players_projectile_amount,omitempty"`
/*
	 damage_from_players_projectile_num_shots integer */
	DamageFromPlayersProjectileNumShots int64 `json:"damage_from_players_projectile_num_shots,omitempty"`
/*
	 damage_from_players_smart_bomb_amount integer */
	DamageFromPlayersSmartBombAmount int64 `json:"damage_from_players_smart_bomb_amount,omitempty"`
/*
	 damage_from_players_smart_bomb_num_shots integer */
	DamageFromPlayersSmartBombNumShots int64 `json:"damage_from_players_smart_bomb_num_shots,omitempty"`
/*
	 damage_from_players_super_amount integer */
	DamageFromPlayersSuperAmount int64 `json:"damage_from_players_super_amount,omitempty"`
/*
	 damage_from_players_super_num_shots integer */
	DamageFromPlayersSuperNumShots int64 `json:"damage_from_players_super_num_shots,omitempty"`
/*
	 damage_from_structures_total_amount integer */
	DamageFromStructuresTotalAmount int64 `json:"damage_from_structures_total_amount,omitempty"`
/*
	 damage_from_structures_total_num_shots integer */
	DamageFromStructuresTotalNumShots int64 `json:"damage_from_structures_total_num_shots,omitempty"`
/*
	 damage_to_players_bomb_amount integer */
	DamageToPlayersBombAmount int64 `json:"damage_to_players_bomb_amount,omitempty"`
/*
	 damage_to_players_bomb_num_shots integer */
	DamageToPlayersBombNumShots int64 `json:"damage_to_players_bomb_num_shots,omitempty"`
/*
	 damage_to_players_combat_drone_amount integer */
	DamageToPlayersCombatDroneAmount int64 `json:"damage_to_players_combat_drone_amount,omitempty"`
/*
	 damage_to_players_combat_drone_num_shots integer */
	DamageToPlayersCombatDroneNumShots int64 `json:"damage_to_players_combat_drone_num_shots,omitempty"`
/*
	 damage_to_players_energy_amount integer */
	DamageToPlayersEnergyAmount int64 `json:"damage_to_players_energy_amount,omitempty"`
/*
	 damage_to_players_energy_num_shots integer */
	DamageToPlayersEnergyNumShots int64 `json:"damage_to_players_energy_num_shots,omitempty"`
/*
	 damage_to_players_fighter_bomber_amount integer */
	DamageToPlayersFighterBomberAmount int64 `json:"damage_to_players_fighter_bomber_amount,omitempty"`
/*
	 damage_to_players_fighter_bomber_num_shots integer */
	DamageToPlayersFighterBomberNumShots int64 `json:"damage_to_players_fighter_bomber_num_shots,omitempty"`
/*
	 damage_to_players_fighter_drone_amount integer */
	DamageToPlayersFighterDroneAmount int64 `json:"damage_to_players_fighter_drone_amount,omitempty"`
/*
	 damage_to_players_fighter_drone_num_shots integer */
	DamageToPlayersFighterDroneNumShots int64 `json:"damage_to_players_fighter_drone_num_shots,omitempty"`
/*
	 damage_to_players_hybrid_amount integer */
	DamageToPlayersHybridAmount int64 `json:"damage_to_players_hybrid_amount,omitempty"`
/*
	 damage_to_players_hybrid_num_shots integer */
	DamageToPlayersHybridNumShots int64 `json:"damage_to_players_hybrid_num_shots,omitempty"`
/*
	 damage_to_players_missile_amount integer */
	DamageToPlayersMissileAmount int64 `json:"damage_to_players_missile_amount,omitempty"`
/*
	 damage_to_players_missile_num_shots integer */
	DamageToPlayersMissileNumShots int64 `json:"damage_to_players_missile_num_shots,omitempty"`
/*
	 damage_to_players_projectile_amount integer */
	DamageToPlayersProjectileAmount int64 `json:"damage_to_players_projectile_amount,omitempty"`
/*
	 damage_to_players_projectile_num_shots integer */
	DamageToPlayersProjectileNumShots int64 `json:"damage_to_players_projectile_num_shots,omitempty"`
/*
	 damage_to_players_smart_bomb_amount integer */
	DamageToPlayersSmartBombAmount int64 `json:"damage_to_players_smart_bomb_amount,omitempty"`
/*
	 damage_to_players_smart_bomb_num_shots integer */
	DamageToPlayersSmartBombNumShots int64 `json:"damage_to_players_smart_bomb_num_shots,omitempty"`
/*
	 damage_to_players_super_amount integer */
	DamageToPlayersSuperAmount int64 `json:"damage_to_players_super_amount,omitempty"`
/*
	 damage_to_players_super_num_shots integer */
	DamageToPlayersSuperNumShots int64 `json:"damage_to_players_super_num_shots,omitempty"`
/*
	 damage_to_structures_total_amount integer */
	DamageToStructuresTotalAmount int64 `json:"damage_to_structures_total_amount,omitempty"`
/*
	 damage_to_structures_total_num_shots integer */
	DamageToStructuresTotalNumShots int64 `json:"damage_to_structures_total_num_shots,omitempty"`
/*
	 deaths_high_sec integer */
	DeathsHighSec int64 `json:"deaths_high_sec,omitempty"`
/*
	 deaths_low_sec integer */
	DeathsLowSec int64 `json:"deaths_low_sec,omitempty"`
/*
	 deaths_null_sec integer */
	DeathsNullSec int64 `json:"deaths_null_sec,omitempty"`
/*
	 deaths_pod_high_sec integer */
	DeathsPodHighSec int64 `json:"deaths_pod_high_sec,omitempty"`
/*
	 deaths_pod_low_sec integer */
	DeathsPodLowSec int64 `json:"deaths_pod_low_sec,omitempty"`
/*
	 deaths_pod_null_sec integer */
	DeathsPodNullSec int64 `json:"deaths_pod_null_sec,omitempty"`
/*
	 deaths_pod_wormhole integer */
	DeathsPodWormhole int64 `json:"deaths_pod_wormhole,omitempty"`
/*
	 deaths_wormhole integer */
	DeathsWormhole int64 `json:"deaths_wormhole,omitempty"`
/*
	 drone_engage integer */
	DroneEngage int64 `json:"drone_engage,omitempty"`
/*
	 dscans integer */
	Dscans int64 `json:"dscans,omitempty"`
/*
	 duel_requested integer */
	DuelRequested int64 `json:"duel_requested,omitempty"`
/*
	 engagement_register integer */
	EngagementRegister int64 `json:"engagement_register,omitempty"`
/*
	 kills_assists integer */
	KillsAssists int64 `json:"kills_assists,omitempty"`
/*
	 kills_high_sec integer */
	KillsHighSec int64 `json:"kills_high_sec,omitempty"`
/*
	 kills_low_sec integer */
	KillsLowSec int64 `json:"kills_low_sec,omitempty"`
/*
	 kills_null_sec integer */
	KillsNullSec int64 `json:"kills_null_sec,omitempty"`
/*
	 kills_pod_high_sec integer */
	KillsPodHighSec int64 `json:"kills_pod_high_sec,omitempty"`
/*
	 kills_pod_low_sec integer */
	KillsPodLowSec int64 `json:"kills_pod_low_sec,omitempty"`
/*
	 kills_pod_null_sec integer */
	KillsPodNullSec int64 `json:"kills_pod_null_sec,omitempty"`
/*
	 kills_pod_wormhole integer */
	KillsPodWormhole int64 `json:"kills_pod_wormhole,omitempty"`
/*
	 kills_wormhole integer */
	KillsWormhole int64 `json:"kills_wormhole,omitempty"`
/*
	 npc_flag_set integer */
	NpcFlagSet int64 `json:"npc_flag_set,omitempty"`
/*
	 probe_scans integer */
	ProbeScans int64 `json:"probe_scans,omitempty"`
/*
	 pvp_flag_set integer */
	PvpFlagSet int64 `json:"pvp_flag_set,omitempty"`
/*
	 repair_armor_by_remote_amount integer */
	RepairArmorByRemoteAmount int64 `json:"repair_armor_by_remote_amount,omitempty"`
/*
	 repair_armor_remote_amount integer */
	RepairArmorRemoteAmount int64 `json:"repair_armor_remote_amount,omitempty"`
/*
	 repair_armor_self_amount integer */
	RepairArmorSelfAmount int64 `json:"repair_armor_self_amount,omitempty"`
/*
	 repair_capacitor_by_remote_amount integer */
	RepairCapacitorByRemoteAmount int64 `json:"repair_capacitor_by_remote_amount,omitempty"`
/*
	 repair_capacitor_remote_amount integer */
	RepairCapacitorRemoteAmount int64 `json:"repair_capacitor_remote_amount,omitempty"`
/*
	 repair_capacitor_self_amount integer */
	RepairCapacitorSelfAmount int64 `json:"repair_capacitor_self_amount,omitempty"`
/*
	 repair_hull_by_remote_amount integer */
	RepairHullByRemoteAmount int64 `json:"repair_hull_by_remote_amount,omitempty"`
/*
	 repair_hull_remote_amount integer */
	RepairHullRemoteAmount int64 `json:"repair_hull_remote_amount,omitempty"`
/*
	 repair_hull_self_amount integer */
	RepairHullSelfAmount int64 `json:"repair_hull_self_amount,omitempty"`
/*
	 repair_shield_by_remote_amount integer */
	RepairShieldByRemoteAmount int64 `json:"repair_shield_by_remote_amount,omitempty"`
/*
	 repair_shield_remote_amount integer */
	RepairShieldRemoteAmount int64 `json:"repair_shield_remote_amount,omitempty"`
/*
	 repair_shield_self_amount integer */
	RepairShieldSelfAmount int64 `json:"repair_shield_self_amount,omitempty"`
/*
	 self_destructs integer */
	SelfDestructs int64 `json:"self_destructs,omitempty"`
/*
	 warp_scramble_pc integer */
	WarpScramblePc int64 `json:"warp_scramble_pc,omitempty"`
/*
	 warp_scrambledby_npc integer */
	WarpScrambledbyNpc int64 `json:"warp_scrambledby_npc,omitempty"`
/*
	 warp_scrambledby_pc integer */
	WarpScrambledbyPc int64 `json:"warp_scrambledby_pc,omitempty"`
/*
	 weapon_flag_set integer */
	WeaponFlagSet int64 `json:"weapon_flag_set,omitempty"`
/*
	 webifiedby_npc integer */
	WebifiedbyNpc int64 `json:"webifiedby_npc,omitempty"`
/*
	 webifiedby_pc integer */
	WebifiedbyPc int64 `json:"webifiedby_pc,omitempty"`
/*
	 webifying_pc integer */
	WebifyingPc int64 `json:"webifying_pc,omitempty"`
}
