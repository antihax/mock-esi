package esiv2

import (
	"net/http"
	"reflect"
	"strconv"

	"github.com/antihax/mock-esi/mockesi"
)

func init() {

	mockesi.NewRoute(
		"GetAlliancesAllianceId",
		"Get",
		"/v2/alliances/{alliance_id}/",
		GetAlliancesAllianceId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAssets",
		"Get",
		"/v2/characters/{character_id}/assets/",
		GetCharactersCharacterIdAssets,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdAssets",
		"Get",
		"/v2/corporations/{corporation_id}/assets/",
		GetCorporationsCorporationIdAssets,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdAssetsLocations",
		"Post",
		"/v2/characters/{character_id}/assets/locations/",
		PostCharactersCharacterIdAssetsLocations,
	)

	mockesi.NewRoute(
		"PostCorporationsCorporationIdAssetsLocations",
		"Post",
		"/v2/corporations/{corporation_id}/assets/locations/",
		PostCorporationsCorporationIdAssetsLocations,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBookmarks",
		"Get",
		"/v2/characters/{character_id}/bookmarks/",
		GetCharactersCharacterIdBookmarks,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBookmarksFolders",
		"Get",
		"/v2/characters/{character_id}/bookmarks/folders/",
		GetCharactersCharacterIdBookmarksFolders,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCalendarEventId",
		"Get",
		"/v2/characters/{character_id}/calendar/{event_id}/",
		GetCharactersCharacterIdCalendarEventId,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdCalendarEventId",
		"Put",
		"/v2/characters/{character_id}/calendar/{event_id}/",
		PutCharactersCharacterIdCalendarEventId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBlueprints",
		"Get",
		"/v2/characters/{character_id}/blueprints/",
		GetCharactersCharacterIdBlueprints,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPortrait",
		"Get",
		"/v2/characters/{character_id}/portrait/",
		GetCharactersCharacterIdPortrait,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdRoles",
		"Get",
		"/v2/characters/{character_id}/roles/",
		GetCharactersCharacterIdRoles,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdStats",
		"Get",
		"/v2/characters/{character_id}/stats/",
		GetCharactersCharacterIdStats,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdClones",
		"Get",
		"/v2/characters/{character_id}/clones/",
		GetCharactersCharacterIdClones,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdContacts",
		"Delete",
		"/v2/characters/{character_id}/contacts/",
		DeleteCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdContacts",
		"Get",
		"/v2/alliances/{alliance_id}/contacts/",
		GetAlliancesAllianceIdContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContacts",
		"Get",
		"/v2/characters/{character_id}/contacts/",
		GetCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContacts",
		"Get",
		"/v2/corporations/{corporation_id}/contacts/",
		GetCorporationsCorporationIdContacts,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdContacts",
		"Post",
		"/v2/characters/{character_id}/contacts/",
		PostCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdContacts",
		"Put",
		"/v2/characters/{character_id}/contacts/",
		PutCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdAlliancehistory",
		"Get",
		"/v2/corporations/{corporation_id}/alliancehistory/",
		GetCorporationsCorporationIdAlliancehistory,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdBlueprints",
		"Get",
		"/v2/corporations/{corporation_id}/blueprints/",
		GetCorporationsCorporationIdBlueprints,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContainersLogs",
		"Get",
		"/v2/corporations/{corporation_id}/containers/logs/",
		GetCorporationsCorporationIdContainersLogs,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMembers",
		"Get",
		"/v2/corporations/{corporation_id}/members/",
		GetCorporationsCorporationIdMembers,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdStructures",
		"Get",
		"/v2/corporations/{corporation_id}/structures/",
		GetCorporationsCorporationIdStructures,
	)

	mockesi.NewRoute(
		"GetDogmaEffectsEffectId",
		"Get",
		"/v2/dogma/effects/{effect_id}/",
		GetDogmaEffectsEffectId,
	)

	mockesi.NewRoute(
		"GetFwSystems",
		"Get",
		"/v2/fw/systems/",
		GetFwSystems,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFittings",
		"Get",
		"/v2/characters/{character_id}/fittings/",
		GetCharactersCharacterIdFittings,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdFittings",
		"Post",
		"/v2/characters/{character_id}/fittings/",
		PostCharactersCharacterIdFittings,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFleet",
		"Get",
		"/v2/characters/{character_id}/fleet/",
		GetCharactersCharacterIdFleet,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOnline",
		"Get",
		"/v2/characters/{character_id}/online/",
		GetCharactersCharacterIdOnline,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailLabels",
		"Get",
		"/v2/characters/{character_id}/mail/labels/",
		GetCharactersCharacterIdMailLabels,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdMailLabels",
		"Post",
		"/v2/characters/{character_id}/mail/labels/",
		PostCharactersCharacterIdMailLabels,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOrders",
		"Get",
		"/v2/characters/{character_id}/orders/",
		GetCharactersCharacterIdOrders,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdOrders",
		"Get",
		"/v2/corporations/{corporation_id}/orders/",
		GetCorporationsCorporationIdOrders,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdOrdersHistory",
		"Get",
		"/v2/corporations/{corporation_id}/orders/history/",
		GetCorporationsCorporationIdOrdersHistory,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPlanetsPlanetId",
		"Get",
		"/v2/characters/{character_id}/planets/{planet_id}/",
		GetCharactersCharacterIdPlanetsPlanetId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdSearch",
		"Get",
		"/v2/characters/{character_id}/search/",
		GetCharactersCharacterIdSearch,
	)

	mockesi.NewRoute(
		"GetSearch",
		"Get",
		"/v2/search/",
		GetSearch,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdSkillqueue",
		"Get",
		"/v2/characters/{character_id}/skillqueue/",
		GetCharactersCharacterIdSkillqueue,
	)

	mockesi.NewRoute(
		"GetUniverseFactions",
		"Get",
		"/v2/universe/factions/",
		GetUniverseFactions,
	)

	mockesi.NewRoute(
		"GetUniverseStationsStationId",
		"Get",
		"/v2/universe/stations/{station_id}/",
		GetUniverseStationsStationId,
	)

	mockesi.NewRoute(
		"GetUniverseStructuresStructureId",
		"Get",
		"/v2/universe/structures/{structure_id}/",
		GetUniverseStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetUniverseSystemKills",
		"Get",
		"/v2/universe/system_kills/",
		GetUniverseSystemKills,
	)

	mockesi.NewRoute(
		"GetUniverseTypesTypeId",
		"Get",
		"/v2/universe/types/{type_id}/",
		GetUniverseTypesTypeId,
	)

	mockesi.NewRoute(
		"PostUniverseNames",
		"Post",
		"/v2/universe/names/",
		PostUniverseNames,
	)

	mockesi.NewRoute(
		"PostUiAutopilotWaypoint",
		"Post",
		"/v2/ui/autopilot/waypoint/",
		PostUiAutopilotWaypoint,
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
