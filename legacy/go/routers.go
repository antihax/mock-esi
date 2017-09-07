package esilegacy

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/antihax/mock-esi/mockesi"
)

func init() {

	mockesi.NewRoute(
		"GetAlliances",
		"GET",
		"/legacy/alliances/",
		GetAlliances,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceId",
		"GET",
		"/legacy/alliances/{alliance_id}/",
		GetAlliancesAllianceId,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdCorporations",
		"GET",
		"/legacy/alliances/{alliance_id}/corporations/",
		GetAlliancesAllianceIdCorporations,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdIcons",
		"GET",
		"/legacy/alliances/{alliance_id}/icons/",
		GetAlliancesAllianceIdIcons,
	)

	mockesi.NewRoute(
		"GetAlliancesNames",
		"GET",
		"/legacy/alliances/names/",
		GetAlliancesNames,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAssets",
		"GET",
		"/legacy/characters/{character_id}/assets/",
		GetCharactersCharacterIdAssets,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBookmarks",
		"GET",
		"/legacy/characters/{character_id}/bookmarks/",
		GetCharactersCharacterIdBookmarks,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBookmarksFolders",
		"GET",
		"/legacy/characters/{character_id}/bookmarks/folders/",
		GetCharactersCharacterIdBookmarksFolders,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCalendar",
		"GET",
		"/legacy/characters/{character_id}/calendar/",
		GetCharactersCharacterIdCalendar,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCalendarEventId",
		"GET",
		"/legacy/characters/{character_id}/calendar/{event_id}/",
		GetCharactersCharacterIdCalendarEventId,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdCalendarEventId",
		"PUT",
		"/legacy/characters/{character_id}/calendar/{event_id}/",
		PutCharactersCharacterIdCalendarEventId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterId",
		"GET",
		"/legacy/characters/{character_id}/",
		GetCharactersCharacterId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAgentsResearch",
		"GET",
		"/legacy/characters/{character_id}/agents_research/",
		GetCharactersCharacterIdAgentsResearch,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBlueprints",
		"GET",
		"/legacy/characters/{character_id}/blueprints/",
		GetCharactersCharacterIdBlueprints,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdChatChannels",
		"GET",
		"/legacy/characters/{character_id}/chat_channels/",
		GetCharactersCharacterIdChatChannels,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCorporationhistory",
		"GET",
		"/legacy/characters/{character_id}/corporationhistory/",
		GetCharactersCharacterIdCorporationhistory,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFatigue",
		"GET",
		"/legacy/characters/{character_id}/fatigue/",
		GetCharactersCharacterIdFatigue,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMedals",
		"GET",
		"/legacy/characters/{character_id}/medals/",
		GetCharactersCharacterIdMedals,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdNotificationsContacts",
		"GET",
		"/legacy/characters/{character_id}/notifications/contacts/",
		GetCharactersCharacterIdNotificationsContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPortrait",
		"GET",
		"/legacy/characters/{character_id}/portrait/",
		GetCharactersCharacterIdPortrait,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdRoles",
		"GET",
		"/legacy/characters/{character_id}/roles/",
		GetCharactersCharacterIdRoles,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdStandings",
		"GET",
		"/legacy/characters/{character_id}/standings/",
		GetCharactersCharacterIdStandings,
	)

	mockesi.NewRoute(
		"GetCharactersNames",
		"GET",
		"/legacy/characters/names/",
		GetCharactersNames,
	)

	mockesi.NewRoute(
		"PostCharactersAffiliation",
		"POST",
		"/legacy/characters/affiliation/",
		PostCharactersAffiliation,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdCspa",
		"POST",
		"/legacy/characters/{character_id}/cspa/",
		PostCharactersCharacterIdCspa,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdClones",
		"GET",
		"/legacy/characters/{character_id}/clones/",
		GetCharactersCharacterIdClones,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdImplants",
		"GET",
		"/legacy/characters/{character_id}/implants/",
		GetCharactersCharacterIdImplants,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdContacts",
		"DELETE",
		"/legacy/characters/{character_id}/contacts/",
		DeleteCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContacts",
		"GET",
		"/legacy/characters/{character_id}/contacts/",
		GetCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContactsLabels",
		"GET",
		"/legacy/characters/{character_id}/contacts/labels/",
		GetCharactersCharacterIdContactsLabels,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContacts",
		"GET",
		"/legacy/corporations/{corporation_id}/contacts/",
		GetCorporationsCorporationIdContacts,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdContacts",
		"POST",
		"/legacy/characters/{character_id}/contacts/",
		PostCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdContacts",
		"PUT",
		"/legacy/characters/{character_id}/contacts/",
		PutCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContracts",
		"GET",
		"/legacy/characters/{character_id}/contracts/",
		GetCharactersCharacterIdContracts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContractsContractIdBids",
		"GET",
		"/legacy/characters/{character_id}/contracts/{contract_id}/bids/",
		GetCharactersCharacterIdContractsContractIdBids,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContractsContractIdItems",
		"GET",
		"/legacy/characters/{character_id}/contracts/{contract_id}/items/",
		GetCharactersCharacterIdContractsContractIdItems,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationId",
		"GET",
		"/legacy/corporations/{corporation_id}/",
		GetCorporationsCorporationId,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdAlliancehistory",
		"GET",
		"/legacy/corporations/{corporation_id}/alliancehistory/",
		GetCorporationsCorporationIdAlliancehistory,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdDivisions",
		"GET",
		"/legacy/corporations/{corporation_id}/divisions/",
		GetCorporationsCorporationIdDivisions,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdIcons",
		"GET",
		"/legacy/corporations/{corporation_id}/icons/",
		GetCorporationsCorporationIdIcons,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMembers",
		"GET",
		"/legacy/corporations/{corporation_id}/members/",
		GetCorporationsCorporationIdMembers,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMembersLimit",
		"GET",
		"/legacy/corporations/{corporation_id}/members/limit/",
		GetCorporationsCorporationIdMembersLimit,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMembertracking",
		"GET",
		"/legacy/corporations/{corporation_id}/membertracking/",
		GetCorporationsCorporationIdMembertracking,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdRoles",
		"GET",
		"/legacy/corporations/{corporation_id}/roles/",
		GetCorporationsCorporationIdRoles,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdStructures",
		"GET",
		"/legacy/corporations/{corporation_id}/structures/",
		GetCorporationsCorporationIdStructures,
	)

	mockesi.NewRoute(
		"GetCorporationsNames",
		"GET",
		"/legacy/corporations/names/",
		GetCorporationsNames,
	)

	mockesi.NewRoute(
		"GetCorporationsNpccorps",
		"GET",
		"/legacy/corporations/npccorps/",
		GetCorporationsNpccorps,
	)

	mockesi.NewRoute(
		"PutCorporationsCorporationIdStructuresStructureId",
		"PUT",
		"/legacy/corporations/{corporation_id}/structures/{structure_id}/",
		PutCorporationsCorporationIdStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetDogmaAttributes",
		"GET",
		"/legacy/dogma/attributes/",
		GetDogmaAttributes,
	)

	mockesi.NewRoute(
		"GetDogmaAttributesAttributeId",
		"GET",
		"/legacy/dogma/attributes/{attribute_id}/",
		GetDogmaAttributesAttributeId,
	)

	mockesi.NewRoute(
		"GetDogmaEffects",
		"GET",
		"/legacy/dogma/effects/",
		GetDogmaEffects,
	)

	mockesi.NewRoute(
		"GetDogmaEffectsEffectId",
		"GET",
		"/legacy/dogma/effects/{effect_id}/",
		GetDogmaEffectsEffectId,
	)

	mockesi.NewRoute(
		"GetFwLeaderboards",
		"GET",
		"/legacy/fw/leaderboards/",
		GetFwLeaderboards,
	)

	mockesi.NewRoute(
		"GetFwLeaderboardsCharacters",
		"GET",
		"/legacy/fw/leaderboards/characters/",
		GetFwLeaderboardsCharacters,
	)

	mockesi.NewRoute(
		"GetFwLeaderboardsCorporations",
		"GET",
		"/legacy/fw/leaderboards/corporations/",
		GetFwLeaderboardsCorporations,
	)

	mockesi.NewRoute(
		"GetFwStats",
		"GET",
		"/legacy/fw/stats/",
		GetFwStats,
	)

	mockesi.NewRoute(
		"GetFwSystems",
		"GET",
		"/legacy/fw/systems/",
		GetFwSystems,
	)

	mockesi.NewRoute(
		"GetFwWars",
		"GET",
		"/legacy/fw/wars/",
		GetFwWars,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdFittingsFittingId",
		"DELETE",
		"/legacy/characters/{character_id}/fittings/{fitting_id}/",
		DeleteCharactersCharacterIdFittingsFittingId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFittings",
		"GET",
		"/legacy/characters/{character_id}/fittings/",
		GetCharactersCharacterIdFittings,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdFittings",
		"POST",
		"/legacy/characters/{character_id}/fittings/",
		PostCharactersCharacterIdFittings,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdMembersMemberId",
		"DELETE",
		"/legacy/fleets/{fleet_id}/members/{member_id}/",
		DeleteFleetsFleetIdMembersMemberId,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdSquadsSquadId",
		"DELETE",
		"/legacy/fleets/{fleet_id}/squads/{squad_id}/",
		DeleteFleetsFleetIdSquadsSquadId,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdWingsWingId",
		"DELETE",
		"/legacy/fleets/{fleet_id}/wings/{wing_id}/",
		DeleteFleetsFleetIdWingsWingId,
	)

	mockesi.NewRoute(
		"GetFleetsFleetId",
		"GET",
		"/legacy/fleets/{fleet_id}/",
		GetFleetsFleetId,
	)

	mockesi.NewRoute(
		"GetFleetsFleetIdMembers",
		"GET",
		"/legacy/fleets/{fleet_id}/members/",
		GetFleetsFleetIdMembers,
	)

	mockesi.NewRoute(
		"GetFleetsFleetIdWings",
		"GET",
		"/legacy/fleets/{fleet_id}/wings/",
		GetFleetsFleetIdWings,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdMembers",
		"POST",
		"/legacy/fleets/{fleet_id}/members/",
		PostFleetsFleetIdMembers,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdWings",
		"POST",
		"/legacy/fleets/{fleet_id}/wings/",
		PostFleetsFleetIdWings,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdWingsWingIdSquads",
		"POST",
		"/legacy/fleets/{fleet_id}/wings/{wing_id}/squads/",
		PostFleetsFleetIdWingsWingIdSquads,
	)

	mockesi.NewRoute(
		"PutFleetsFleetId",
		"PUT",
		"/legacy/fleets/{fleet_id}/",
		PutFleetsFleetId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdMembersMemberId",
		"PUT",
		"/legacy/fleets/{fleet_id}/members/{member_id}/",
		PutFleetsFleetIdMembersMemberId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdSquadsSquadId",
		"PUT",
		"/legacy/fleets/{fleet_id}/squads/{squad_id}/",
		PutFleetsFleetIdSquadsSquadId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdWingsWingId",
		"PUT",
		"/legacy/fleets/{fleet_id}/wings/{wing_id}/",
		PutFleetsFleetIdWingsWingId,
	)

	mockesi.NewRoute(
		"GetIncursions",
		"GET",
		"/legacy/incursions/",
		GetIncursions,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdIndustryJobs",
		"GET",
		"/legacy/characters/{character_id}/industry/jobs/",
		GetCharactersCharacterIdIndustryJobs,
	)

	mockesi.NewRoute(
		"GetIndustryFacilities",
		"GET",
		"/legacy/industry/facilities/",
		GetIndustryFacilities,
	)

	mockesi.NewRoute(
		"GetIndustrySystems",
		"GET",
		"/legacy/industry/systems/",
		GetIndustrySystems,
	)

	mockesi.NewRoute(
		"GetInsurancePrices",
		"GET",
		"/legacy/insurance/prices/",
		GetInsurancePrices,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdKillmailsRecent",
		"GET",
		"/legacy/characters/{character_id}/killmails/recent/",
		GetCharactersCharacterIdKillmailsRecent,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdKillmailsRecent",
		"GET",
		"/legacy/corporations/{corporation_id}/killmails/recent/",
		GetCorporationsCorporationIdKillmailsRecent,
	)

	mockesi.NewRoute(
		"GetKillmailsKillmailIdKillmailHash",
		"GET",
		"/legacy/killmails/{killmail_id}/{killmail_hash}/",
		GetKillmailsKillmailIdKillmailHash,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdLocation",
		"GET",
		"/legacy/characters/{character_id}/location/",
		GetCharactersCharacterIdLocation,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOnline",
		"GET",
		"/legacy/characters/{character_id}/online/",
		GetCharactersCharacterIdOnline,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdShip",
		"GET",
		"/legacy/characters/{character_id}/ship/",
		GetCharactersCharacterIdShip,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdLoyaltyPoints",
		"GET",
		"/legacy/characters/{character_id}/loyalty/points/",
		GetCharactersCharacterIdLoyaltyPoints,
	)

	mockesi.NewRoute(
		"GetLoyaltyStoresCorporationIdOffers",
		"GET",
		"/legacy/loyalty/stores/{corporation_id}/offers/",
		GetLoyaltyStoresCorporationIdOffers,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdMailLabelsLabelId",
		"DELETE",
		"/legacy/characters/{character_id}/mail/labels/{label_id}/",
		DeleteCharactersCharacterIdMailLabelsLabelId,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdMailMailId",
		"DELETE",
		"/legacy/characters/{character_id}/mail/{mail_id}/",
		DeleteCharactersCharacterIdMailMailId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMail",
		"GET",
		"/legacy/characters/{character_id}/mail/",
		GetCharactersCharacterIdMail,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailLabels",
		"GET",
		"/legacy/characters/{character_id}/mail/labels/",
		GetCharactersCharacterIdMailLabels,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailLists",
		"GET",
		"/legacy/characters/{character_id}/mail/lists/",
		GetCharactersCharacterIdMailLists,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailMailId",
		"GET",
		"/legacy/characters/{character_id}/mail/{mail_id}/",
		GetCharactersCharacterIdMailMailId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailUnread",
		"GET",
		"/legacy/characters/{character_id}/mail/unread/",
		GetCharactersCharacterIdMailUnread,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdMail",
		"POST",
		"/legacy/characters/{character_id}/mail/",
		PostCharactersCharacterIdMail,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdMailLabels",
		"POST",
		"/legacy/characters/{character_id}/mail/labels/",
		PostCharactersCharacterIdMailLabels,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdMailMailId",
		"PUT",
		"/legacy/characters/{character_id}/mail/{mail_id}/",
		PutCharactersCharacterIdMailMailId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOrders",
		"GET",
		"/legacy/characters/{character_id}/orders/",
		GetCharactersCharacterIdOrders,
	)

	mockesi.NewRoute(
		"GetMarketsGroups",
		"GET",
		"/legacy/markets/groups/",
		GetMarketsGroups,
	)

	mockesi.NewRoute(
		"GetMarketsGroupsMarketGroupId",
		"GET",
		"/legacy/markets/groups/{market_group_id}/",
		GetMarketsGroupsMarketGroupId,
	)

	mockesi.NewRoute(
		"GetMarketsPrices",
		"GET",
		"/legacy/markets/prices/",
		GetMarketsPrices,
	)

	mockesi.NewRoute(
		"GetMarketsRegionIdHistory",
		"GET",
		"/legacy/markets/{region_id}/history/",
		GetMarketsRegionIdHistory,
	)

	mockesi.NewRoute(
		"GetMarketsRegionIdOrders",
		"GET",
		"/legacy/markets/{region_id}/orders/",
		GetMarketsRegionIdOrders,
	)

	mockesi.NewRoute(
		"GetMarketsRegionIdTypes",
		"GET",
		"/legacy/markets/{region_id}/types/",
		GetMarketsRegionIdTypes,
	)

	mockesi.NewRoute(
		"GetMarketsStructuresStructureId",
		"GET",
		"/legacy/markets/structures/{structure_id}/",
		GetMarketsStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOpportunities",
		"GET",
		"/legacy/characters/{character_id}/opportunities/",
		GetCharactersCharacterIdOpportunities,
	)

	mockesi.NewRoute(
		"GetOpportunitiesGroups",
		"GET",
		"/legacy/opportunities/groups/",
		GetOpportunitiesGroups,
	)

	mockesi.NewRoute(
		"GetOpportunitiesGroupsGroupId",
		"GET",
		"/legacy/opportunities/groups/{group_id}/",
		GetOpportunitiesGroupsGroupId,
	)

	mockesi.NewRoute(
		"GetOpportunitiesTasks",
		"GET",
		"/legacy/opportunities/tasks/",
		GetOpportunitiesTasks,
	)

	mockesi.NewRoute(
		"GetOpportunitiesTasksTaskId",
		"GET",
		"/legacy/opportunities/tasks/{task_id}/",
		GetOpportunitiesTasksTaskId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPlanets",
		"GET",
		"/legacy/characters/{character_id}/planets/",
		GetCharactersCharacterIdPlanets,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPlanetsPlanetId",
		"GET",
		"/legacy/characters/{character_id}/planets/{planet_id}/",
		GetCharactersCharacterIdPlanetsPlanetId,
	)

	mockesi.NewRoute(
		"GetUniverseSchematicsSchematicId",
		"GET",
		"/legacy/universe/schematics/{schematic_id}/",
		GetUniverseSchematicsSchematicId,
	)

	mockesi.NewRoute(
		"GetRouteOriginDestination",
		"GET",
		"/legacy/route/{origin}/{destination}/",
		GetRouteOriginDestination,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdSearch",
		"GET",
		"/legacy/characters/{character_id}/search/",
		GetCharactersCharacterIdSearch,
	)

	mockesi.NewRoute(
		"GetSearch",
		"GET",
		"/legacy/search/",
		GetSearch,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAttributes",
		"GET",
		"/legacy/characters/{character_id}/attributes/",
		GetCharactersCharacterIdAttributes,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdSkillqueue",
		"GET",
		"/legacy/characters/{character_id}/skillqueue/",
		GetCharactersCharacterIdSkillqueue,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdSkills",
		"GET",
		"/legacy/characters/{character_id}/skills/",
		GetCharactersCharacterIdSkills,
	)

	mockesi.NewRoute(
		"GetSovereigntyCampaigns",
		"GET",
		"/legacy/sovereignty/campaigns/",
		GetSovereigntyCampaigns,
	)

	mockesi.NewRoute(
		"GetSovereigntyMap",
		"GET",
		"/legacy/sovereignty/map/",
		GetSovereigntyMap,
	)

	mockesi.NewRoute(
		"GetSovereigntyStructures",
		"GET",
		"/legacy/sovereignty/structures/",
		GetSovereigntyStructures,
	)

	mockesi.NewRoute(
		"GetStatus",
		"GET",
		"/legacy/status/",
		GetStatus,
	)

	mockesi.NewRoute(
		"GetUniverseBloodlines",
		"GET",
		"/legacy/universe/bloodlines/",
		GetUniverseBloodlines,
	)

	mockesi.NewRoute(
		"GetUniverseCategories",
		"GET",
		"/legacy/universe/categories/",
		GetUniverseCategories,
	)

	mockesi.NewRoute(
		"GetUniverseCategoriesCategoryId",
		"GET",
		"/legacy/universe/categories/{category_id}/",
		GetUniverseCategoriesCategoryId,
	)

	mockesi.NewRoute(
		"GetUniverseConstellations",
		"GET",
		"/legacy/universe/constellations/",
		GetUniverseConstellations,
	)

	mockesi.NewRoute(
		"GetUniverseConstellationsConstellationId",
		"GET",
		"/legacy/universe/constellations/{constellation_id}/",
		GetUniverseConstellationsConstellationId,
	)

	mockesi.NewRoute(
		"GetUniverseFactions",
		"GET",
		"/legacy/universe/factions/",
		GetUniverseFactions,
	)

	mockesi.NewRoute(
		"GetUniverseGraphics",
		"GET",
		"/legacy/universe/graphics/",
		GetUniverseGraphics,
	)

	mockesi.NewRoute(
		"GetUniverseGraphicsGraphicId",
		"GET",
		"/legacy/universe/graphics/{graphic_id}/",
		GetUniverseGraphicsGraphicId,
	)

	mockesi.NewRoute(
		"GetUniverseGroups",
		"GET",
		"/legacy/universe/groups/",
		GetUniverseGroups,
	)

	mockesi.NewRoute(
		"GetUniverseGroupsGroupId",
		"GET",
		"/legacy/universe/groups/{group_id}/",
		GetUniverseGroupsGroupId,
	)

	mockesi.NewRoute(
		"GetUniverseMoonsMoonId",
		"GET",
		"/legacy/universe/moons/{moon_id}/",
		GetUniverseMoonsMoonId,
	)

	mockesi.NewRoute(
		"GetUniversePlanetsPlanetId",
		"GET",
		"/legacy/universe/planets/{planet_id}/",
		GetUniversePlanetsPlanetId,
	)

	mockesi.NewRoute(
		"GetUniverseRaces",
		"GET",
		"/legacy/universe/races/",
		GetUniverseRaces,
	)

	mockesi.NewRoute(
		"GetUniverseRegions",
		"GET",
		"/legacy/universe/regions/",
		GetUniverseRegions,
	)

	mockesi.NewRoute(
		"GetUniverseRegionsRegionId",
		"GET",
		"/legacy/universe/regions/{region_id}/",
		GetUniverseRegionsRegionId,
	)

	mockesi.NewRoute(
		"GetUniverseStargatesStargateId",
		"GET",
		"/legacy/universe/stargates/{stargate_id}/",
		GetUniverseStargatesStargateId,
	)

	mockesi.NewRoute(
		"GetUniverseStarsStarId",
		"GET",
		"/legacy/universe/stars/{star_id}/",
		GetUniverseStarsStarId,
	)

	mockesi.NewRoute(
		"GetUniverseStationsStationId",
		"GET",
		"/legacy/universe/stations/{station_id}/",
		GetUniverseStationsStationId,
	)

	mockesi.NewRoute(
		"GetUniverseStructures",
		"GET",
		"/legacy/universe/structures/",
		GetUniverseStructures,
	)

	mockesi.NewRoute(
		"GetUniverseStructuresStructureId",
		"GET",
		"/legacy/universe/structures/{structure_id}/",
		GetUniverseStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetUniverseSystemJumps",
		"GET",
		"/legacy/universe/system_jumps/",
		GetUniverseSystemJumps,
	)

	mockesi.NewRoute(
		"GetUniverseSystemKills",
		"GET",
		"/legacy/universe/system_kills/",
		GetUniverseSystemKills,
	)

	mockesi.NewRoute(
		"GetUniverseSystems",
		"GET",
		"/legacy/universe/systems/",
		GetUniverseSystems,
	)

	mockesi.NewRoute(
		"GetUniverseSystemsSystemId",
		"GET",
		"/legacy/universe/systems/{system_id}/",
		GetUniverseSystemsSystemId,
	)

	mockesi.NewRoute(
		"GetUniverseTypes",
		"GET",
		"/legacy/universe/types/",
		GetUniverseTypes,
	)

	mockesi.NewRoute(
		"GetUniverseTypesTypeId",
		"GET",
		"/legacy/universe/types/{type_id}/",
		GetUniverseTypesTypeId,
	)

	mockesi.NewRoute(
		"PostUniverseNames",
		"POST",
		"/legacy/universe/names/",
		PostUniverseNames,
	)

	mockesi.NewRoute(
		"PostUiAutopilotWaypoint",
		"POST",
		"/legacy/ui/autopilot/waypoint/",
		PostUiAutopilotWaypoint,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowContract",
		"POST",
		"/legacy/ui/openwindow/contract/",
		PostUiOpenwindowContract,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowInformation",
		"POST",
		"/legacy/ui/openwindow/information/",
		PostUiOpenwindowInformation,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowMarketdetails",
		"POST",
		"/legacy/ui/openwindow/marketdetails/",
		PostUiOpenwindowMarketdetails,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowNewmail",
		"POST",
		"/legacy/ui/openwindow/newmail/",
		PostUiOpenwindowNewmail,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdWallet",
		"GET",
		"/legacy/characters/{character_id}/wallet/",
		GetCharactersCharacterIdWallet,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdWalletJournal",
		"GET",
		"/legacy/characters/{character_id}/wallet/journal/",
		GetCharactersCharacterIdWalletJournal,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdWalletTransactions",
		"GET",
		"/legacy/characters/{character_id}/wallet/transactions/",
		GetCharactersCharacterIdWalletTransactions,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdWallets",
		"GET",
		"/legacy/corporations/{corporation_id}/wallets/",
		GetCorporationsCorporationIdWallets,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdWalletsDivisionJournal",
		"GET",
		"/legacy/corporations/{corporation_id}/wallets/{division}/journal/",
		GetCorporationsCorporationIdWalletsDivisionJournal,
	)

	mockesi.NewRoute(
		"GetWars",
		"GET",
		"/legacy/wars/",
		GetWars,
	)

	mockesi.NewRoute(
		"GetWarsWarId",
		"GET",
		"/legacy/wars/{war_id}/",
		GetWarsWarId,
	)

	mockesi.NewRoute(
		"GetWarsWarIdKillmails",
		"GET",
		"/legacy/wars/{war_id}/killmails/",
		GetWarsWarIdKillmails,
	)

}

func errorOut(w http.ResponseWriter, r *http.Request, e error) {
	w.WriteHeader(http.StatusInternalServerError)

	w.Write([]byte(e.Error()))
}

func processParameters(data interface{}, input string) (v interface{}, err error) {
	switch reflect.TypeOf(data).String() {
	case "int":
		{
			v, err = strconv.Atoi(input)
			v = v.(int)
		}
	case "int8":
		{
			v, err = strconv.ParseInt(input, 10, 8)
			v = (int8)(v.(int64))
		}
	case "int16":
		{
			v, err = strconv.ParseInt(input, 10, 16)
			v = (int16)(v.(int64))
		}
	case "rune":
		{
			v, err = strconv.ParseInt(input, 10, 32)
			v = (rune)(v.(int64))
		}
	case "int32":
		{
			v, err = strconv.ParseInt(input, 10, 32)
			v = (int32)(v.(int64))
		}
	case "int64":
		{
			v, err = strconv.ParseInt(input, 10, 64)
			v = v.(int64)
		}
	case "uint8":
		{
			v, err = strconv.ParseUint(input, 10, 8)
			v = (uint8)(v.(uint64))
		}
	case "byte":
		{
			v, err = strconv.ParseUint(input, 10, 8)
			v = (byte)(v.(uint64))
		}
	case "uint16":
		{
			v, err = strconv.ParseUint(input, 10, 16)
			v = (uint16)(v.(uint64))
		}
	case "uint32":
		{
			v, err = strconv.ParseUint(input, 10, 32)
			v = (uint32)(v.(uint64))
		}
	case "uint64":
		{
			v, err = strconv.ParseUint(input, 10, 64)
			v = v.(uint64)
		}
	case "float32":
		{
			v, err = strconv.ParseFloat(input, 32)
			v = (float32)(v.(float64))
		}
	case "float64":
		{
			v, err = strconv.ParseFloat(input, 64)
			v = v.(float64)
		}
	case "bool":
		{
			v, err = strconv.ParseBool(input)
			v = v.(bool)
		}
	case "string":
		{
			v = input
		}
	}

	return
}
