package esiV1

import (
	"net/http"
	"github.com/antihax/mock-esi/mockesi"
	"strconv"
	"reflect"
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
		"GetCharactersCharacterIdCorporationhistory",
		"GET",
		"/v1/characters/{character_id}/corporationhistory/",
		GetCharactersCharacterIdCorporationhistory,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPortrait",
		"GET",
		"/v1/characters/{character_id}/portrait/",
		GetCharactersCharacterIdPortrait,
	)

	mockesi.NewRoute(
		"GetCharactersNames",
		"GET",
		"/v1/characters/names/",
		GetCharactersNames,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdClones",
		"GET",
		"/v1/characters/{character_id}/clones/",
		GetCharactersCharacterIdClones,
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
		"GetCorporationsCorporationId",
		"GET",
		"/v1/corporations/{corporation_id}/",
		GetCorporationsCorporationId,
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
		"DeleteCharactersCharacterIdFittingsFittingId",
		"DELETE",
		"/v1/characters/{character_id}/fittings/{fitting_id}/",
		DeleteCharactersCharacterIdFittingsFittingId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFittings",
		"GET",
		"/v1/characters/{character_id}/fittings/",
		GetCharactersCharacterIdFittings,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdFittings",
		"POST",
		"/v1/characters/{character_id}/fittings/",
		PostCharactersCharacterIdFittings,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdMembersMemberId",
		"DELETE",
		"/v1/fleets/{fleet_id}/members/{member_id}/",
		DeleteFleetsFleetIdMembersMemberId,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdSquadsSquadId",
		"DELETE",
		"/v1/fleets/{fleet_id}/squads/{squad_id}/",
		DeleteFleetsFleetIdSquadsSquadId,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdWingsWingId",
		"DELETE",
		"/v1/fleets/{fleet_id}/wings/{wing_id}/",
		DeleteFleetsFleetIdWingsWingId,
	)

	mockesi.NewRoute(
		"GetFleetsFleetId",
		"GET",
		"/v1/fleets/{fleet_id}/",
		GetFleetsFleetId,
	)

	mockesi.NewRoute(
		"GetFleetsFleetIdMembers",
		"GET",
		"/v1/fleets/{fleet_id}/members/",
		GetFleetsFleetIdMembers,
	)

	mockesi.NewRoute(
		"GetFleetsFleetIdWings",
		"GET",
		"/v1/fleets/{fleet_id}/wings/",
		GetFleetsFleetIdWings,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdMembers",
		"POST",
		"/v1/fleets/{fleet_id}/members/",
		PostFleetsFleetIdMembers,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdWings",
		"POST",
		"/v1/fleets/{fleet_id}/wings/",
		PostFleetsFleetIdWings,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdWingsWingIdSquads",
		"POST",
		"/v1/fleets/{fleet_id}/wings/{wing_id}/squads/",
		PostFleetsFleetIdWingsWingIdSquads,
	)

	mockesi.NewRoute(
		"PutFleetsFleetId",
		"PUT",
		"/v1/fleets/{fleet_id}/",
		PutFleetsFleetId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdMembersMemberId",
		"PUT",
		"/v1/fleets/{fleet_id}/members/{member_id}/",
		PutFleetsFleetIdMembersMemberId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdSquadsSquadId",
		"PUT",
		"/v1/fleets/{fleet_id}/squads/{squad_id}/",
		PutFleetsFleetIdSquadsSquadId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdWingsWingId",
		"PUT",
		"/v1/fleets/{fleet_id}/wings/{wing_id}/",
		PutFleetsFleetIdWingsWingId,
	)

	mockesi.NewRoute(
		"GetIncursions",
		"GET",
		"/v1/incursions/",
		GetIncursions,
	)

	mockesi.NewRoute(
		"GetIndustryFacilities",
		"GET",
		"/v1/industry/facilities/",
		GetIndustryFacilities,
	)

	mockesi.NewRoute(
		"GetIndustrySystems",
		"GET",
		"/v1/industry/systems/",
		GetIndustrySystems,
	)

	mockesi.NewRoute(
		"GetInsurancePrices",
		"GET",
		"/v1/insurance/prices/",
		GetInsurancePrices,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdKillmailsRecent",
		"GET",
		"/v1/characters/{character_id}/killmails/recent/",
		GetCharactersCharacterIdKillmailsRecent,
	)

	mockesi.NewRoute(
		"GetKillmailsKillmailIdKillmailHash",
		"GET",
		"/v1/killmails/{killmail_id}/{killmail_hash}/",
		GetKillmailsKillmailIdKillmailHash,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdLocation",
		"GET",
		"/v1/characters/{character_id}/location/",
		GetCharactersCharacterIdLocation,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdShip",
		"GET",
		"/v1/characters/{character_id}/ship/",
		GetCharactersCharacterIdShip,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdMailLabelsLabelId",
		"DELETE",
		"/v1/characters/{character_id}/mail/labels/{label_id}/",
		DeleteCharactersCharacterIdMailLabelsLabelId,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdMailMailId",
		"DELETE",
		"/v1/characters/{character_id}/mail/{mail_id}/",
		DeleteCharactersCharacterIdMailMailId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMail",
		"GET",
		"/v1/characters/{character_id}/mail/",
		GetCharactersCharacterIdMail,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailLists",
		"GET",
		"/v1/characters/{character_id}/mail/lists/",
		GetCharactersCharacterIdMailLists,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailMailId",
		"GET",
		"/v1/characters/{character_id}/mail/{mail_id}/",
		GetCharactersCharacterIdMailMailId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailUnread",
		"GET",
		"/v1/characters/{character_id}/mail/unread/",
		GetCharactersCharacterIdMailUnread,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdMail",
		"POST",
		"/v1/characters/{character_id}/mail/",
		PostCharactersCharacterIdMail,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdMailMailId",
		"PUT",
		"/v1/characters/{character_id}/mail/{mail_id}/",
		PutCharactersCharacterIdMailMailId,
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
		"GetCharactersCharacterIdSearch",
		"GET",
		"/v1/characters/{character_id}/search/",
		GetCharactersCharacterIdSearch,
	)

	mockesi.NewRoute(
		"GetSearch",
		"GET",
		"/v1/search/",
		GetSearch,
	)

	mockesi.NewRoute(
		"GetSovereigntyCampaigns",
		"GET",
		"/v1/sovereignty/campaigns/",
		GetSovereigntyCampaigns,
	)

	mockesi.NewRoute(
		"GetSovereigntyStructures",
		"GET",
		"/v1/sovereignty/structures/",
		GetSovereigntyStructures,
	)

	mockesi.NewRoute(
		"GetUniverseBloodlines",
		"GET",
		"/v1/universe/bloodlines/",
		GetUniverseBloodlines,
	)

	mockesi.NewRoute(
		"GetUniverseCategories",
		"GET",
		"/v1/universe/categories/",
		GetUniverseCategories,
	)

	mockesi.NewRoute(
		"GetUniverseCategoriesCategoryId",
		"GET",
		"/v1/universe/categories/{category_id}/",
		GetUniverseCategoriesCategoryId,
	)

	mockesi.NewRoute(
		"GetUniverseConstellations",
		"GET",
		"/v1/universe/constellations/",
		GetUniverseConstellations,
	)

	mockesi.NewRoute(
		"GetUniverseConstellationsConstellationId",
		"GET",
		"/v1/universe/constellations/{constellation_id}/",
		GetUniverseConstellationsConstellationId,
	)

	mockesi.NewRoute(
		"GetUniverseFactions",
		"GET",
		"/v1/universe/factions/",
		GetUniverseFactions,
	)

	mockesi.NewRoute(
		"GetUniverseGroups",
		"GET",
		"/v1/universe/groups/",
		GetUniverseGroups,
	)

	mockesi.NewRoute(
		"GetUniverseGroupsGroupId",
		"GET",
		"/v1/universe/groups/{group_id}/",
		GetUniverseGroupsGroupId,
	)

	mockesi.NewRoute(
		"GetUniverseMoonsMoonId",
		"GET",
		"/v1/universe/moons/{moon_id}/",
		GetUniverseMoonsMoonId,
	)

	mockesi.NewRoute(
		"GetUniversePlanetsPlanetId",
		"GET",
		"/v1/universe/planets/{planet_id}/",
		GetUniversePlanetsPlanetId,
	)

	mockesi.NewRoute(
		"GetUniverseRaces",
		"GET",
		"/v1/universe/races/",
		GetUniverseRaces,
	)

	mockesi.NewRoute(
		"GetUniverseRegions",
		"GET",
		"/v1/universe/regions/",
		GetUniverseRegions,
	)

	mockesi.NewRoute(
		"GetUniverseRegionsRegionId",
		"GET",
		"/v1/universe/regions/{region_id}/",
		GetUniverseRegionsRegionId,
	)

	mockesi.NewRoute(
		"GetUniverseStargatesStargateId",
		"GET",
		"/v1/universe/stargates/{stargate_id}/",
		GetUniverseStargatesStargateId,
	)

	mockesi.NewRoute(
		"GetUniverseStationsStationId",
		"GET",
		"/v1/universe/stations/{station_id}/",
		GetUniverseStationsStationId,
	)

	mockesi.NewRoute(
		"GetUniverseStructures",
		"GET",
		"/v1/universe/structures/",
		GetUniverseStructures,
	)

	mockesi.NewRoute(
		"GetUniverseStructuresStructureId",
		"GET",
		"/v1/universe/structures/{structure_id}/",
		GetUniverseStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetUniverseSystems",
		"GET",
		"/v1/universe/systems/",
		GetUniverseSystems,
	)

	mockesi.NewRoute(
		"GetUniverseSystemsSystemId",
		"GET",
		"/v1/universe/systems/{system_id}/",
		GetUniverseSystemsSystemId,
	)

	mockesi.NewRoute(
		"GetUniverseTypes",
		"GET",
		"/v1/universe/types/",
		GetUniverseTypes,
	)

	mockesi.NewRoute(
		"GetUniverseTypesTypeId",
		"GET",
		"/v1/universe/types/{type_id}/",
		GetUniverseTypesTypeId,
	)

	mockesi.NewRoute(
		"PostUniverseNames",
		"POST",
		"/v1/universe/names/",
		PostUniverseNames,
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

	mockesi.NewRoute(
		"GetCharactersCharacterIdWallets",
		"GET",
		"/v1/characters/{character_id}/wallets/",
		GetCharactersCharacterIdWallets,
	)

	mockesi.NewRoute(
		"GetWars",
		"GET",
		"/v1/wars/",
		GetWars,
	)

	mockesi.NewRoute(
		"GetWarsWarId",
		"GET",
		"/v1/wars/{war_id}/",
		GetWarsWarId,
	)

	mockesi.NewRoute(
		"GetWarsWarIdKillmails",
		"GET",
		"/v1/wars/{war_id}/killmails/",
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