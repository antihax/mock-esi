package esiv1



/* 
Aggregate stats for a year */
type GetCharactersCharacterIdStats200Ok struct {
/*
	 character_minutes integer */
	CharacterMinutes int64 `json:"character_minutes,omitempty"`
/*
	 character_sessions_started integer */
	CharacterSessionsStarted int64 `json:"character_sessions_started,omitempty"`
/*
	 combat_cap_drainedby_npc integer */
	CombatCapDrainedbyNpc int64 `json:"combat_cap_drainedby_npc,omitempty"`
/*
	 combat_cap_drainedby_pc integer */
	CombatCapDrainedbyPc int64 `json:"combat_cap_drainedby_pc,omitempty"`
/*
	 combat_cap_draining_pc integer */
	CombatCapDrainingPc int64 `json:"combat_cap_draining_pc,omitempty"`
/*
	 combat_criminal_flag_set integer */
	CombatCriminalFlagSet int64 `json:"combat_criminal_flag_set,omitempty"`
/*
	 combat_damage_from_np_cs_amount integer */
	CombatDamageFromNpCsAmount int64 `json:"combat_damage_from_np_cs_amount,omitempty"`
/*
	 combat_damage_from_np_cs_num_shots integer */
	CombatDamageFromNpCsNumShots int64 `json:"combat_damage_from_np_cs_num_shots,omitempty"`
/*
	 combat_damage_from_players_bomb_amount integer */
	CombatDamageFromPlayersBombAmount int64 `json:"combat_damage_from_players_bomb_amount,omitempty"`
/*
	 combat_damage_from_players_bomb_num_shots integer */
	CombatDamageFromPlayersBombNumShots int64 `json:"combat_damage_from_players_bomb_num_shots,omitempty"`
/*
	 combat_damage_from_players_combat_drone_amount integer */
	CombatDamageFromPlayersCombatDroneAmount int64 `json:"combat_damage_from_players_combat_drone_amount,omitempty"`
/*
	 combat_damage_from_players_combat_drone_num_shots integer */
	CombatDamageFromPlayersCombatDroneNumShots int64 `json:"combat_damage_from_players_combat_drone_num_shots,omitempty"`
/*
	 combat_damage_from_players_energy_amount integer */
	CombatDamageFromPlayersEnergyAmount int64 `json:"combat_damage_from_players_energy_amount,omitempty"`
/*
	 combat_damage_from_players_energy_num_shots integer */
	CombatDamageFromPlayersEnergyNumShots int64 `json:"combat_damage_from_players_energy_num_shots,omitempty"`
/*
	 combat_damage_from_players_fighter_bomber_amount integer */
	CombatDamageFromPlayersFighterBomberAmount int64 `json:"combat_damage_from_players_fighter_bomber_amount,omitempty"`
/*
	 combat_damage_from_players_fighter_bomber_num_shots integer */
	CombatDamageFromPlayersFighterBomberNumShots int64 `json:"combat_damage_from_players_fighter_bomber_num_shots,omitempty"`
/*
	 combat_damage_from_players_fighter_drone_amount integer */
	CombatDamageFromPlayersFighterDroneAmount int64 `json:"combat_damage_from_players_fighter_drone_amount,omitempty"`
/*
	 combat_damage_from_players_fighter_drone_num_shots integer */
	CombatDamageFromPlayersFighterDroneNumShots int64 `json:"combat_damage_from_players_fighter_drone_num_shots,omitempty"`
/*
	 combat_damage_from_players_hybrid_amount integer */
	CombatDamageFromPlayersHybridAmount int64 `json:"combat_damage_from_players_hybrid_amount,omitempty"`
/*
	 combat_damage_from_players_hybrid_num_shots integer */
	CombatDamageFromPlayersHybridNumShots int64 `json:"combat_damage_from_players_hybrid_num_shots,omitempty"`
/*
	 combat_damage_from_players_missile_amount integer */
	CombatDamageFromPlayersMissileAmount int64 `json:"combat_damage_from_players_missile_amount,omitempty"`
/*
	 combat_damage_from_players_missile_num_shots integer */
	CombatDamageFromPlayersMissileNumShots int64 `json:"combat_damage_from_players_missile_num_shots,omitempty"`
/*
	 combat_damage_from_players_projectile_amount integer */
	CombatDamageFromPlayersProjectileAmount int64 `json:"combat_damage_from_players_projectile_amount,omitempty"`
/*
	 combat_damage_from_players_projectile_num_shots integer */
	CombatDamageFromPlayersProjectileNumShots int64 `json:"combat_damage_from_players_projectile_num_shots,omitempty"`
/*
	 combat_damage_from_players_smart_bomb_amount integer */
	CombatDamageFromPlayersSmartBombAmount int64 `json:"combat_damage_from_players_smart_bomb_amount,omitempty"`
/*
	 combat_damage_from_players_smart_bomb_num_shots integer */
	CombatDamageFromPlayersSmartBombNumShots int64 `json:"combat_damage_from_players_smart_bomb_num_shots,omitempty"`
/*
	 combat_damage_from_players_super_amount integer */
	CombatDamageFromPlayersSuperAmount int64 `json:"combat_damage_from_players_super_amount,omitempty"`
/*
	 combat_damage_from_players_super_num_shots integer */
	CombatDamageFromPlayersSuperNumShots int64 `json:"combat_damage_from_players_super_num_shots,omitempty"`
/*
	 combat_damage_from_structures_total_amount integer */
	CombatDamageFromStructuresTotalAmount int64 `json:"combat_damage_from_structures_total_amount,omitempty"`
/*
	 combat_damage_from_structures_total_num_shots integer */
	CombatDamageFromStructuresTotalNumShots int64 `json:"combat_damage_from_structures_total_num_shots,omitempty"`
/*
	 combat_damage_to_players_bomb_amount integer */
	CombatDamageToPlayersBombAmount int64 `json:"combat_damage_to_players_bomb_amount,omitempty"`
/*
	 combat_damage_to_players_bomb_num_shots integer */
	CombatDamageToPlayersBombNumShots int64 `json:"combat_damage_to_players_bomb_num_shots,omitempty"`
/*
	 combat_damage_to_players_combat_drone_amount integer */
	CombatDamageToPlayersCombatDroneAmount int64 `json:"combat_damage_to_players_combat_drone_amount,omitempty"`
/*
	 combat_damage_to_players_combat_drone_num_shots integer */
	CombatDamageToPlayersCombatDroneNumShots int64 `json:"combat_damage_to_players_combat_drone_num_shots,omitempty"`
/*
	 combat_damage_to_players_energy_amount integer */
	CombatDamageToPlayersEnergyAmount int64 `json:"combat_damage_to_players_energy_amount,omitempty"`
/*
	 combat_damage_to_players_energy_num_shots integer */
	CombatDamageToPlayersEnergyNumShots int64 `json:"combat_damage_to_players_energy_num_shots,omitempty"`
/*
	 combat_damage_to_players_fighter_bomber_amount integer */
	CombatDamageToPlayersFighterBomberAmount int64 `json:"combat_damage_to_players_fighter_bomber_amount,omitempty"`
/*
	 combat_damage_to_players_fighter_bomber_num_shots integer */
	CombatDamageToPlayersFighterBomberNumShots int64 `json:"combat_damage_to_players_fighter_bomber_num_shots,omitempty"`
/*
	 combat_damage_to_players_fighter_drone_amount integer */
	CombatDamageToPlayersFighterDroneAmount int64 `json:"combat_damage_to_players_fighter_drone_amount,omitempty"`
/*
	 combat_damage_to_players_fighter_drone_num_shots integer */
	CombatDamageToPlayersFighterDroneNumShots int64 `json:"combat_damage_to_players_fighter_drone_num_shots,omitempty"`
/*
	 combat_damage_to_players_hybrid_amount integer */
	CombatDamageToPlayersHybridAmount int64 `json:"combat_damage_to_players_hybrid_amount,omitempty"`
/*
	 combat_damage_to_players_hybrid_num_shots integer */
	CombatDamageToPlayersHybridNumShots int64 `json:"combat_damage_to_players_hybrid_num_shots,omitempty"`
/*
	 combat_damage_to_players_missile_amount integer */
	CombatDamageToPlayersMissileAmount int64 `json:"combat_damage_to_players_missile_amount,omitempty"`
/*
	 combat_damage_to_players_missile_num_shots integer */
	CombatDamageToPlayersMissileNumShots int64 `json:"combat_damage_to_players_missile_num_shots,omitempty"`
/*
	 combat_damage_to_players_projectile_amount integer */
	CombatDamageToPlayersProjectileAmount int64 `json:"combat_damage_to_players_projectile_amount,omitempty"`
/*
	 combat_damage_to_players_projectile_num_shots integer */
	CombatDamageToPlayersProjectileNumShots int64 `json:"combat_damage_to_players_projectile_num_shots,omitempty"`
/*
	 combat_damage_to_players_smart_bomb_amount integer */
	CombatDamageToPlayersSmartBombAmount int64 `json:"combat_damage_to_players_smart_bomb_amount,omitempty"`
/*
	 combat_damage_to_players_smart_bomb_num_shots integer */
	CombatDamageToPlayersSmartBombNumShots int64 `json:"combat_damage_to_players_smart_bomb_num_shots,omitempty"`
/*
	 combat_damage_to_players_super_amount integer */
	CombatDamageToPlayersSuperAmount int64 `json:"combat_damage_to_players_super_amount,omitempty"`
/*
	 combat_damage_to_players_super_num_shots integer */
	CombatDamageToPlayersSuperNumShots int64 `json:"combat_damage_to_players_super_num_shots,omitempty"`
/*
	 combat_damage_to_structures_total_amount integer */
	CombatDamageToStructuresTotalAmount int64 `json:"combat_damage_to_structures_total_amount,omitempty"`
/*
	 combat_damage_to_structures_total_num_shots integer */
	CombatDamageToStructuresTotalNumShots int64 `json:"combat_damage_to_structures_total_num_shots,omitempty"`
/*
	 combat_deaths_high_sec integer */
	CombatDeathsHighSec int64 `json:"combat_deaths_high_sec,omitempty"`
/*
	 combat_deaths_low_sec integer */
	CombatDeathsLowSec int64 `json:"combat_deaths_low_sec,omitempty"`
/*
	 combat_deaths_null_sec integer */
	CombatDeathsNullSec int64 `json:"combat_deaths_null_sec,omitempty"`
/*
	 combat_deaths_pod_high_sec integer */
	CombatDeathsPodHighSec int64 `json:"combat_deaths_pod_high_sec,omitempty"`
/*
	 combat_deaths_pod_low_sec integer */
	CombatDeathsPodLowSec int64 `json:"combat_deaths_pod_low_sec,omitempty"`
/*
	 combat_deaths_pod_null_sec integer */
	CombatDeathsPodNullSec int64 `json:"combat_deaths_pod_null_sec,omitempty"`
/*
	 combat_deaths_pod_wormhole integer */
	CombatDeathsPodWormhole int64 `json:"combat_deaths_pod_wormhole,omitempty"`
/*
	 combat_deaths_wormhole integer */
	CombatDeathsWormhole int64 `json:"combat_deaths_wormhole,omitempty"`
/*
	 combat_drone_engage integer */
	CombatDroneEngage int64 `json:"combat_drone_engage,omitempty"`
/*
	 combat_duel_requested integer */
	CombatDuelRequested int64 `json:"combat_duel_requested,omitempty"`
/*
	 combat_engagement_register integer */
	CombatEngagementRegister int64 `json:"combat_engagement_register,omitempty"`
/*
	 combat_kills_assists integer */
	CombatKillsAssists int64 `json:"combat_kills_assists,omitempty"`
/*
	 combat_kills_high_sec integer */
	CombatKillsHighSec int64 `json:"combat_kills_high_sec,omitempty"`
/*
	 combat_kills_low_sec integer */
	CombatKillsLowSec int64 `json:"combat_kills_low_sec,omitempty"`
/*
	 combat_kills_null_sec integer */
	CombatKillsNullSec int64 `json:"combat_kills_null_sec,omitempty"`
/*
	 combat_kills_pod_high_sec integer */
	CombatKillsPodHighSec int64 `json:"combat_kills_pod_high_sec,omitempty"`
/*
	 combat_kills_pod_low_sec integer */
	CombatKillsPodLowSec int64 `json:"combat_kills_pod_low_sec,omitempty"`
/*
	 combat_kills_pod_null_sec integer */
	CombatKillsPodNullSec int64 `json:"combat_kills_pod_null_sec,omitempty"`
/*
	 combat_kills_pod_wormhole integer */
	CombatKillsPodWormhole int64 `json:"combat_kills_pod_wormhole,omitempty"`
/*
	 combat_kills_wormhole integer */
	CombatKillsWormhole int64 `json:"combat_kills_wormhole,omitempty"`
/*
	 combat_npc_flag_set integer */
	CombatNpcFlagSet int64 `json:"combat_npc_flag_set,omitempty"`
/*
	 combat_pvp_flag_set integer */
	CombatPvpFlagSet int64 `json:"combat_pvp_flag_set,omitempty"`
/*
	 combat_repair_armor_by_remote_amount integer */
	CombatRepairArmorByRemoteAmount int64 `json:"combat_repair_armor_by_remote_amount,omitempty"`
/*
	 combat_repair_armor_remote_amount integer */
	CombatRepairArmorRemoteAmount int64 `json:"combat_repair_armor_remote_amount,omitempty"`
/*
	 combat_repair_armor_self_amount integer */
	CombatRepairArmorSelfAmount int64 `json:"combat_repair_armor_self_amount,omitempty"`
/*
	 combat_repair_capacitor_by_remote_amount integer */
	CombatRepairCapacitorByRemoteAmount int64 `json:"combat_repair_capacitor_by_remote_amount,omitempty"`
/*
	 combat_repair_capacitor_remote_amount integer */
	CombatRepairCapacitorRemoteAmount int64 `json:"combat_repair_capacitor_remote_amount,omitempty"`
/*
	 combat_repair_capacitor_self_amount integer */
	CombatRepairCapacitorSelfAmount int64 `json:"combat_repair_capacitor_self_amount,omitempty"`
/*
	 combat_repair_hull_by_remote_amount integer */
	CombatRepairHullByRemoteAmount int64 `json:"combat_repair_hull_by_remote_amount,omitempty"`
/*
	 combat_repair_hull_remote_amount integer */
	CombatRepairHullRemoteAmount int64 `json:"combat_repair_hull_remote_amount,omitempty"`
/*
	 combat_repair_hull_self_amount integer */
	CombatRepairHullSelfAmount int64 `json:"combat_repair_hull_self_amount,omitempty"`
/*
	 combat_repair_shield_by_remote_amount integer */
	CombatRepairShieldByRemoteAmount int64 `json:"combat_repair_shield_by_remote_amount,omitempty"`
/*
	 combat_repair_shield_remote_amount integer */
	CombatRepairShieldRemoteAmount int64 `json:"combat_repair_shield_remote_amount,omitempty"`
/*
	 combat_repair_shield_self_amount integer */
	CombatRepairShieldSelfAmount int64 `json:"combat_repair_shield_self_amount,omitempty"`
/*
	 combat_self_destructs integer */
	CombatSelfDestructs int64 `json:"combat_self_destructs,omitempty"`
/*
	 combat_warp_scramble_pc integer */
	CombatWarpScramblePc int64 `json:"combat_warp_scramble_pc,omitempty"`
/*
	 combat_warp_scrambledby_npc integer */
	CombatWarpScrambledbyNpc int64 `json:"combat_warp_scrambledby_npc,omitempty"`
/*
	 combat_warp_scrambledby_pc integer */
	CombatWarpScrambledbyPc int64 `json:"combat_warp_scrambledby_pc,omitempty"`
/*
	 combat_weapon_flag_set integer */
	CombatWeaponFlagSet int64 `json:"combat_weapon_flag_set,omitempty"`
/*
	 combat_webifiedby_npc integer */
	CombatWebifiedbyNpc int64 `json:"combat_webifiedby_npc,omitempty"`
/*
	 combat_webifiedby_pc integer */
	CombatWebifiedbyPc int64 `json:"combat_webifiedby_pc,omitempty"`
/*
	 combat_webifying_pc integer */
	CombatWebifyingPc int64 `json:"combat_webifying_pc,omitempty"`
/*
	 days_of_activity integer */
	DaysOfActivity int64 `json:"days_of_activity,omitempty"`
/*
	 generic_cone_scans integer */
	GenericConeScans int64 `json:"generic_cone_scans,omitempty"`
/*
	 generic_request_scans integer */
	GenericRequestScans int64 `json:"generic_request_scans,omitempty"`
/*
	 industry_hacking_successes integer */
	IndustryHackingSuccesses int64 `json:"industry_hacking_successes,omitempty"`
/*
	 industry_jobs_cancelled integer */
	IndustryJobsCancelled int64 `json:"industry_jobs_cancelled,omitempty"`
/*
	 industry_jobs_completed_copy_blueprint integer */
	IndustryJobsCompletedCopyBlueprint int64 `json:"industry_jobs_completed_copy_blueprint,omitempty"`
/*
	 industry_jobs_completed_invention integer */
	IndustryJobsCompletedInvention int64 `json:"industry_jobs_completed_invention,omitempty"`
/*
	 industry_jobs_completed_manufacture integer */
	IndustryJobsCompletedManufacture int64 `json:"industry_jobs_completed_manufacture,omitempty"`
/*
	 industry_jobs_completed_manufacture_asteroid integer */
	IndustryJobsCompletedManufactureAsteroid int64 `json:"industry_jobs_completed_manufacture_asteroid,omitempty"`
/*
	 industry_jobs_completed_manufacture_asteroid_quantity integer */
	IndustryJobsCompletedManufactureAsteroidQuantity int64 `json:"industry_jobs_completed_manufacture_asteroid_quantity,omitempty"`
/*
	 industry_jobs_completed_manufacture_charge integer */
	IndustryJobsCompletedManufactureCharge int64 `json:"industry_jobs_completed_manufacture_charge,omitempty"`
/*
	 industry_jobs_completed_manufacture_charge_quantity integer */
	IndustryJobsCompletedManufactureChargeQuantity int64 `json:"industry_jobs_completed_manufacture_charge_quantity,omitempty"`
/*
	 industry_jobs_completed_manufacture_commodity integer */
	IndustryJobsCompletedManufactureCommodity int64 `json:"industry_jobs_completed_manufacture_commodity,omitempty"`
/*
	 industry_jobs_completed_manufacture_commodity_quantity integer */
	IndustryJobsCompletedManufactureCommodityQuantity int64 `json:"industry_jobs_completed_manufacture_commodity_quantity,omitempty"`
/*
	 industry_jobs_completed_manufacture_deployable integer */
	IndustryJobsCompletedManufactureDeployable int64 `json:"industry_jobs_completed_manufacture_deployable,omitempty"`
/*
	 industry_jobs_completed_manufacture_deployable_quantity integer */
	IndustryJobsCompletedManufactureDeployableQuantity int64 `json:"industry_jobs_completed_manufacture_deployable_quantity,omitempty"`
/*
	 industry_jobs_completed_manufacture_drone integer */
	IndustryJobsCompletedManufactureDrone int64 `json:"industry_jobs_completed_manufacture_drone,omitempty"`
/*
	 industry_jobs_completed_manufacture_drone_quantity integer */
	IndustryJobsCompletedManufactureDroneQuantity int64 `json:"industry_jobs_completed_manufacture_drone_quantity,omitempty"`
/*
	 industry_jobs_completed_manufacture_implant integer */
	IndustryJobsCompletedManufactureImplant int64 `json:"industry_jobs_completed_manufacture_implant,omitempty"`
/*
	 industry_jobs_completed_manufacture_implant_quantity integer */
	IndustryJobsCompletedManufactureImplantQuantity int64 `json:"industry_jobs_completed_manufacture_implant_quantity,omitempty"`
/*
	 industry_jobs_completed_manufacture_module integer */
	IndustryJobsCompletedManufactureModule int64 `json:"industry_jobs_completed_manufacture_module,omitempty"`
/*
	 industry_jobs_completed_manufacture_module_quantity integer */
	IndustryJobsCompletedManufactureModuleQuantity int64 `json:"industry_jobs_completed_manufacture_module_quantity,omitempty"`
/*
	 industry_jobs_completed_manufacture_other integer */
	IndustryJobsCompletedManufactureOther int64 `json:"industry_jobs_completed_manufacture_other,omitempty"`
/*
	 industry_jobs_completed_manufacture_other_quantity integer */
	IndustryJobsCompletedManufactureOtherQuantity int64 `json:"industry_jobs_completed_manufacture_other_quantity,omitempty"`
/*
	 industry_jobs_completed_manufacture_ship integer */
	IndustryJobsCompletedManufactureShip int64 `json:"industry_jobs_completed_manufacture_ship,omitempty"`
/*
	 industry_jobs_completed_manufacture_ship_quantity integer */
	IndustryJobsCompletedManufactureShipQuantity int64 `json:"industry_jobs_completed_manufacture_ship_quantity,omitempty"`
/*
	 industry_jobs_completed_manufacture_structure integer */
	IndustryJobsCompletedManufactureStructure int64 `json:"industry_jobs_completed_manufacture_structure,omitempty"`
/*
	 industry_jobs_completed_manufacture_structure_quantity integer */
	IndustryJobsCompletedManufactureStructureQuantity int64 `json:"industry_jobs_completed_manufacture_structure_quantity,omitempty"`
/*
	 industry_jobs_completed_manufacture_subsystem integer */
	IndustryJobsCompletedManufactureSubsystem int64 `json:"industry_jobs_completed_manufacture_subsystem,omitempty"`
/*
	 industry_jobs_completed_manufacture_subsystem_quantity integer */
	IndustryJobsCompletedManufactureSubsystemQuantity int64 `json:"industry_jobs_completed_manufacture_subsystem_quantity,omitempty"`
/*
	 industry_jobs_completed_material_productivity integer */
	IndustryJobsCompletedMaterialProductivity int64 `json:"industry_jobs_completed_material_productivity,omitempty"`
/*
	 industry_jobs_completed_time_productivity integer */
	IndustryJobsCompletedTimeProductivity int64 `json:"industry_jobs_completed_time_productivity,omitempty"`
/*
	 industry_jobs_started_copy_blueprint integer */
	IndustryJobsStartedCopyBlueprint int64 `json:"industry_jobs_started_copy_blueprint,omitempty"`
/*
	 industry_jobs_started_invention integer */
	IndustryJobsStartedInvention int64 `json:"industry_jobs_started_invention,omitempty"`
/*
	 industry_jobs_started_manufacture integer */
	IndustryJobsStartedManufacture int64 `json:"industry_jobs_started_manufacture,omitempty"`
/*
	 industry_jobs_started_material_productivity integer */
	IndustryJobsStartedMaterialProductivity int64 `json:"industry_jobs_started_material_productivity,omitempty"`
/*
	 industry_jobs_started_time_productivity integer */
	IndustryJobsStartedTimeProductivity int64 `json:"industry_jobs_started_time_productivity,omitempty"`
/*
	 industry_reprocess_item integer */
	IndustryReprocessItem int64 `json:"industry_reprocess_item,omitempty"`
/*
	 industry_reprocess_item_quantity integer */
	IndustryReprocessItemQuantity int64 `json:"industry_reprocess_item_quantity,omitempty"`
/*
	 inventory_abandon_loot_quantity integer */
	InventoryAbandonLootQuantity int64 `json:"inventory_abandon_loot_quantity,omitempty"`
/*
	 inventory_trash_item_quantity integer */
	InventoryTrashItemQuantity int64 `json:"inventory_trash_item_quantity,omitempty"`
/*
	 isk_in integer */
	IskIn int64 `json:"isk_in,omitempty"`
/*
	 isk_out integer */
	IskOut int64 `json:"isk_out,omitempty"`
/*
	 market_accept_contracts_courier integer */
	MarketAcceptContractsCourier int64 `json:"market_accept_contracts_courier,omitempty"`
/*
	 market_accept_contracts_item_exchange integer */
	MarketAcceptContractsItemExchange int64 `json:"market_accept_contracts_item_exchange,omitempty"`
/*
	 market_buy_orders_placed integer */
	MarketBuyOrdersPlaced int64 `json:"market_buy_orders_placed,omitempty"`
/*
	 market_cancel_market_order integer */
	MarketCancelMarketOrder int64 `json:"market_cancel_market_order,omitempty"`
/*
	 market_create_contracts_auction integer */
	MarketCreateContractsAuction int64 `json:"market_create_contracts_auction,omitempty"`
/*
	 market_create_contracts_courier integer */
	MarketCreateContractsCourier int64 `json:"market_create_contracts_courier,omitempty"`
/*
	 market_create_contracts_item_exchange integer */
	MarketCreateContractsItemExchange int64 `json:"market_create_contracts_item_exchange,omitempty"`
/*
	 market_deliver_courier_contract integer */
	MarketDeliverCourierContract int64 `json:"market_deliver_courier_contract,omitempty"`
/*
	 market_isk_gained integer */
	MarketIskGained int64 `json:"market_isk_gained,omitempty"`
/*
	 market_isk_spent integer */
	MarketIskSpent int64 `json:"market_isk_spent,omitempty"`
/*
	 market_modify_market_order integer */
	MarketModifyMarketOrder int64 `json:"market_modify_market_order,omitempty"`
/*
	 market_search_contracts integer */
	MarketSearchContracts int64 `json:"market_search_contracts,omitempty"`
/*
	 market_sell_orders_placed integer */
	MarketSellOrdersPlaced int64 `json:"market_sell_orders_placed,omitempty"`
/*
	 mining_drone_mine integer */
	MiningDroneMine int64 `json:"mining_drone_mine,omitempty"`
/*
	 mining_ore_arkonor integer */
	MiningOreArkonor int64 `json:"mining_ore_arkonor,omitempty"`
/*
	 mining_ore_bistot integer */
	MiningOreBistot int64 `json:"mining_ore_bistot,omitempty"`
/*
	 mining_ore_crokite integer */
	MiningOreCrokite int64 `json:"mining_ore_crokite,omitempty"`
/*
	 mining_ore_dark_ochre integer */
	MiningOreDarkOchre int64 `json:"mining_ore_dark_ochre,omitempty"`
/*
	 mining_ore_gneiss integer */
	MiningOreGneiss int64 `json:"mining_ore_gneiss,omitempty"`
/*
	 mining_ore_harvestable_cloud integer */
	MiningOreHarvestableCloud int64 `json:"mining_ore_harvestable_cloud,omitempty"`
/*
	 mining_ore_hedbergite integer */
	MiningOreHedbergite int64 `json:"mining_ore_hedbergite,omitempty"`
/*
	 mining_ore_hemorphite integer */
	MiningOreHemorphite int64 `json:"mining_ore_hemorphite,omitempty"`
/*
	 mining_ore_ice integer */
	MiningOreIce int64 `json:"mining_ore_ice,omitempty"`
/*
	 mining_ore_jaspet integer */
	MiningOreJaspet int64 `json:"mining_ore_jaspet,omitempty"`
/*
	 mining_ore_kernite integer */
	MiningOreKernite int64 `json:"mining_ore_kernite,omitempty"`
/*
	 mining_ore_mercoxit integer */
	MiningOreMercoxit int64 `json:"mining_ore_mercoxit,omitempty"`
/*
	 mining_ore_omber integer */
	MiningOreOmber int64 `json:"mining_ore_omber,omitempty"`
/*
	 mining_ore_plagioclase integer */
	MiningOrePlagioclase int64 `json:"mining_ore_plagioclase,omitempty"`
/*
	 mining_ore_pyroxeres integer */
	MiningOrePyroxeres int64 `json:"mining_ore_pyroxeres,omitempty"`
/*
	 mining_ore_scordite integer */
	MiningOreScordite int64 `json:"mining_ore_scordite,omitempty"`
/*
	 mining_ore_spodumain integer */
	MiningOreSpodumain int64 `json:"mining_ore_spodumain,omitempty"`
/*
	 mining_ore_veldspar integer */
	MiningOreVeldspar int64 `json:"mining_ore_veldspar,omitempty"`
/*
	 module_activations_armor_hardener integer */
	ModuleActivationsArmorHardener int64 `json:"module_activations_armor_hardener,omitempty"`
/*
	 module_activations_armor_repair_unit integer */
	ModuleActivationsArmorRepairUnit int64 `json:"module_activations_armor_repair_unit,omitempty"`
/*
	 module_activations_armor_resistance_shift_hardener integer */
	ModuleActivationsArmorResistanceShiftHardener int64 `json:"module_activations_armor_resistance_shift_hardener,omitempty"`
/*
	 module_activations_automated_targeting_system integer */
	ModuleActivationsAutomatedTargetingSystem int64 `json:"module_activations_automated_targeting_system,omitempty"`
/*
	 module_activations_bastion integer */
	ModuleActivationsBastion int64 `json:"module_activations_bastion,omitempty"`
/*
	 module_activations_bomb_launcher integer */
	ModuleActivationsBombLauncher int64 `json:"module_activations_bomb_launcher,omitempty"`
/*
	 module_activations_capacitor_booster integer */
	ModuleActivationsCapacitorBooster int64 `json:"module_activations_capacitor_booster,omitempty"`
/*
	 module_activations_cargo_scanner integer */
	ModuleActivationsCargoScanner int64 `json:"module_activations_cargo_scanner,omitempty"`
/*
	 module_activations_cloaking_device integer */
	ModuleActivationsCloakingDevice int64 `json:"module_activations_cloaking_device,omitempty"`
/*
	 module_activations_clone_vat_bay integer */
	ModuleActivationsCloneVatBay int64 `json:"module_activations_clone_vat_bay,omitempty"`
/*
	 module_activations_cynosural_field integer */
	ModuleActivationsCynosuralField int64 `json:"module_activations_cynosural_field,omitempty"`
/*
	 module_activations_damage_control integer */
	ModuleActivationsDamageControl int64 `json:"module_activations_damage_control,omitempty"`
/*
	 module_activations_data_miners integer */
	ModuleActivationsDataMiners int64 `json:"module_activations_data_miners,omitempty"`
/*
	 module_activations_drone_control_unit integer */
	ModuleActivationsDroneControlUnit int64 `json:"module_activations_drone_control_unit,omitempty"`
/*
	 module_activations_drone_tracking_modules integer */
	ModuleActivationsDroneTrackingModules int64 `json:"module_activations_drone_tracking_modules,omitempty"`
/*
	 module_activations_eccm integer */
	ModuleActivationsEccm int64 `json:"module_activations_eccm,omitempty"`
/*
	 module_activations_ecm integer */
	ModuleActivationsEcm int64 `json:"module_activations_ecm,omitempty"`
/*
	 module_activations_ecm_burst integer */
	ModuleActivationsEcmBurst int64 `json:"module_activations_ecm_burst,omitempty"`
/*
	 module_activations_energy_destabilizer integer */
	ModuleActivationsEnergyDestabilizer int64 `json:"module_activations_energy_destabilizer,omitempty"`
/*
	 module_activations_energy_vampire integer */
	ModuleActivationsEnergyVampire int64 `json:"module_activations_energy_vampire,omitempty"`
/*
	 module_activations_energy_weapon integer */
	ModuleActivationsEnergyWeapon int64 `json:"module_activations_energy_weapon,omitempty"`
/*
	 module_activations_festival_launcher integer */
	ModuleActivationsFestivalLauncher int64 `json:"module_activations_festival_launcher,omitempty"`
/*
	 module_activations_frequency_mining_laser integer */
	ModuleActivationsFrequencyMiningLaser int64 `json:"module_activations_frequency_mining_laser,omitempty"`
/*
	 module_activations_fueled_armor_repairer integer */
	ModuleActivationsFueledArmorRepairer int64 `json:"module_activations_fueled_armor_repairer,omitempty"`
/*
	 module_activations_fueled_shield_booster integer */
	ModuleActivationsFueledShieldBooster int64 `json:"module_activations_fueled_shield_booster,omitempty"`
/*
	 module_activations_gang_coordinator integer */
	ModuleActivationsGangCoordinator int64 `json:"module_activations_gang_coordinator,omitempty"`
/*
	 module_activations_gas_cloud_harvester integer */
	ModuleActivationsGasCloudHarvester int64 `json:"module_activations_gas_cloud_harvester,omitempty"`
/*
	 module_activations_hull_repair_unit integer */
	ModuleActivationsHullRepairUnit int64 `json:"module_activations_hull_repair_unit,omitempty"`
/*
	 module_activations_hybrid_weapon integer */
	ModuleActivationsHybridWeapon int64 `json:"module_activations_hybrid_weapon,omitempty"`
/*
	 module_activations_industrial_core integer */
	ModuleActivationsIndustrialCore int64 `json:"module_activations_industrial_core,omitempty"`
/*
	 module_activations_interdiction_sphere_launcher integer */
	ModuleActivationsInterdictionSphereLauncher int64 `json:"module_activations_interdiction_sphere_launcher,omitempty"`
/*
	 module_activations_micro_jump_drive integer */
	ModuleActivationsMicroJumpDrive int64 `json:"module_activations_micro_jump_drive,omitempty"`
/*
	 module_activations_mining_laser integer */
	ModuleActivationsMiningLaser int64 `json:"module_activations_mining_laser,omitempty"`
/*
	 module_activations_missile_launcher integer */
	ModuleActivationsMissileLauncher int64 `json:"module_activations_missile_launcher,omitempty"`
/*
	 module_activations_passive_targeting_system integer */
	ModuleActivationsPassiveTargetingSystem int64 `json:"module_activations_passive_targeting_system,omitempty"`
/*
	 module_activations_probe_launcher integer */
	ModuleActivationsProbeLauncher int64 `json:"module_activations_probe_launcher,omitempty"`
/*
	 module_activations_projected_eccm integer */
	ModuleActivationsProjectedEccm int64 `json:"module_activations_projected_eccm,omitempty"`
/*
	 module_activations_projectile_weapon integer */
	ModuleActivationsProjectileWeapon int64 `json:"module_activations_projectile_weapon,omitempty"`
/*
	 module_activations_propulsion_module integer */
	ModuleActivationsPropulsionModule int64 `json:"module_activations_propulsion_module,omitempty"`
/*
	 module_activations_remote_armor_repairer integer */
	ModuleActivationsRemoteArmorRepairer int64 `json:"module_activations_remote_armor_repairer,omitempty"`
/*
	 module_activations_remote_capacitor_transmitter integer */
	ModuleActivationsRemoteCapacitorTransmitter int64 `json:"module_activations_remote_capacitor_transmitter,omitempty"`
/*
	 module_activations_remote_ecm_burst integer */
	ModuleActivationsRemoteEcmBurst int64 `json:"module_activations_remote_ecm_burst,omitempty"`
/*
	 module_activations_remote_hull_repairer integer */
	ModuleActivationsRemoteHullRepairer int64 `json:"module_activations_remote_hull_repairer,omitempty"`
/*
	 module_activations_remote_sensor_booster integer */
	ModuleActivationsRemoteSensorBooster int64 `json:"module_activations_remote_sensor_booster,omitempty"`
/*
	 module_activations_remote_sensor_damper integer */
	ModuleActivationsRemoteSensorDamper int64 `json:"module_activations_remote_sensor_damper,omitempty"`
/*
	 module_activations_remote_shield_booster integer */
	ModuleActivationsRemoteShieldBooster int64 `json:"module_activations_remote_shield_booster,omitempty"`
/*
	 module_activations_remote_tracking_computer integer */
	ModuleActivationsRemoteTrackingComputer int64 `json:"module_activations_remote_tracking_computer,omitempty"`
/*
	 module_activations_salvager integer */
	ModuleActivationsSalvager int64 `json:"module_activations_salvager,omitempty"`
/*
	 module_activations_sensor_booster integer */
	ModuleActivationsSensorBooster int64 `json:"module_activations_sensor_booster,omitempty"`
/*
	 module_activations_shield_booster integer */
	ModuleActivationsShieldBooster int64 `json:"module_activations_shield_booster,omitempty"`
/*
	 module_activations_shield_hardener integer */
	ModuleActivationsShieldHardener int64 `json:"module_activations_shield_hardener,omitempty"`
/*
	 module_activations_ship_scanner integer */
	ModuleActivationsShipScanner int64 `json:"module_activations_ship_scanner,omitempty"`
/*
	 module_activations_siege integer */
	ModuleActivationsSiege int64 `json:"module_activations_siege,omitempty"`
/*
	 module_activations_smart_bomb integer */
	ModuleActivationsSmartBomb int64 `json:"module_activations_smart_bomb,omitempty"`
/*
	 module_activations_stasis_web integer */
	ModuleActivationsStasisWeb int64 `json:"module_activations_stasis_web,omitempty"`
/*
	 module_activations_strip_miner integer */
	ModuleActivationsStripMiner int64 `json:"module_activations_strip_miner,omitempty"`
/*
	 module_activations_super_weapon integer */
	ModuleActivationsSuperWeapon int64 `json:"module_activations_super_weapon,omitempty"`
/*
	 module_activations_survey_scanner integer */
	ModuleActivationsSurveyScanner int64 `json:"module_activations_survey_scanner,omitempty"`
/*
	 module_activations_target_breaker integer */
	ModuleActivationsTargetBreaker int64 `json:"module_activations_target_breaker,omitempty"`
/*
	 module_activations_target_painter integer */
	ModuleActivationsTargetPainter int64 `json:"module_activations_target_painter,omitempty"`
/*
	 module_activations_tracking_computer integer */
	ModuleActivationsTrackingComputer int64 `json:"module_activations_tracking_computer,omitempty"`
/*
	 module_activations_tracking_disruptor integer */
	ModuleActivationsTrackingDisruptor int64 `json:"module_activations_tracking_disruptor,omitempty"`
/*
	 module_activations_tractor_beam integer */
	ModuleActivationsTractorBeam int64 `json:"module_activations_tractor_beam,omitempty"`
/*
	 module_activations_triage integer */
	ModuleActivationsTriage int64 `json:"module_activations_triage,omitempty"`
/*
	 module_activations_warp_disrupt_field_generator integer */
	ModuleActivationsWarpDisruptFieldGenerator int64 `json:"module_activations_warp_disrupt_field_generator,omitempty"`
/*
	 module_activations_warp_scrambler integer */
	ModuleActivationsWarpScrambler int64 `json:"module_activations_warp_scrambler,omitempty"`
/*
	 module_link_weapons integer */
	ModuleLinkWeapons int64 `json:"module_link_weapons,omitempty"`
/*
	 module_overload integer */
	ModuleOverload int64 `json:"module_overload,omitempty"`
/*
	 module_repairs integer */
	ModuleRepairs int64 `json:"module_repairs,omitempty"`
/*
	 orbital_strike_characters_killed integer */
	OrbitalStrikeCharactersKilled int64 `json:"orbital_strike_characters_killed,omitempty"`
/*
	 orbital_strike_damage_to_players_armor_amount integer */
	OrbitalStrikeDamageToPlayersArmorAmount int64 `json:"orbital_strike_damage_to_players_armor_amount,omitempty"`
/*
	 orbital_strike_damage_to_players_shield_amount integer */
	OrbitalStrikeDamageToPlayersShieldAmount int64 `json:"orbital_strike_damage_to_players_shield_amount,omitempty"`
/*
	 pve_dungeons_completed_agent integer */
	PveDungeonsCompletedAgent int64 `json:"pve_dungeons_completed_agent,omitempty"`
/*
	 pve_dungeons_completed_distribution integer */
	PveDungeonsCompletedDistribution int64 `json:"pve_dungeons_completed_distribution,omitempty"`
/*
	 pve_missions_succeeded integer */
	PveMissionsSucceeded int64 `json:"pve_missions_succeeded,omitempty"`
/*
	 pve_missions_succeeded_epic_arc integer */
	PveMissionsSucceededEpicArc int64 `json:"pve_missions_succeeded_epic_arc,omitempty"`
/*
	 social_add_contact_bad integer */
	SocialAddContactBad int64 `json:"social_add_contact_bad,omitempty"`
/*
	 social_add_contact_good integer */
	SocialAddContactGood int64 `json:"social_add_contact_good,omitempty"`
/*
	 social_add_contact_high integer */
	SocialAddContactHigh int64 `json:"social_add_contact_high,omitempty"`
/*
	 social_add_contact_horrible integer */
	SocialAddContactHorrible int64 `json:"social_add_contact_horrible,omitempty"`
/*
	 social_add_contact_neutral integer */
	SocialAddContactNeutral int64 `json:"social_add_contact_neutral,omitempty"`
/*
	 social_add_note integer */
	SocialAddNote int64 `json:"social_add_note,omitempty"`
/*
	 social_added_as_contact_bad integer */
	SocialAddedAsContactBad int64 `json:"social_added_as_contact_bad,omitempty"`
/*
	 social_added_as_contact_good integer */
	SocialAddedAsContactGood int64 `json:"social_added_as_contact_good,omitempty"`
/*
	 social_added_as_contact_high integer */
	SocialAddedAsContactHigh int64 `json:"social_added_as_contact_high,omitempty"`
/*
	 social_added_as_contact_horrible integer */
	SocialAddedAsContactHorrible int64 `json:"social_added_as_contact_horrible,omitempty"`
/*
	 social_added_as_contact_neutral integer */
	SocialAddedAsContactNeutral int64 `json:"social_added_as_contact_neutral,omitempty"`
/*
	 social_calendar_event_created integer */
	SocialCalendarEventCreated int64 `json:"social_calendar_event_created,omitempty"`
/*
	 social_chat_messages_alliance integer */
	SocialChatMessagesAlliance int64 `json:"social_chat_messages_alliance,omitempty"`
/*
	 social_chat_messages_constellation integer */
	SocialChatMessagesConstellation int64 `json:"social_chat_messages_constellation,omitempty"`
/*
	 social_chat_messages_corporation integer */
	SocialChatMessagesCorporation int64 `json:"social_chat_messages_corporation,omitempty"`
/*
	 social_chat_messages_fleet integer */
	SocialChatMessagesFleet int64 `json:"social_chat_messages_fleet,omitempty"`
/*
	 social_chat_messages_region integer */
	SocialChatMessagesRegion int64 `json:"social_chat_messages_region,omitempty"`
/*
	 social_chat_messages_solarsystem integer */
	SocialChatMessagesSolarsystem int64 `json:"social_chat_messages_solarsystem,omitempty"`
/*
	 social_chat_messages_warfaction integer */
	SocialChatMessagesWarfaction int64 `json:"social_chat_messages_warfaction,omitempty"`
/*
	 social_chat_total_message_length integer */
	SocialChatTotalMessageLength int64 `json:"social_chat_total_message_length,omitempty"`
/*
	 social_direct_trades integer */
	SocialDirectTrades int64 `json:"social_direct_trades,omitempty"`
/*
	 social_fleet_broadcasts integer */
	SocialFleetBroadcasts int64 `json:"social_fleet_broadcasts,omitempty"`
/*
	 social_fleet_joins integer */
	SocialFleetJoins int64 `json:"social_fleet_joins,omitempty"`
/*
	 social_mails_received integer */
	SocialMailsReceived int64 `json:"social_mails_received,omitempty"`
/*
	 social_mails_sent integer */
	SocialMailsSent int64 `json:"social_mails_sent,omitempty"`
/*
	 travel_acceleration_gate_activations integer */
	TravelAccelerationGateActivations int64 `json:"travel_acceleration_gate_activations,omitempty"`
/*
	 travel_align_to integer */
	TravelAlignTo int64 `json:"travel_align_to,omitempty"`
/*
	 travel_distance_warped_high_sec integer */
	TravelDistanceWarpedHighSec int64 `json:"travel_distance_warped_high_sec,omitempty"`
/*
	 travel_distance_warped_low_sec integer */
	TravelDistanceWarpedLowSec int64 `json:"travel_distance_warped_low_sec,omitempty"`
/*
	 travel_distance_warped_null_sec integer */
	TravelDistanceWarpedNullSec int64 `json:"travel_distance_warped_null_sec,omitempty"`
/*
	 travel_distance_warped_wormhole integer */
	TravelDistanceWarpedWormhole int64 `json:"travel_distance_warped_wormhole,omitempty"`
/*
	 travel_docks_high_sec integer */
	TravelDocksHighSec int64 `json:"travel_docks_high_sec,omitempty"`
/*
	 travel_docks_low_sec integer */
	TravelDocksLowSec int64 `json:"travel_docks_low_sec,omitempty"`
/*
	 travel_docks_null_sec integer */
	TravelDocksNullSec int64 `json:"travel_docks_null_sec,omitempty"`
/*
	 travel_jumps_stargate_high_sec integer */
	TravelJumpsStargateHighSec int64 `json:"travel_jumps_stargate_high_sec,omitempty"`
/*
	 travel_jumps_stargate_low_sec integer */
	TravelJumpsStargateLowSec int64 `json:"travel_jumps_stargate_low_sec,omitempty"`
/*
	 travel_jumps_stargate_null_sec integer */
	TravelJumpsStargateNullSec int64 `json:"travel_jumps_stargate_null_sec,omitempty"`
/*
	 travel_jumps_wormhole integer */
	TravelJumpsWormhole int64 `json:"travel_jumps_wormhole,omitempty"`
/*
	 travel_warps_high_sec integer */
	TravelWarpsHighSec int64 `json:"travel_warps_high_sec,omitempty"`
/*
	 travel_warps_low_sec integer */
	TravelWarpsLowSec int64 `json:"travel_warps_low_sec,omitempty"`
/*
	 travel_warps_null_sec integer */
	TravelWarpsNullSec int64 `json:"travel_warps_null_sec,omitempty"`
/*
	 travel_warps_to_bookmark integer */
	TravelWarpsToBookmark int64 `json:"travel_warps_to_bookmark,omitempty"`
/*
	 travel_warps_to_celestial integer */
	TravelWarpsToCelestial int64 `json:"travel_warps_to_celestial,omitempty"`
/*
	 travel_warps_to_fleet_member integer */
	TravelWarpsToFleetMember int64 `json:"travel_warps_to_fleet_member,omitempty"`
/*
	 travel_warps_to_scan_result integer */
	TravelWarpsToScanResult int64 `json:"travel_warps_to_scan_result,omitempty"`
/*
	 travel_warps_wormhole integer */
	TravelWarpsWormhole int64 `json:"travel_warps_wormhole,omitempty"`
/*
	 Gregorian year for this set of aggregates */
	Year int32 `json:"year,omitempty"`
}
