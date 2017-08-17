package esidev

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
		"/dev/alliances/",
		GetAlliances,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceId",
		"GET",
		"/dev/alliances/{alliance_id}/",
		GetAlliancesAllianceId,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdCorporations",
		"GET",
		"/dev/alliances/{alliance_id}/corporations/",
		GetAlliancesAllianceIdCorporations,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdIcons",
		"GET",
		"/dev/alliances/{alliance_id}/icons/",
		GetAlliancesAllianceIdIcons,
	)

	mockesi.NewRoute(
		"GetAlliancesNames",
		"GET",
		"/dev/alliances/names/",
		GetAlliancesNames,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAssets",
		"GET",
		"/dev/characters/{character_id}/assets/",
		GetCharactersCharacterIdAssets,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBookmarks",
		"GET",
		"/dev/characters/{character_id}/bookmarks/",
		GetCharactersCharacterIdBookmarks,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBookmarksFolders",
		"GET",
		"/dev/characters/{character_id}/bookmarks/folders/",
		GetCharactersCharacterIdBookmarksFolders,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCalendar",
		"GET",
		"/dev/characters/{character_id}/calendar/",
		GetCharactersCharacterIdCalendar,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCalendarEventId",
		"GET",
		"/dev/characters/{character_id}/calendar/{event_id}/",
		GetCharactersCharacterIdCalendarEventId,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdCalendarEventId",
		"PUT",
		"/dev/characters/{character_id}/calendar/{event_id}/",
		PutCharactersCharacterIdCalendarEventId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterId",
		"GET",
		"/dev/characters/{character_id}/",
		GetCharactersCharacterId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAgentsResearch",
		"GET",
		"/dev/characters/{character_id}/agents_research/",
		GetCharactersCharacterIdAgentsResearch,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBlueprints",
		"GET",
		"/dev/characters/{character_id}/blueprints/",
		GetCharactersCharacterIdBlueprints,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdChatChannels",
		"GET",
		"/dev/characters/{character_id}/chat_channels/",
		GetCharactersCharacterIdChatChannels,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCorporationhistory",
		"GET",
		"/dev/characters/{character_id}/corporationhistory/",
		GetCharactersCharacterIdCorporationhistory,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFatigue",
		"GET",
		"/dev/characters/{character_id}/fatigue/",
		GetCharactersCharacterIdFatigue,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMedals",
		"GET",
		"/dev/characters/{character_id}/medals/",
		GetCharactersCharacterIdMedals,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPortrait",
		"GET",
		"/dev/characters/{character_id}/portrait/",
		GetCharactersCharacterIdPortrait,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdRoles",
		"GET",
		"/dev/characters/{character_id}/roles/",
		GetCharactersCharacterIdRoles,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdStandings",
		"GET",
		"/dev/characters/{character_id}/standings/",
		GetCharactersCharacterIdStandings,
	)

	mockesi.NewRoute(
		"GetCharactersNames",
		"GET",
		"/dev/characters/names/",
		GetCharactersNames,
	)

	mockesi.NewRoute(
		"PostCharactersAffiliation",
		"POST",
		"/dev/characters/affiliation/",
		PostCharactersAffiliation,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdCspa",
		"POST",
		"/dev/characters/{character_id}/cspa/",
		PostCharactersCharacterIdCspa,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdClones",
		"GET",
		"/dev/characters/{character_id}/clones/",
		GetCharactersCharacterIdClones,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdImplants",
		"GET",
		"/dev/characters/{character_id}/implants/",
		GetCharactersCharacterIdImplants,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdContacts",
		"DELETE",
		"/dev/characters/{character_id}/contacts/",
		DeleteCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContacts",
		"GET",
		"/dev/characters/{character_id}/contacts/",
		GetCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContactsLabels",
		"GET",
		"/dev/characters/{character_id}/contacts/labels/",
		GetCharactersCharacterIdContactsLabels,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdContacts",
		"POST",
		"/dev/characters/{character_id}/contacts/",
		PostCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdContacts",
		"PUT",
		"/dev/characters/{character_id}/contacts/",
		PutCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationId",
		"GET",
		"/dev/corporations/{corporation_id}/",
		GetCorporationsCorporationId,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdAlliancehistory",
		"GET",
		"/dev/corporations/{corporation_id}/alliancehistory/",
		GetCorporationsCorporationIdAlliancehistory,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdIcons",
		"GET",
		"/dev/corporations/{corporation_id}/icons/",
		GetCorporationsCorporationIdIcons,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMembers",
		"GET",
		"/dev/corporations/{corporation_id}/members/",
		GetCorporationsCorporationIdMembers,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdRoles",
		"GET",
		"/dev/corporations/{corporation_id}/roles/",
		GetCorporationsCorporationIdRoles,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdStructures",
		"GET",
		"/dev/corporations/{corporation_id}/structures/",
		GetCorporationsCorporationIdStructures,
	)

	mockesi.NewRoute(
		"GetCorporationsNames",
		"GET",
		"/dev/corporations/names/",
		GetCorporationsNames,
	)

	mockesi.NewRoute(
		"GetCorporationsNpccorps",
		"GET",
		"/dev/corporations/npccorps/",
		GetCorporationsNpccorps,
	)

	mockesi.NewRoute(
		"PutCorporationsCorporationIdStructuresStructureId",
		"PUT",
		"/dev/corporations/{corporation_id}/structures/{structure_id}/",
		PutCorporationsCorporationIdStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetDogmaAttributes",
		"GET",
		"/dev/dogma/attributes/",
		GetDogmaAttributes,
	)

	mockesi.NewRoute(
		"GetDogmaAttributesAttributeId",
		"GET",
		"/dev/dogma/attributes/{attribute_id}/",
		GetDogmaAttributesAttributeId,
	)

	mockesi.NewRoute(
		"GetDogmaEffects",
		"GET",
		"/dev/dogma/effects/",
		GetDogmaEffects,
	)

	mockesi.NewRoute(
		"GetDogmaEffectsEffectId",
		"GET",
		"/dev/dogma/effects/{effect_id}/",
		GetDogmaEffectsEffectId,
	)

	mockesi.NewRoute(
		"GetFwStats",
		"GET",
		"/dev/fw/stats/",
		GetFwStats,
	)

	mockesi.NewRoute(
		"GetFwWars",
		"GET",
		"/dev/fw/wars/",
		GetFwWars,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdFittingsFittingId",
		"DELETE",
		"/dev/characters/{character_id}/fittings/{fitting_id}/",
		DeleteCharactersCharacterIdFittingsFittingId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFittings",
		"GET",
		"/dev/characters/{character_id}/fittings/",
		GetCharactersCharacterIdFittings,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdFittings",
		"POST",
		"/dev/characters/{character_id}/fittings/",
		PostCharactersCharacterIdFittings,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdMembersMemberId",
		"DELETE",
		"/dev/fleets/{fleet_id}/members/{member_id}/",
		DeleteFleetsFleetIdMembersMemberId,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdSquadsSquadId",
		"DELETE",
		"/dev/fleets/{fleet_id}/squads/{squad_id}/",
		DeleteFleetsFleetIdSquadsSquadId,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdWingsWingId",
		"DELETE",
		"/dev/fleets/{fleet_id}/wings/{wing_id}/",
		DeleteFleetsFleetIdWingsWingId,
	)

	mockesi.NewRoute(
		"GetFleetsFleetId",
		"GET",
		"/dev/fleets/{fleet_id}/",
		GetFleetsFleetId,
	)

	mockesi.NewRoute(
		"GetFleetsFleetIdMembers",
		"GET",
		"/dev/fleets/{fleet_id}/members/",
		GetFleetsFleetIdMembers,
	)

	mockesi.NewRoute(
		"GetFleetsFleetIdWings",
		"GET",
		"/dev/fleets/{fleet_id}/wings/",
		GetFleetsFleetIdWings,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdMembers",
		"POST",
		"/dev/fleets/{fleet_id}/members/",
		PostFleetsFleetIdMembers,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdWings",
		"POST",
		"/dev/fleets/{fleet_id}/wings/",
		PostFleetsFleetIdWings,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdWingsWingIdSquads",
		"POST",
		"/dev/fleets/{fleet_id}/wings/{wing_id}/squads/",
		PostFleetsFleetIdWingsWingIdSquads,
	)

	mockesi.NewRoute(
		"PutFleetsFleetId",
		"PUT",
		"/dev/fleets/{fleet_id}/",
		PutFleetsFleetId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdMembersMemberId",
		"PUT",
		"/dev/fleets/{fleet_id}/members/{member_id}/",
		PutFleetsFleetIdMembersMemberId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdSquadsSquadId",
		"PUT",
		"/dev/fleets/{fleet_id}/squads/{squad_id}/",
		PutFleetsFleetIdSquadsSquadId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdWingsWingId",
		"PUT",
		"/dev/fleets/{fleet_id}/wings/{wing_id}/",
		PutFleetsFleetIdWingsWingId,
	)

	mockesi.NewRoute(
		"GetIncursions",
		"GET",
		"/dev/incursions/",
		GetIncursions,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdIndustryJobs",
		"GET",
		"/dev/characters/{character_id}/industry/jobs/",
		GetCharactersCharacterIdIndustryJobs,
	)

	mockesi.NewRoute(
		"GetIndustryFacilities",
		"GET",
		"/dev/industry/facilities/",
		GetIndustryFacilities,
	)

	mockesi.NewRoute(
		"GetIndustrySystems",
		"GET",
		"/dev/industry/systems/",
		GetIndustrySystems,
	)

	mockesi.NewRoute(
		"GetInsurancePrices",
		"GET",
		"/dev/insurance/prices/",
		GetInsurancePrices,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdLoyaltyPoints",
		"GET",
		"/dev/characters/{character_id}/loyalty/points/",
		GetCharactersCharacterIdLoyaltyPoints,
	)

	mockesi.NewRoute(
		"GetLoyaltyStoresCorporationIdOffers",
		"GET",
		"/dev/loyalty/stores/{corporation_id}/offers/",
		GetLoyaltyStoresCorporationIdOffers,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOrders",
		"GET",
		"/dev/characters/{character_id}/orders/",
		GetCharactersCharacterIdOrders,
	)

	mockesi.NewRoute(
		"GetMarketsGroups",
		"GET",
		"/dev/markets/groups/",
		GetMarketsGroups,
	)

	mockesi.NewRoute(
		"GetMarketsGroupsMarketGroupId",
		"GET",
		"/dev/markets/groups/{market_group_id}/",
		GetMarketsGroupsMarketGroupId,
	)

	mockesi.NewRoute(
		"GetMarketsPrices",
		"GET",
		"/dev/markets/prices/",
		GetMarketsPrices,
	)

	mockesi.NewRoute(
		"GetMarketsRegionIdHistory",
		"GET",
		"/dev/markets/{region_id}/history/",
		GetMarketsRegionIdHistory,
	)

	mockesi.NewRoute(
		"GetMarketsRegionIdOrders",
		"GET",
		"/dev/markets/{region_id}/orders/",
		GetMarketsRegionIdOrders,
	)

	mockesi.NewRoute(
		"GetMarketsStructuresStructureId",
		"GET",
		"/dev/markets/structures/{structure_id}/",
		GetMarketsStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOpportunities",
		"GET",
		"/dev/characters/{character_id}/opportunities/",
		GetCharactersCharacterIdOpportunities,
	)

	mockesi.NewRoute(
		"GetOpportunitiesGroups",
		"GET",
		"/dev/opportunities/groups/",
		GetOpportunitiesGroups,
	)

	mockesi.NewRoute(
		"GetOpportunitiesGroupsGroupId",
		"GET",
		"/dev/opportunities/groups/{group_id}/",
		GetOpportunitiesGroupsGroupId,
	)

	mockesi.NewRoute(
		"GetOpportunitiesTasks",
		"GET",
		"/dev/opportunities/tasks/",
		GetOpportunitiesTasks,
	)

	mockesi.NewRoute(
		"GetOpportunitiesTasksTaskId",
		"GET",
		"/dev/opportunities/tasks/{task_id}/",
		GetOpportunitiesTasksTaskId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPlanets",
		"GET",
		"/dev/characters/{character_id}/planets/",
		GetCharactersCharacterIdPlanets,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPlanetsPlanetId",
		"GET",
		"/dev/characters/{character_id}/planets/{planet_id}/",
		GetCharactersCharacterIdPlanetsPlanetId,
	)

	mockesi.NewRoute(
		"GetUniverseSchematicsSchematicId",
		"GET",
		"/dev/universe/schematics/{schematic_id}/",
		GetUniverseSchematicsSchematicId,
	)

	mockesi.NewRoute(
		"GetRouteOriginDestination",
		"GET",
		"/dev/route/{origin}/{destination}/",
		GetRouteOriginDestination,
	)

	mockesi.NewRoute(
		"GetSovereigntyCampaigns",
		"GET",
		"/dev/sovereignty/campaigns/",
		GetSovereigntyCampaigns,
	)

	mockesi.NewRoute(
		"GetSovereigntyMap",
		"GET",
		"/dev/sovereignty/map/",
		GetSovereigntyMap,
	)

	mockesi.NewRoute(
		"GetSovereigntyStructures",
		"GET",
		"/dev/sovereignty/structures/",
		GetSovereigntyStructures,
	)

	mockesi.NewRoute(
		"GetStatus",
		"GET",
		"/dev/status/",
		GetStatus,
	)

	mockesi.NewRoute(
		"PostUiAutopilotWaypoint",
		"POST",
		"/dev/ui/autopilot/waypoint/",
		PostUiAutopilotWaypoint,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowContract",
		"POST",
		"/dev/ui/openwindow/contract/",
		PostUiOpenwindowContract,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowInformation",
		"POST",
		"/dev/ui/openwindow/information/",
		PostUiOpenwindowInformation,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowMarketdetails",
		"POST",
		"/dev/ui/openwindow/marketdetails/",
		PostUiOpenwindowMarketdetails,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowNewmail",
		"POST",
		"/dev/ui/openwindow/newmail/",
		PostUiOpenwindowNewmail,
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
