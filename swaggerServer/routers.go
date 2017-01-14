package swaggerServer

import (
	"net/http"
	"fmt"
	"reflect"
	"strconv"
	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/latest/",
		Index,
	},

	Route{
		"GetAlliances",
		"GET",
		"/latest/alliances/",
		GetAlliances,
	},

	Route{
		"GetAlliancesAllianceId",
		"GET",
		"/latest/alliances/{alliance_id}/",
		GetAlliancesAllianceId,
	},

	Route{
		"GetAlliancesAllianceIdCorporations",
		"GET",
		"/latest/alliances/{alliance_id}/corporations/",
		GetAlliancesAllianceIdCorporations,
	},

	Route{
		"GetAlliancesAllianceIdIcons",
		"GET",
		"/latest/alliances/{alliance_id}/icons/",
		GetAlliancesAllianceIdIcons,
	},

	Route{
		"GetAlliancesNames",
		"GET",
		"/latest/alliances/names/",
		GetAlliancesNames,
	},

	Route{
		"GetCharactersCharacterIdAssets",
		"GET",
		"/latest/characters/{character_id}/assets/",
		GetCharactersCharacterIdAssets,
	},

	Route{
		"GetCharactersCharacterIdBookmarks",
		"GET",
		"/latest/characters/{character_id}/bookmarks/",
		GetCharactersCharacterIdBookmarks,
	},

	Route{
		"GetCharactersCharacterIdBookmarksFolders",
		"GET",
		"/latest/characters/{character_id}/bookmarks/folders/",
		GetCharactersCharacterIdBookmarksFolders,
	},

	Route{
		"GetCharactersCharacterIdCalendar",
		"GET",
		"/latest/characters/{character_id}/calendar/",
		GetCharactersCharacterIdCalendar,
	},

	Route{
		"GetCharactersCharacterIdCalendarEventId",
		"GET",
		"/latest/characters/{character_id}/calendar/{event_id}/",
		GetCharactersCharacterIdCalendarEventId,
	},

	Route{
		"PutCharactersCharacterIdCalendarEventId",
		"PUT",
		"/latest/characters/{character_id}/calendar/{event_id}/",
		PutCharactersCharacterIdCalendarEventId,
	},

	Route{
		"GetCharactersCharacterId",
		"GET",
		"/latest/characters/{character_id}/",
		GetCharactersCharacterId,
	},

	Route{
		"GetCharactersCharacterIdCorporationhistory",
		"GET",
		"/latest/characters/{character_id}/corporationhistory/",
		GetCharactersCharacterIdCorporationhistory,
	},

	Route{
		"GetCharactersCharacterIdPortrait",
		"GET",
		"/latest/characters/{character_id}/portrait/",
		GetCharactersCharacterIdPortrait,
	},

	Route{
		"GetCharactersNames",
		"GET",
		"/latest/characters/names/",
		GetCharactersNames,
	},

	Route{
		"PostCharactersCharacterIdCspa",
		"POST",
		"/latest/characters/{character_id}/cspa/",
		PostCharactersCharacterIdCspa,
	},

	Route{
		"GetCharactersCharacterIdClones",
		"GET",
		"/latest/characters/{character_id}/clones/",
		GetCharactersCharacterIdClones,
	},

	Route{
		"DeleteCharactersCharacterIdContacts",
		"DELETE",
		"/latest/characters/{character_id}/contacts/",
		DeleteCharactersCharacterIdContacts,
	},

	Route{
		"GetCharactersCharacterIdContacts",
		"GET",
		"/latest/characters/{character_id}/contacts/",
		GetCharactersCharacterIdContacts,
	},

	Route{
		"GetCharactersCharacterIdContactsLabels",
		"GET",
		"/latest/characters/{character_id}/contacts/labels/",
		GetCharactersCharacterIdContactsLabels,
	},

	Route{
		"PostCharactersCharacterIdContacts",
		"POST",
		"/latest/characters/{character_id}/contacts/",
		PostCharactersCharacterIdContacts,
	},

	Route{
		"PutCharactersCharacterIdContacts",
		"PUT",
		"/latest/characters/{character_id}/contacts/",
		PutCharactersCharacterIdContacts,
	},

	Route{
		"GetCorporationsCorporationId",
		"GET",
		"/latest/corporations/{corporation_id}/",
		GetCorporationsCorporationId,
	},

	Route{
		"GetCorporationsCorporationIdAlliancehistory",
		"GET",
		"/latest/corporations/{corporation_id}/alliancehistory/",
		GetCorporationsCorporationIdAlliancehistory,
	},

	Route{
		"GetCorporationsCorporationIdIcons",
		"GET",
		"/latest/corporations/{corporation_id}/icons/",
		GetCorporationsCorporationIdIcons,
	},

	Route{
		"GetCorporationsCorporationIdMembers",
		"GET",
		"/latest/corporations/{corporation_id}/members/",
		GetCorporationsCorporationIdMembers,
	},

	Route{
		"GetCorporationsCorporationIdRoles",
		"GET",
		"/latest/corporations/{corporation_id}/roles/",
		GetCorporationsCorporationIdRoles,
	},

	Route{
		"GetCorporationsNames",
		"GET",
		"/latest/corporations/names/",
		GetCorporationsNames,
	},

	Route{
		"DeleteFleetsFleetIdMembersMemberId",
		"DELETE",
		"/latest/fleets/{fleet_id}/members/{member_id}/",
		DeleteFleetsFleetIdMembersMemberId,
	},

	Route{
		"DeleteFleetsFleetIdSquadsSquadId",
		"DELETE",
		"/latest/fleets/{fleet_id}/squads/{squad_id}/",
		DeleteFleetsFleetIdSquadsSquadId,
	},

	Route{
		"DeleteFleetsFleetIdWingsWingId",
		"DELETE",
		"/latest/fleets/{fleet_id}/wings/{wing_id}/",
		DeleteFleetsFleetIdWingsWingId,
	},

	Route{
		"GetFleetsFleetId",
		"GET",
		"/latest/fleets/{fleet_id}/",
		GetFleetsFleetId,
	},

	Route{
		"GetFleetsFleetIdMembers",
		"GET",
		"/latest/fleets/{fleet_id}/members/",
		GetFleetsFleetIdMembers,
	},

	Route{
		"GetFleetsFleetIdWings",
		"GET",
		"/latest/fleets/{fleet_id}/wings/",
		GetFleetsFleetIdWings,
	},

	Route{
		"PostFleetsFleetIdMembers",
		"POST",
		"/latest/fleets/{fleet_id}/members/",
		PostFleetsFleetIdMembers,
	},

	Route{
		"PostFleetsFleetIdWings",
		"POST",
		"/latest/fleets/{fleet_id}/wings/",
		PostFleetsFleetIdWings,
	},

	Route{
		"PostFleetsFleetIdWingsWingIdSquads",
		"POST",
		"/latest/fleets/{fleet_id}/wings/{wing_id}/squads/",
		PostFleetsFleetIdWingsWingIdSquads,
	},

	Route{
		"PutFleetsFleetId",
		"PUT",
		"/latest/fleets/{fleet_id}/",
		PutFleetsFleetId,
	},

	Route{
		"PutFleetsFleetIdMembersMemberId",
		"PUT",
		"/latest/fleets/{fleet_id}/members/{member_id}/",
		PutFleetsFleetIdMembersMemberId,
	},

	Route{
		"PutFleetsFleetIdSquadsSquadId",
		"PUT",
		"/latest/fleets/{fleet_id}/squads/{squad_id}/",
		PutFleetsFleetIdSquadsSquadId,
	},

	Route{
		"PutFleetsFleetIdWingsWingId",
		"PUT",
		"/latest/fleets/{fleet_id}/wings/{wing_id}/",
		PutFleetsFleetIdWingsWingId,
	},

	Route{
		"GetIncursions",
		"GET",
		"/latest/incursions/",
		GetIncursions,
	},

	Route{
		"GetIndustryFacilities",
		"GET",
		"/latest/industry/facilities/",
		GetIndustryFacilities,
	},

	Route{
		"GetIndustrySystems",
		"GET",
		"/latest/industry/systems/",
		GetIndustrySystems,
	},

	Route{
		"GetInsurancePrices",
		"GET",
		"/latest/insurance/prices/",
		GetInsurancePrices,
	},

	Route{
		"GetCharactersCharacterIdKillmailsRecent",
		"GET",
		"/latest/characters/{character_id}/killmails/recent/",
		GetCharactersCharacterIdKillmailsRecent,
	},

	Route{
		"GetKillmailsKillmailIdKillmailHash",
		"GET",
		"/latest/killmails/{killmail_id}/{killmail_hash}/",
		GetKillmailsKillmailIdKillmailHash,
	},

	Route{
		"GetCharactersCharacterIdLocation",
		"GET",
		"/latest/characters/{character_id}/location/",
		GetCharactersCharacterIdLocation,
	},

	Route{
		"GetCharactersCharacterIdShip",
		"GET",
		"/latest/characters/{character_id}/ship/",
		GetCharactersCharacterIdShip,
	},

	Route{
		"DeleteCharactersCharacterIdMailLabelsLabelId",
		"DELETE",
		"/latest/characters/{character_id}/mail/labels/{label_id}/",
		DeleteCharactersCharacterIdMailLabelsLabelId,
	},

	Route{
		"DeleteCharactersCharacterIdMailMailId",
		"DELETE",
		"/latest/characters/{character_id}/mail/{mail_id}/",
		DeleteCharactersCharacterIdMailMailId,
	},

	Route{
		"GetCharactersCharacterIdMail",
		"GET",
		"/latest/characters/{character_id}/mail/",
		GetCharactersCharacterIdMail,
	},

	Route{
		"GetCharactersCharacterIdMailLabels",
		"GET",
		"/latest/characters/{character_id}/mail/labels/",
		GetCharactersCharacterIdMailLabels,
	},

	Route{
		"GetCharactersCharacterIdMailLists",
		"GET",
		"/latest/characters/{character_id}/mail/lists/",
		GetCharactersCharacterIdMailLists,
	},

	Route{
		"GetCharactersCharacterIdMailMailId",
		"GET",
		"/latest/characters/{character_id}/mail/{mail_id}/",
		GetCharactersCharacterIdMailMailId,
	},

	Route{
		"PostCharactersCharacterIdMail",
		"POST",
		"/latest/characters/{character_id}/mail/",
		PostCharactersCharacterIdMail,
	},

	Route{
		"PostCharactersCharacterIdMailLabels",
		"POST",
		"/latest/characters/{character_id}/mail/labels/",
		PostCharactersCharacterIdMailLabels,
	},

	Route{
		"PutCharactersCharacterIdMailMailId",
		"PUT",
		"/latest/characters/{character_id}/mail/{mail_id}/",
		PutCharactersCharacterIdMailMailId,
	},

	Route{
		"GetMarketsPrices",
		"GET",
		"/latest/markets/prices/",
		GetMarketsPrices,
	},

	Route{
		"GetMarketsRegionIdHistory",
		"GET",
		"/latest/markets/{region_id}/history/",
		GetMarketsRegionIdHistory,
	},

	Route{
		"GetMarketsRegionIdOrders",
		"GET",
		"/latest/markets/{region_id}/orders/",
		GetMarketsRegionIdOrders,
	},

	Route{
		"GetCharactersCharacterIdPlanets",
		"GET",
		"/latest/characters/{character_id}/planets/",
		GetCharactersCharacterIdPlanets,
	},

	Route{
		"GetCharactersCharacterIdPlanetsPlanetId",
		"GET",
		"/latest/characters/{character_id}/planets/{planet_id}/",
		GetCharactersCharacterIdPlanetsPlanetId,
	},

	Route{
		"GetUniverseSchematicsSchematicId",
		"GET",
		"/latest/universe/schematics/{schematic_id}/",
		GetUniverseSchematicsSchematicId,
	},

	Route{
		"GetCharactersCharacterIdSearch",
		"GET",
		"/latest/characters/{character_id}/search/",
		GetCharactersCharacterIdSearch,
	},

	Route{
		"GetSearch",
		"GET",
		"/latest/search/",
		GetSearch,
	},

	Route{
		"GetCharactersCharacterIdSkillqueue",
		"GET",
		"/latest/characters/{character_id}/skillqueue/",
		GetCharactersCharacterIdSkillqueue,
	},

	Route{
		"GetCharactersCharacterIdSkills",
		"GET",
		"/latest/characters/{character_id}/skills/",
		GetCharactersCharacterIdSkills,
	},

	Route{
		"GetSovereigntyCampaigns",
		"GET",
		"/latest/sovereignty/campaigns/",
		GetSovereigntyCampaigns,
	},

	Route{
		"GetSovereigntyStructures",
		"GET",
		"/latest/sovereignty/structures/",
		GetSovereigntyStructures,
	},

	Route{
		"GetUniverseStationsStationId",
		"GET",
		"/latest/universe/stations/{station_id}/",
		GetUniverseStationsStationId,
	},

	Route{
		"GetUniverseStructures",
		"GET",
		"/latest/universe/structures/",
		GetUniverseStructures,
	},

	Route{
		"GetUniverseStructuresStructureId",
		"GET",
		"/latest/universe/structures/{structure_id}/",
		GetUniverseStructuresStructureId,
	},

	Route{
		"GetUniverseSystemsSystemId",
		"GET",
		"/latest/universe/systems/{system_id}/",
		GetUniverseSystemsSystemId,
	},

	Route{
		"GetUniverseTypesTypeId",
		"GET",
		"/latest/universe/types/{type_id}/",
		GetUniverseTypesTypeId,
	},

	Route{
		"PostUniverseNames",
		"POST",
		"/latest/universe/names/",
		PostUniverseNames,
	},

	Route{
		"PostUiAutopilotWaypoint",
		"POST",
		"/latest/ui/autopilot/waypoint/",
		PostUiAutopilotWaypoint,
	},

	Route{
		"PostUiOpenwindowContract",
		"POST",
		"/latest/ui/openwindow/contract/",
		PostUiOpenwindowContract,
	},

	Route{
		"PostUiOpenwindowInformation",
		"POST",
		"/latest/ui/openwindow/information/",
		PostUiOpenwindowInformation,
	},

	Route{
		"PostUiOpenwindowMarketdetails",
		"POST",
		"/latest/ui/openwindow/marketdetails/",
		PostUiOpenwindowMarketdetails,
	},

	Route{
		"PostUiOpenwindowNewmail",
		"POST",
		"/latest/ui/openwindow/newmail/",
		PostUiOpenwindowNewmail,
	},

	Route{
		"GetCharactersCharacterIdWallets",
		"GET",
		"/latest/characters/{character_id}/wallets/",
		GetCharactersCharacterIdWallets,
	},

	Route{
		"GetWars",
		"GET",
		"/latest/wars/",
		GetWars,
	},

	Route{
		"GetWarsWarId",
		"GET",
		"/latest/wars/{war_id}/",
		GetWarsWarId,
	},

	Route{
		"GetWarsWarIdKillmails",
		"GET",
		"/latest/wars/{war_id}/killmails/",
		GetWarsWarIdKillmails,
	},

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