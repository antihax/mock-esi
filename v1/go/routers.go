package esiv1

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
		"/v1/alliances/",
		GetAlliances,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceId",
		"GET",
		"/v1/alliances/{alliance_id}/",
		GetAlliancesAllianceId,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdCorporations",
		"GET",
		"/v1/alliances/{alliance_id}/corporations/",
		GetAlliancesAllianceIdCorporations,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdIcons",
		"GET",
		"/v1/alliances/{alliance_id}/icons/",
		GetAlliancesAllianceIdIcons,
	)

	mockesi.NewRoute(
		"GetAlliancesNames",
		"GET",
		"/v1/alliances/names/",
		GetAlliancesNames,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAssets",
		"GET",
		"/v1/characters/{character_id}/assets/",
		GetCharactersCharacterIdAssets,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBookmarks",
		"GET",
		"/v1/characters/{character_id}/bookmarks/",
		GetCharactersCharacterIdBookmarks,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBookmarksFolders",
		"GET",
		"/v1/characters/{character_id}/bookmarks/folders/",
		GetCharactersCharacterIdBookmarksFolders,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCalendar",
		"GET",
		"/v1/characters/{character_id}/calendar/",
		GetCharactersCharacterIdCalendar,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAgentsResearch",
		"GET",
		"/v1/characters/{character_id}/agents_research/",
		GetCharactersCharacterIdAgentsResearch,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBlueprints",
		"GET",
		"/v1/characters/{character_id}/blueprints/",
		GetCharactersCharacterIdBlueprints,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdChatChannels",
		"GET",
		"/v1/characters/{character_id}/chat_channels/",
		GetCharactersCharacterIdChatChannels,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCorporationhistory",
		"GET",
		"/v1/characters/{character_id}/corporationhistory/",
		GetCharactersCharacterIdCorporationhistory,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFatigue",
		"GET",
		"/v1/characters/{character_id}/fatigue/",
		GetCharactersCharacterIdFatigue,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMedals",
		"GET",
		"/v1/characters/{character_id}/medals/",
		GetCharactersCharacterIdMedals,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPortrait",
		"GET",
		"/v1/characters/{character_id}/portrait/",
		GetCharactersCharacterIdPortrait,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdRoles",
		"GET",
		"/v1/characters/{character_id}/roles/",
		GetCharactersCharacterIdRoles,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdStandings",
		"GET",
		"/v1/characters/{character_id}/standings/",
		GetCharactersCharacterIdStandings,
	)

	mockesi.NewRoute(
		"GetCharactersNames",
		"GET",
		"/v1/characters/names/",
		GetCharactersNames,
	)

	mockesi.NewRoute(
		"PostCharactersAffiliation",
		"POST",
		"/v1/characters/affiliation/",
		PostCharactersAffiliation,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdImplants",
		"GET",
		"/v1/characters/{character_id}/implants/",
		GetCharactersCharacterIdImplants,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdContacts",
		"DELETE",
		"/v1/characters/{character_id}/contacts/",
		DeleteCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContacts",
		"GET",
		"/v1/characters/{character_id}/contacts/",
		GetCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContactsLabels",
		"GET",
		"/v1/characters/{character_id}/contacts/labels/",
		GetCharactersCharacterIdContactsLabels,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdContacts",
		"POST",
		"/v1/characters/{character_id}/contacts/",
		PostCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdContacts",
		"PUT",
		"/v1/characters/{character_id}/contacts/",
		PutCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdAlliancehistory",
		"GET",
		"/v1/corporations/{corporation_id}/alliancehistory/",
		GetCorporationsCorporationIdAlliancehistory,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdIcons",
		"GET",
		"/v1/corporations/{corporation_id}/icons/",
		GetCorporationsCorporationIdIcons,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdRoles",
		"GET",
		"/v1/corporations/{corporation_id}/roles/",
		GetCorporationsCorporationIdRoles,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdStructures",
		"GET",
		"/v1/corporations/{corporation_id}/structures/",
		GetCorporationsCorporationIdStructures,
	)

	mockesi.NewRoute(
		"GetCorporationsNames",
		"GET",
		"/v1/corporations/names/",
		GetCorporationsNames,
	)

	mockesi.NewRoute(
		"GetCorporationsNpccorps",
		"GET",
		"/v1/corporations/npccorps/",
		GetCorporationsNpccorps,
	)

	mockesi.NewRoute(
		"PutCorporationsCorporationIdStructuresStructureId",
		"PUT",
		"/v1/corporations/{corporation_id}/structures/{structure_id}/",
		PutCorporationsCorporationIdStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetDogmaAttributes",
		"GET",
		"/v1/dogma/attributes/",
		GetDogmaAttributes,
	)

	mockesi.NewRoute(
		"GetDogmaAttributesAttributeId",
		"GET",
		"/v1/dogma/attributes/{attribute_id}/",
		GetDogmaAttributesAttributeId,
	)

	mockesi.NewRoute(
		"GetDogmaEffects",
		"GET",
		"/v1/dogma/effects/",
		GetDogmaEffects,
	)

	mockesi.NewRoute(
		"GetDogmaEffectsEffectId",
		"GET",
		"/v1/dogma/effects/{effect_id}/",
		GetDogmaEffectsEffectId,
	)

	mockesi.NewRoute(
		"GetFwStats",
		"GET",
		"/v1/fw/stats/",
		GetFwStats,
	)

	mockesi.NewRoute(
		"GetFwWars",
		"GET",
		"/v1/fw/wars/",
		GetFwWars,
	)

	mockesi.NewRoute(
		"GetIncursions",
		"GET",
		"/v1/incursions/",
		GetIncursions,
	)

	mockesi.NewRoute(
		"GetInsurancePrices",
		"GET",
		"/v1/insurance/prices/",
		GetInsurancePrices,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdLoyaltyPoints",
		"GET",
		"/v1/characters/{character_id}/loyalty/points/",
		GetCharactersCharacterIdLoyaltyPoints,
	)

	mockesi.NewRoute(
		"GetLoyaltyStoresCorporationIdOffers",
		"GET",
		"/v1/loyalty/stores/{corporation_id}/offers/",
		GetLoyaltyStoresCorporationIdOffers,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOrders",
		"GET",
		"/v1/characters/{character_id}/orders/",
		GetCharactersCharacterIdOrders,
	)

	mockesi.NewRoute(
		"GetMarketsGroups",
		"GET",
		"/v1/markets/groups/",
		GetMarketsGroups,
	)

	mockesi.NewRoute(
		"GetMarketsGroupsMarketGroupId",
		"GET",
		"/v1/markets/groups/{market_group_id}/",
		GetMarketsGroupsMarketGroupId,
	)

	mockesi.NewRoute(
		"GetMarketsPrices",
		"GET",
		"/v1/markets/prices/",
		GetMarketsPrices,
	)

	mockesi.NewRoute(
		"GetMarketsRegionIdHistory",
		"GET",
		"/v1/markets/{region_id}/history/",
		GetMarketsRegionIdHistory,
	)

	mockesi.NewRoute(
		"GetMarketsRegionIdOrders",
		"GET",
		"/v1/markets/{region_id}/orders/",
		GetMarketsRegionIdOrders,
	)

	mockesi.NewRoute(
		"GetMarketsStructuresStructureId",
		"GET",
		"/v1/markets/structures/{structure_id}/",
		GetMarketsStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOpportunities",
		"GET",
		"/v1/characters/{character_id}/opportunities/",
		GetCharactersCharacterIdOpportunities,
	)

	mockesi.NewRoute(
		"GetOpportunitiesGroups",
		"GET",
		"/v1/opportunities/groups/",
		GetOpportunitiesGroups,
	)

	mockesi.NewRoute(
		"GetOpportunitiesGroupsGroupId",
		"GET",
		"/v1/opportunities/groups/{group_id}/",
		GetOpportunitiesGroupsGroupId,
	)

	mockesi.NewRoute(
		"GetOpportunitiesTasks",
		"GET",
		"/v1/opportunities/tasks/",
		GetOpportunitiesTasks,
	)

	mockesi.NewRoute(
		"GetOpportunitiesTasksTaskId",
		"GET",
		"/v1/opportunities/tasks/{task_id}/",
		GetOpportunitiesTasksTaskId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPlanets",
		"GET",
		"/v1/characters/{character_id}/planets/",
		GetCharactersCharacterIdPlanets,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPlanetsPlanetId",
		"GET",
		"/v1/characters/{character_id}/planets/{planet_id}/",
		GetCharactersCharacterIdPlanetsPlanetId,
	)

	mockesi.NewRoute(
		"GetUniverseSchematicsSchematicId",
		"GET",
		"/v1/universe/schematics/{schematic_id}/",
		GetUniverseSchematicsSchematicId,
	)

	mockesi.NewRoute(
		"GetRouteOriginDestination",
		"GET",
		"/v1/route/{origin}/{destination}/",
		GetRouteOriginDestination,
	)

	mockesi.NewRoute(
		"GetSovereigntyCampaigns",
		"GET",
		"/v1/sovereignty/campaigns/",
		GetSovereigntyCampaigns,
	)

	mockesi.NewRoute(
		"GetSovereigntyMap",
		"GET",
		"/v1/sovereignty/map/",
		GetSovereigntyMap,
	)

	mockesi.NewRoute(
		"GetSovereigntyStructures",
		"GET",
		"/v1/sovereignty/structures/",
		GetSovereigntyStructures,
	)

	mockesi.NewRoute(
		"GetStatus",
		"GET",
		"/v1/status/",
		GetStatus,
	)

	mockesi.NewRoute(
		"PostUiAutopilotWaypoint",
		"POST",
		"/v1/ui/autopilot/waypoint/",
		PostUiAutopilotWaypoint,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowContract",
		"POST",
		"/v1/ui/openwindow/contract/",
		PostUiOpenwindowContract,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowInformation",
		"POST",
		"/v1/ui/openwindow/information/",
		PostUiOpenwindowInformation,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowMarketdetails",
		"POST",
		"/v1/ui/openwindow/marketdetails/",
		PostUiOpenwindowMarketdetails,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowNewmail",
		"POST",
		"/v1/ui/openwindow/newmail/",
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
