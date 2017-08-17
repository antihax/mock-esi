package esilatest

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
		"/latest/alliances/",
		GetAlliances,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceId",
		"GET",
		"/latest/alliances/{alliance_id}/",
		GetAlliancesAllianceId,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdCorporations",
		"GET",
		"/latest/alliances/{alliance_id}/corporations/",
		GetAlliancesAllianceIdCorporations,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdIcons",
		"GET",
		"/latest/alliances/{alliance_id}/icons/",
		GetAlliancesAllianceIdIcons,
	)

	mockesi.NewRoute(
		"GetAlliancesNames",
		"GET",
		"/latest/alliances/names/",
		GetAlliancesNames,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAssets",
		"GET",
		"/latest/characters/{character_id}/assets/",
		GetCharactersCharacterIdAssets,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBookmarks",
		"GET",
		"/latest/characters/{character_id}/bookmarks/",
		GetCharactersCharacterIdBookmarks,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBookmarksFolders",
		"GET",
		"/latest/characters/{character_id}/bookmarks/folders/",
		GetCharactersCharacterIdBookmarksFolders,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCalendar",
		"GET",
		"/latest/characters/{character_id}/calendar/",
		GetCharactersCharacterIdCalendar,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCalendarEventId",
		"GET",
		"/latest/characters/{character_id}/calendar/{event_id}/",
		GetCharactersCharacterIdCalendarEventId,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdCalendarEventId",
		"PUT",
		"/latest/characters/{character_id}/calendar/{event_id}/",
		PutCharactersCharacterIdCalendarEventId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterId",
		"GET",
		"/latest/characters/{character_id}/",
		GetCharactersCharacterId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAgentsResearch",
		"GET",
		"/latest/characters/{character_id}/agents_research/",
		GetCharactersCharacterIdAgentsResearch,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBlueprints",
		"GET",
		"/latest/characters/{character_id}/blueprints/",
		GetCharactersCharacterIdBlueprints,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdChatChannels",
		"GET",
		"/latest/characters/{character_id}/chat_channels/",
		GetCharactersCharacterIdChatChannels,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCorporationhistory",
		"GET",
		"/latest/characters/{character_id}/corporationhistory/",
		GetCharactersCharacterIdCorporationhistory,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFatigue",
		"GET",
		"/latest/characters/{character_id}/fatigue/",
		GetCharactersCharacterIdFatigue,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMedals",
		"GET",
		"/latest/characters/{character_id}/medals/",
		GetCharactersCharacterIdMedals,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPortrait",
		"GET",
		"/latest/characters/{character_id}/portrait/",
		GetCharactersCharacterIdPortrait,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdRoles",
		"GET",
		"/latest/characters/{character_id}/roles/",
		GetCharactersCharacterIdRoles,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdStandings",
		"GET",
		"/latest/characters/{character_id}/standings/",
		GetCharactersCharacterIdStandings,
	)

	mockesi.NewRoute(
		"GetCharactersNames",
		"GET",
		"/latest/characters/names/",
		GetCharactersNames,
	)

	mockesi.NewRoute(
		"PostCharactersAffiliation",
		"POST",
		"/latest/characters/affiliation/",
		PostCharactersAffiliation,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdCspa",
		"POST",
		"/latest/characters/{character_id}/cspa/",
		PostCharactersCharacterIdCspa,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdClones",
		"GET",
		"/latest/characters/{character_id}/clones/",
		GetCharactersCharacterIdClones,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdImplants",
		"GET",
		"/latest/characters/{character_id}/implants/",
		GetCharactersCharacterIdImplants,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdContacts",
		"DELETE",
		"/latest/characters/{character_id}/contacts/",
		DeleteCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContacts",
		"GET",
		"/latest/characters/{character_id}/contacts/",
		GetCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContactsLabels",
		"GET",
		"/latest/characters/{character_id}/contacts/labels/",
		GetCharactersCharacterIdContactsLabels,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdContacts",
		"POST",
		"/latest/characters/{character_id}/contacts/",
		PostCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdContacts",
		"PUT",
		"/latest/characters/{character_id}/contacts/",
		PutCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContracts",
		"GET",
		"/latest/characters/{character_id}/contracts/",
		GetCharactersCharacterIdContracts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContractsContractIdBids",
		"GET",
		"/latest/characters/{character_id}/contracts/{contract_id}/bids/",
		GetCharactersCharacterIdContractsContractIdBids,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContractsContractIdItems",
		"GET",
		"/latest/characters/{character_id}/contracts/{contract_id}/items/",
		GetCharactersCharacterIdContractsContractIdItems,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationId",
		"GET",
		"/latest/corporations/{corporation_id}/",
		GetCorporationsCorporationId,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdAlliancehistory",
		"GET",
		"/latest/corporations/{corporation_id}/alliancehistory/",
		GetCorporationsCorporationIdAlliancehistory,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdIcons",
		"GET",
		"/latest/corporations/{corporation_id}/icons/",
		GetCorporationsCorporationIdIcons,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMembers",
		"GET",
		"/latest/corporations/{corporation_id}/members/",
		GetCorporationsCorporationIdMembers,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdRoles",
		"GET",
		"/latest/corporations/{corporation_id}/roles/",
		GetCorporationsCorporationIdRoles,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdStructures",
		"GET",
		"/latest/corporations/{corporation_id}/structures/",
		GetCorporationsCorporationIdStructures,
	)

	mockesi.NewRoute(
		"GetCorporationsNames",
		"GET",
		"/latest/corporations/names/",
		GetCorporationsNames,
	)

	mockesi.NewRoute(
		"GetCorporationsNpccorps",
		"GET",
		"/latest/corporations/npccorps/",
		GetCorporationsNpccorps,
	)

	mockesi.NewRoute(
		"PutCorporationsCorporationIdStructuresStructureId",
		"PUT",
		"/latest/corporations/{corporation_id}/structures/{structure_id}/",
		PutCorporationsCorporationIdStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetDogmaAttributes",
		"GET",
		"/latest/dogma/attributes/",
		GetDogmaAttributes,
	)

	mockesi.NewRoute(
		"GetDogmaAttributesAttributeId",
		"GET",
		"/latest/dogma/attributes/{attribute_id}/",
		GetDogmaAttributesAttributeId,
	)

	mockesi.NewRoute(
		"GetDogmaEffects",
		"GET",
		"/latest/dogma/effects/",
		GetDogmaEffects,
	)

	mockesi.NewRoute(
		"GetDogmaEffectsEffectId",
		"GET",
		"/latest/dogma/effects/{effect_id}/",
		GetDogmaEffectsEffectId,
	)

	mockesi.NewRoute(
		"GetFwStats",
		"GET",
		"/latest/fw/stats/",
		GetFwStats,
	)

	mockesi.NewRoute(
		"GetFwWars",
		"GET",
		"/latest/fw/wars/",
		GetFwWars,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdFittingsFittingId",
		"DELETE",
		"/latest/characters/{character_id}/fittings/{fitting_id}/",
		DeleteCharactersCharacterIdFittingsFittingId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFittings",
		"GET",
		"/latest/characters/{character_id}/fittings/",
		GetCharactersCharacterIdFittings,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdFittings",
		"POST",
		"/latest/characters/{character_id}/fittings/",
		PostCharactersCharacterIdFittings,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdMembersMemberId",
		"DELETE",
		"/latest/fleets/{fleet_id}/members/{member_id}/",
		DeleteFleetsFleetIdMembersMemberId,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdSquadsSquadId",
		"DELETE",
		"/latest/fleets/{fleet_id}/squads/{squad_id}/",
		DeleteFleetsFleetIdSquadsSquadId,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdWingsWingId",
		"DELETE",
		"/latest/fleets/{fleet_id}/wings/{wing_id}/",
		DeleteFleetsFleetIdWingsWingId,
	)

	mockesi.NewRoute(
		"GetFleetsFleetId",
		"GET",
		"/latest/fleets/{fleet_id}/",
		GetFleetsFleetId,
	)

	mockesi.NewRoute(
		"GetFleetsFleetIdMembers",
		"GET",
		"/latest/fleets/{fleet_id}/members/",
		GetFleetsFleetIdMembers,
	)

	mockesi.NewRoute(
		"GetFleetsFleetIdWings",
		"GET",
		"/latest/fleets/{fleet_id}/wings/",
		GetFleetsFleetIdWings,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdMembers",
		"POST",
		"/latest/fleets/{fleet_id}/members/",
		PostFleetsFleetIdMembers,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdWings",
		"POST",
		"/latest/fleets/{fleet_id}/wings/",
		PostFleetsFleetIdWings,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdWingsWingIdSquads",
		"POST",
		"/latest/fleets/{fleet_id}/wings/{wing_id}/squads/",
		PostFleetsFleetIdWingsWingIdSquads,
	)

	mockesi.NewRoute(
		"PutFleetsFleetId",
		"PUT",
		"/latest/fleets/{fleet_id}/",
		PutFleetsFleetId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdMembersMemberId",
		"PUT",
		"/latest/fleets/{fleet_id}/members/{member_id}/",
		PutFleetsFleetIdMembersMemberId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdSquadsSquadId",
		"PUT",
		"/latest/fleets/{fleet_id}/squads/{squad_id}/",
		PutFleetsFleetIdSquadsSquadId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdWingsWingId",
		"PUT",
		"/latest/fleets/{fleet_id}/wings/{wing_id}/",
		PutFleetsFleetIdWingsWingId,
	)

	mockesi.NewRoute(
		"GetIncursions",
		"GET",
		"/latest/incursions/",
		GetIncursions,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdIndustryJobs",
		"GET",
		"/latest/characters/{character_id}/industry/jobs/",
		GetCharactersCharacterIdIndustryJobs,
	)

	mockesi.NewRoute(
		"GetIndustryFacilities",
		"GET",
		"/latest/industry/facilities/",
		GetIndustryFacilities,
	)

	mockesi.NewRoute(
		"GetIndustrySystems",
		"GET",
		"/latest/industry/systems/",
		GetIndustrySystems,
	)

	mockesi.NewRoute(
		"GetInsurancePrices",
		"GET",
		"/latest/insurance/prices/",
		GetInsurancePrices,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdKillmailsRecent",
		"GET",
		"/latest/characters/{character_id}/killmails/recent/",
		GetCharactersCharacterIdKillmailsRecent,
	)

	mockesi.NewRoute(
		"GetKillmailsKillmailIdKillmailHash",
		"GET",
		"/latest/killmails/{killmail_id}/{killmail_hash}/",
		GetKillmailsKillmailIdKillmailHash,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdLoyaltyPoints",
		"GET",
		"/latest/characters/{character_id}/loyalty/points/",
		GetCharactersCharacterIdLoyaltyPoints,
	)

	mockesi.NewRoute(
		"GetLoyaltyStoresCorporationIdOffers",
		"GET",
		"/latest/loyalty/stores/{corporation_id}/offers/",
		GetLoyaltyStoresCorporationIdOffers,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdMailLabelsLabelId",
		"DELETE",
		"/latest/characters/{character_id}/mail/labels/{label_id}/",
		DeleteCharactersCharacterIdMailLabelsLabelId,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdMailMailId",
		"DELETE",
		"/latest/characters/{character_id}/mail/{mail_id}/",
		DeleteCharactersCharacterIdMailMailId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMail",
		"GET",
		"/latest/characters/{character_id}/mail/",
		GetCharactersCharacterIdMail,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailLabels",
		"GET",
		"/latest/characters/{character_id}/mail/labels/",
		GetCharactersCharacterIdMailLabels,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailLists",
		"GET",
		"/latest/characters/{character_id}/mail/lists/",
		GetCharactersCharacterIdMailLists,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailMailId",
		"GET",
		"/latest/characters/{character_id}/mail/{mail_id}/",
		GetCharactersCharacterIdMailMailId,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdMail",
		"POST",
		"/latest/characters/{character_id}/mail/",
		PostCharactersCharacterIdMail,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdMailLabels",
		"POST",
		"/latest/characters/{character_id}/mail/labels/",
		PostCharactersCharacterIdMailLabels,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdMailMailId",
		"PUT",
		"/latest/characters/{character_id}/mail/{mail_id}/",
		PutCharactersCharacterIdMailMailId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOrders",
		"GET",
		"/latest/characters/{character_id}/orders/",
		GetCharactersCharacterIdOrders,
	)

	mockesi.NewRoute(
		"GetMarketsGroups",
		"GET",
		"/latest/markets/groups/",
		GetMarketsGroups,
	)

	mockesi.NewRoute(
		"GetMarketsGroupsMarketGroupId",
		"GET",
		"/latest/markets/groups/{market_group_id}/",
		GetMarketsGroupsMarketGroupId,
	)

	mockesi.NewRoute(
		"GetMarketsPrices",
		"GET",
		"/latest/markets/prices/",
		GetMarketsPrices,
	)

	mockesi.NewRoute(
		"GetMarketsRegionIdHistory",
		"GET",
		"/latest/markets/{region_id}/history/",
		GetMarketsRegionIdHistory,
	)

	mockesi.NewRoute(
		"GetMarketsRegionIdOrders",
		"GET",
		"/latest/markets/{region_id}/orders/",
		GetMarketsRegionIdOrders,
	)

	mockesi.NewRoute(
		"GetMarketsStructuresStructureId",
		"GET",
		"/latest/markets/structures/{structure_id}/",
		GetMarketsStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOpportunities",
		"GET",
		"/latest/characters/{character_id}/opportunities/",
		GetCharactersCharacterIdOpportunities,
	)

	mockesi.NewRoute(
		"GetOpportunitiesGroups",
		"GET",
		"/latest/opportunities/groups/",
		GetOpportunitiesGroups,
	)

	mockesi.NewRoute(
		"GetOpportunitiesGroupsGroupId",
		"GET",
		"/latest/opportunities/groups/{group_id}/",
		GetOpportunitiesGroupsGroupId,
	)

	mockesi.NewRoute(
		"GetOpportunitiesTasks",
		"GET",
		"/latest/opportunities/tasks/",
		GetOpportunitiesTasks,
	)

	mockesi.NewRoute(
		"GetOpportunitiesTasksTaskId",
		"GET",
		"/latest/opportunities/tasks/{task_id}/",
		GetOpportunitiesTasksTaskId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPlanets",
		"GET",
		"/latest/characters/{character_id}/planets/",
		GetCharactersCharacterIdPlanets,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPlanetsPlanetId",
		"GET",
		"/latest/characters/{character_id}/planets/{planet_id}/",
		GetCharactersCharacterIdPlanetsPlanetId,
	)

	mockesi.NewRoute(
		"GetUniverseSchematicsSchematicId",
		"GET",
		"/latest/universe/schematics/{schematic_id}/",
		GetUniverseSchematicsSchematicId,
	)

	mockesi.NewRoute(
		"GetRouteOriginDestination",
		"GET",
		"/latest/route/{origin}/{destination}/",
		GetRouteOriginDestination,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdSearch",
		"GET",
		"/latest/characters/{character_id}/search/",
		GetCharactersCharacterIdSearch,
	)

	mockesi.NewRoute(
		"GetSearch",
		"GET",
		"/latest/search/",
		GetSearch,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAttributes",
		"GET",
		"/latest/characters/{character_id}/attributes/",
		GetCharactersCharacterIdAttributes,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdSkillqueue",
		"GET",
		"/latest/characters/{character_id}/skillqueue/",
		GetCharactersCharacterIdSkillqueue,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdSkills",
		"GET",
		"/latest/characters/{character_id}/skills/",
		GetCharactersCharacterIdSkills,
	)

	mockesi.NewRoute(
		"GetSovereigntyCampaigns",
		"GET",
		"/latest/sovereignty/campaigns/",
		GetSovereigntyCampaigns,
	)

	mockesi.NewRoute(
		"GetSovereigntyMap",
		"GET",
		"/latest/sovereignty/map/",
		GetSovereigntyMap,
	)

	mockesi.NewRoute(
		"GetSovereigntyStructures",
		"GET",
		"/latest/sovereignty/structures/",
		GetSovereigntyStructures,
	)

	mockesi.NewRoute(
		"GetStatus",
		"GET",
		"/latest/status/",
		GetStatus,
	)

	mockesi.NewRoute(
		"GetUniverseBloodlines",
		"GET",
		"/latest/universe/bloodlines/",
		GetUniverseBloodlines,
	)

	mockesi.NewRoute(
		"GetUniverseCategories",
		"GET",
		"/latest/universe/categories/",
		GetUniverseCategories,
	)

	mockesi.NewRoute(
		"GetUniverseCategoriesCategoryId",
		"GET",
		"/latest/universe/categories/{category_id}/",
		GetUniverseCategoriesCategoryId,
	)

	mockesi.NewRoute(
		"GetUniverseConstellations",
		"GET",
		"/latest/universe/constellations/",
		GetUniverseConstellations,
	)

	mockesi.NewRoute(
		"GetUniverseConstellationsConstellationId",
		"GET",
		"/latest/universe/constellations/{constellation_id}/",
		GetUniverseConstellationsConstellationId,
	)

	mockesi.NewRoute(
		"GetUniverseFactions",
		"GET",
		"/latest/universe/factions/",
		GetUniverseFactions,
	)

	mockesi.NewRoute(
		"GetUniverseGraphics",
		"GET",
		"/latest/universe/graphics/",
		GetUniverseGraphics,
	)

	mockesi.NewRoute(
		"GetUniverseGraphicsGraphicId",
		"GET",
		"/latest/universe/graphics/{graphic_id}/",
		GetUniverseGraphicsGraphicId,
	)

	mockesi.NewRoute(
		"GetUniverseGroups",
		"GET",
		"/latest/universe/groups/",
		GetUniverseGroups,
	)

	mockesi.NewRoute(
		"GetUniverseGroupsGroupId",
		"GET",
		"/latest/universe/groups/{group_id}/",
		GetUniverseGroupsGroupId,
	)

	mockesi.NewRoute(
		"GetUniverseMoonsMoonId",
		"GET",
		"/latest/universe/moons/{moon_id}/",
		GetUniverseMoonsMoonId,
	)

	mockesi.NewRoute(
		"GetUniversePlanetsPlanetId",
		"GET",
		"/latest/universe/planets/{planet_id}/",
		GetUniversePlanetsPlanetId,
	)

	mockesi.NewRoute(
		"GetUniverseRaces",
		"GET",
		"/latest/universe/races/",
		GetUniverseRaces,
	)

	mockesi.NewRoute(
		"GetUniverseRegions",
		"GET",
		"/latest/universe/regions/",
		GetUniverseRegions,
	)

	mockesi.NewRoute(
		"GetUniverseRegionsRegionId",
		"GET",
		"/latest/universe/regions/{region_id}/",
		GetUniverseRegionsRegionId,
	)

	mockesi.NewRoute(
		"GetUniverseStargatesStargateId",
		"GET",
		"/latest/universe/stargates/{stargate_id}/",
		GetUniverseStargatesStargateId,
	)

	mockesi.NewRoute(
		"GetUniverseStationsStationId",
		"GET",
		"/latest/universe/stations/{station_id}/",
		GetUniverseStationsStationId,
	)

	mockesi.NewRoute(
		"GetUniverseStructures",
		"GET",
		"/latest/universe/structures/",
		GetUniverseStructures,
	)

	mockesi.NewRoute(
		"GetUniverseStructuresStructureId",
		"GET",
		"/latest/universe/structures/{structure_id}/",
		GetUniverseStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetUniverseSystemJumps",
		"GET",
		"/latest/universe/system_jumps/",
		GetUniverseSystemJumps,
	)

	mockesi.NewRoute(
		"GetUniverseSystemKills",
		"GET",
		"/latest/universe/system_kills/",
		GetUniverseSystemKills,
	)

	mockesi.NewRoute(
		"GetUniverseSystems",
		"GET",
		"/latest/universe/systems/",
		GetUniverseSystems,
	)

	mockesi.NewRoute(
		"GetUniverseSystemsSystemId",
		"GET",
		"/latest/universe/systems/{system_id}/",
		GetUniverseSystemsSystemId,
	)

	mockesi.NewRoute(
		"GetUniverseTypes",
		"GET",
		"/latest/universe/types/",
		GetUniverseTypes,
	)

	mockesi.NewRoute(
		"GetUniverseTypesTypeId",
		"GET",
		"/latest/universe/types/{type_id}/",
		GetUniverseTypesTypeId,
	)

	mockesi.NewRoute(
		"PostUiAutopilotWaypoint",
		"POST",
		"/latest/ui/autopilot/waypoint/",
		PostUiAutopilotWaypoint,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowContract",
		"POST",
		"/latest/ui/openwindow/contract/",
		PostUiOpenwindowContract,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowInformation",
		"POST",
		"/latest/ui/openwindow/information/",
		PostUiOpenwindowInformation,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowMarketdetails",
		"POST",
		"/latest/ui/openwindow/marketdetails/",
		PostUiOpenwindowMarketdetails,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowNewmail",
		"POST",
		"/latest/ui/openwindow/newmail/",
		PostUiOpenwindowNewmail,
	)

	mockesi.NewRoute(
		"GetWars",
		"GET",
		"/latest/wars/",
		GetWars,
	)

	mockesi.NewRoute(
		"GetWarsWarId",
		"GET",
		"/latest/wars/{war_id}/",
		GetWarsWarId,
	)

	mockesi.NewRoute(
		"GetWarsWarIdKillmails",
		"GET",
		"/latest/wars/{war_id}/killmails/",
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
