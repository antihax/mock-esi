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
		"Get",
		"/dev/alliances/",
		GetAlliances,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceId",
		"Get",
		"/dev/alliances/{alliance_id}/",
		GetAlliancesAllianceId,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdCorporations",
		"Get",
		"/dev/alliances/{alliance_id}/corporations/",
		GetAlliancesAllianceIdCorporations,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdIcons",
		"Get",
		"/dev/alliances/{alliance_id}/icons/",
		GetAlliancesAllianceIdIcons,
	)

	mockesi.NewRoute(
		"GetAlliancesNames",
		"Get",
		"/dev/alliances/names/",
		GetAlliancesNames,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAssets",
		"Get",
		"/dev/characters/{character_id}/assets/",
		GetCharactersCharacterIdAssets,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdAssets",
		"Get",
		"/dev/corporations/{corporation_id}/assets/",
		GetCorporationsCorporationIdAssets,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdAssetsLocations",
		"Post",
		"/dev/characters/{character_id}/assets/locations/",
		PostCharactersCharacterIdAssetsLocations,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdAssetsNames",
		"Post",
		"/dev/characters/{character_id}/assets/names/",
		PostCharactersCharacterIdAssetsNames,
	)

	mockesi.NewRoute(
		"PostCorporationsCorporationIdAssetsLocations",
		"Post",
		"/dev/corporations/{corporation_id}/assets/locations/",
		PostCorporationsCorporationIdAssetsLocations,
	)

	mockesi.NewRoute(
		"PostCorporationsCorporationIdAssetsNames",
		"Post",
		"/dev/corporations/{corporation_id}/assets/names/",
		PostCorporationsCorporationIdAssetsNames,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBookmarks",
		"Get",
		"/dev/characters/{character_id}/bookmarks/",
		GetCharactersCharacterIdBookmarks,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBookmarksFolders",
		"Get",
		"/dev/characters/{character_id}/bookmarks/folders/",
		GetCharactersCharacterIdBookmarksFolders,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdBookmarks",
		"Get",
		"/dev/corporations/{corporation_id}/bookmarks/",
		GetCorporationsCorporationIdBookmarks,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdBookmarksFolders",
		"Get",
		"/dev/corporations/{corporation_id}/bookmarks/folders/",
		GetCorporationsCorporationIdBookmarksFolders,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCalendar",
		"Get",
		"/dev/characters/{character_id}/calendar/",
		GetCharactersCharacterIdCalendar,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCalendarEventId",
		"Get",
		"/dev/characters/{character_id}/calendar/{event_id}/",
		GetCharactersCharacterIdCalendarEventId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCalendarEventIdAttendees",
		"Get",
		"/dev/characters/{character_id}/calendar/{event_id}/attendees/",
		GetCharactersCharacterIdCalendarEventIdAttendees,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdCalendarEventId",
		"Put",
		"/dev/characters/{character_id}/calendar/{event_id}/",
		PutCharactersCharacterIdCalendarEventId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterId",
		"Get",
		"/dev/characters/{character_id}/",
		GetCharactersCharacterId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAgentsResearch",
		"Get",
		"/dev/characters/{character_id}/agents_research/",
		GetCharactersCharacterIdAgentsResearch,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBlueprints",
		"Get",
		"/dev/characters/{character_id}/blueprints/",
		GetCharactersCharacterIdBlueprints,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdChatChannels",
		"Get",
		"/dev/characters/{character_id}/chat_channels/",
		GetCharactersCharacterIdChatChannels,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCorporationhistory",
		"Get",
		"/dev/characters/{character_id}/corporationhistory/",
		GetCharactersCharacterIdCorporationhistory,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFatigue",
		"Get",
		"/dev/characters/{character_id}/fatigue/",
		GetCharactersCharacterIdFatigue,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMedals",
		"Get",
		"/dev/characters/{character_id}/medals/",
		GetCharactersCharacterIdMedals,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdNotifications",
		"Get",
		"/dev/characters/{character_id}/notifications/",
		GetCharactersCharacterIdNotifications,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdNotificationsContacts",
		"Get",
		"/dev/characters/{character_id}/notifications/contacts/",
		GetCharactersCharacterIdNotificationsContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPortrait",
		"Get",
		"/dev/characters/{character_id}/portrait/",
		GetCharactersCharacterIdPortrait,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdRoles",
		"Get",
		"/dev/characters/{character_id}/roles/",
		GetCharactersCharacterIdRoles,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdStandings",
		"Get",
		"/dev/characters/{character_id}/standings/",
		GetCharactersCharacterIdStandings,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdStats",
		"Get",
		"/dev/characters/{character_id}/stats/",
		GetCharactersCharacterIdStats,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdTitles",
		"Get",
		"/dev/characters/{character_id}/titles/",
		GetCharactersCharacterIdTitles,
	)

	mockesi.NewRoute(
		"GetCharactersNames",
		"Get",
		"/dev/characters/names/",
		GetCharactersNames,
	)

	mockesi.NewRoute(
		"PostCharactersAffiliation",
		"Post",
		"/dev/characters/affiliation/",
		PostCharactersAffiliation,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdCspa",
		"Post",
		"/dev/characters/{character_id}/cspa/",
		PostCharactersCharacterIdCspa,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdClones",
		"Get",
		"/dev/characters/{character_id}/clones/",
		GetCharactersCharacterIdClones,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdImplants",
		"Get",
		"/dev/characters/{character_id}/implants/",
		GetCharactersCharacterIdImplants,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdContacts",
		"Delete",
		"/dev/characters/{character_id}/contacts/",
		DeleteCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdContacts",
		"Get",
		"/dev/alliances/{alliance_id}/contacts/",
		GetAlliancesAllianceIdContacts,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdContactsLabels",
		"Get",
		"/dev/alliances/{alliance_id}/contacts/labels/",
		GetAlliancesAllianceIdContactsLabels,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContacts",
		"Get",
		"/dev/characters/{character_id}/contacts/",
		GetCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContactsLabels",
		"Get",
		"/dev/characters/{character_id}/contacts/labels/",
		GetCharactersCharacterIdContactsLabels,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContacts",
		"Get",
		"/dev/corporations/{corporation_id}/contacts/",
		GetCorporationsCorporationIdContacts,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContactsLabels",
		"Get",
		"/dev/corporations/{corporation_id}/contacts/labels/",
		GetCorporationsCorporationIdContactsLabels,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdContacts",
		"Post",
		"/dev/characters/{character_id}/contacts/",
		PostCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdContacts",
		"Put",
		"/dev/characters/{character_id}/contacts/",
		PutCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContracts",
		"Get",
		"/dev/characters/{character_id}/contracts/",
		GetCharactersCharacterIdContracts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContractsContractIdBids",
		"Get",
		"/dev/characters/{character_id}/contracts/{contract_id}/bids/",
		GetCharactersCharacterIdContractsContractIdBids,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContractsContractIdItems",
		"Get",
		"/dev/characters/{character_id}/contracts/{contract_id}/items/",
		GetCharactersCharacterIdContractsContractIdItems,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContracts",
		"Get",
		"/dev/corporations/{corporation_id}/contracts/",
		GetCorporationsCorporationIdContracts,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContractsContractIdBids",
		"Get",
		"/dev/corporations/{corporation_id}/contracts/{contract_id}/bids/",
		GetCorporationsCorporationIdContractsContractIdBids,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContractsContractIdItems",
		"Get",
		"/dev/corporations/{corporation_id}/contracts/{contract_id}/items/",
		GetCorporationsCorporationIdContractsContractIdItems,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationId",
		"Get",
		"/dev/corporations/{corporation_id}/",
		GetCorporationsCorporationId,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdAlliancehistory",
		"Get",
		"/dev/corporations/{corporation_id}/alliancehistory/",
		GetCorporationsCorporationIdAlliancehistory,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdBlueprints",
		"Get",
		"/dev/corporations/{corporation_id}/blueprints/",
		GetCorporationsCorporationIdBlueprints,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContainersLogs",
		"Get",
		"/dev/corporations/{corporation_id}/containers/logs/",
		GetCorporationsCorporationIdContainersLogs,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdDivisions",
		"Get",
		"/dev/corporations/{corporation_id}/divisions/",
		GetCorporationsCorporationIdDivisions,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdFacilities",
		"Get",
		"/dev/corporations/{corporation_id}/facilities/",
		GetCorporationsCorporationIdFacilities,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdIcons",
		"Get",
		"/dev/corporations/{corporation_id}/icons/",
		GetCorporationsCorporationIdIcons,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMedals",
		"Get",
		"/dev/corporations/{corporation_id}/medals/",
		GetCorporationsCorporationIdMedals,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMedalsIssued",
		"Get",
		"/dev/corporations/{corporation_id}/medals/issued/",
		GetCorporationsCorporationIdMedalsIssued,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMembers",
		"Get",
		"/dev/corporations/{corporation_id}/members/",
		GetCorporationsCorporationIdMembers,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMembersLimit",
		"Get",
		"/dev/corporations/{corporation_id}/members/limit/",
		GetCorporationsCorporationIdMembersLimit,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMembersTitles",
		"Get",
		"/dev/corporations/{corporation_id}/members/titles/",
		GetCorporationsCorporationIdMembersTitles,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMembertracking",
		"Get",
		"/dev/corporations/{corporation_id}/membertracking/",
		GetCorporationsCorporationIdMembertracking,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdOutposts",
		"Get",
		"/dev/corporations/{corporation_id}/outposts/",
		GetCorporationsCorporationIdOutposts,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdOutpostsOutpostId",
		"Get",
		"/dev/corporations/{corporation_id}/outposts/{outpost_id}/",
		GetCorporationsCorporationIdOutpostsOutpostId,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdRoles",
		"Get",
		"/dev/corporations/{corporation_id}/roles/",
		GetCorporationsCorporationIdRoles,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdRolesHistory",
		"Get",
		"/dev/corporations/{corporation_id}/roles/history/",
		GetCorporationsCorporationIdRolesHistory,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdShareholders",
		"Get",
		"/dev/corporations/{corporation_id}/shareholders/",
		GetCorporationsCorporationIdShareholders,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdStandings",
		"Get",
		"/dev/corporations/{corporation_id}/standings/",
		GetCorporationsCorporationIdStandings,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdStarbases",
		"Get",
		"/dev/corporations/{corporation_id}/starbases/",
		GetCorporationsCorporationIdStarbases,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdStarbasesStarbaseId",
		"Get",
		"/dev/corporations/{corporation_id}/starbases/{starbase_id}/",
		GetCorporationsCorporationIdStarbasesStarbaseId,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdStructures",
		"Get",
		"/dev/corporations/{corporation_id}/structures/",
		GetCorporationsCorporationIdStructures,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdTitles",
		"Get",
		"/dev/corporations/{corporation_id}/titles/",
		GetCorporationsCorporationIdTitles,
	)

	mockesi.NewRoute(
		"GetCorporationsNames",
		"Get",
		"/dev/corporations/names/",
		GetCorporationsNames,
	)

	mockesi.NewRoute(
		"GetCorporationsNpccorps",
		"Get",
		"/dev/corporations/npccorps/",
		GetCorporationsNpccorps,
	)

	mockesi.NewRoute(
		"GetDogmaAttributes",
		"Get",
		"/dev/dogma/attributes/",
		GetDogmaAttributes,
	)

	mockesi.NewRoute(
		"GetDogmaAttributesAttributeId",
		"Get",
		"/dev/dogma/attributes/{attribute_id}/",
		GetDogmaAttributesAttributeId,
	)

	mockesi.NewRoute(
		"GetDogmaEffects",
		"Get",
		"/dev/dogma/effects/",
		GetDogmaEffects,
	)

	mockesi.NewRoute(
		"GetDogmaEffectsEffectId",
		"Get",
		"/dev/dogma/effects/{effect_id}/",
		GetDogmaEffectsEffectId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFwStats",
		"Get",
		"/dev/characters/{character_id}/fw/stats/",
		GetCharactersCharacterIdFwStats,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdFwStats",
		"Get",
		"/dev/corporations/{corporation_id}/fw/stats/",
		GetCorporationsCorporationIdFwStats,
	)

	mockesi.NewRoute(
		"GetFwLeaderboards",
		"Get",
		"/dev/fw/leaderboards/",
		GetFwLeaderboards,
	)

	mockesi.NewRoute(
		"GetFwLeaderboardsCharacters",
		"Get",
		"/dev/fw/leaderboards/characters/",
		GetFwLeaderboardsCharacters,
	)

	mockesi.NewRoute(
		"GetFwLeaderboardsCorporations",
		"Get",
		"/dev/fw/leaderboards/corporations/",
		GetFwLeaderboardsCorporations,
	)

	mockesi.NewRoute(
		"GetFwStats",
		"Get",
		"/dev/fw/stats/",
		GetFwStats,
	)

	mockesi.NewRoute(
		"GetFwSystems",
		"Get",
		"/dev/fw/systems/",
		GetFwSystems,
	)

	mockesi.NewRoute(
		"GetFwWars",
		"Get",
		"/dev/fw/wars/",
		GetFwWars,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdFittingsFittingId",
		"Delete",
		"/dev/characters/{character_id}/fittings/{fitting_id}/",
		DeleteCharactersCharacterIdFittingsFittingId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFittings",
		"Get",
		"/dev/characters/{character_id}/fittings/",
		GetCharactersCharacterIdFittings,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdFittings",
		"Post",
		"/dev/characters/{character_id}/fittings/",
		PostCharactersCharacterIdFittings,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdMembersMemberId",
		"Delete",
		"/dev/fleets/{fleet_id}/members/{member_id}/",
		DeleteFleetsFleetIdMembersMemberId,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdSquadsSquadId",
		"Delete",
		"/dev/fleets/{fleet_id}/squads/{squad_id}/",
		DeleteFleetsFleetIdSquadsSquadId,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdWingsWingId",
		"Delete",
		"/dev/fleets/{fleet_id}/wings/{wing_id}/",
		DeleteFleetsFleetIdWingsWingId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFleet",
		"Get",
		"/dev/characters/{character_id}/fleet/",
		GetCharactersCharacterIdFleet,
	)

	mockesi.NewRoute(
		"GetFleetsFleetId",
		"Get",
		"/dev/fleets/{fleet_id}/",
		GetFleetsFleetId,
	)

	mockesi.NewRoute(
		"GetFleetsFleetIdMembers",
		"Get",
		"/dev/fleets/{fleet_id}/members/",
		GetFleetsFleetIdMembers,
	)

	mockesi.NewRoute(
		"GetFleetsFleetIdWings",
		"Get",
		"/dev/fleets/{fleet_id}/wings/",
		GetFleetsFleetIdWings,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdMembers",
		"Post",
		"/dev/fleets/{fleet_id}/members/",
		PostFleetsFleetIdMembers,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdWings",
		"Post",
		"/dev/fleets/{fleet_id}/wings/",
		PostFleetsFleetIdWings,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdWingsWingIdSquads",
		"Post",
		"/dev/fleets/{fleet_id}/wings/{wing_id}/squads/",
		PostFleetsFleetIdWingsWingIdSquads,
	)

	mockesi.NewRoute(
		"PutFleetsFleetId",
		"Put",
		"/dev/fleets/{fleet_id}/",
		PutFleetsFleetId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdMembersMemberId",
		"Put",
		"/dev/fleets/{fleet_id}/members/{member_id}/",
		PutFleetsFleetIdMembersMemberId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdSquadsSquadId",
		"Put",
		"/dev/fleets/{fleet_id}/squads/{squad_id}/",
		PutFleetsFleetIdSquadsSquadId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdWingsWingId",
		"Put",
		"/dev/fleets/{fleet_id}/wings/{wing_id}/",
		PutFleetsFleetIdWingsWingId,
	)

	mockesi.NewRoute(
		"GetIncursions",
		"Get",
		"/dev/incursions/",
		GetIncursions,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdIndustryJobs",
		"Get",
		"/dev/characters/{character_id}/industry/jobs/",
		GetCharactersCharacterIdIndustryJobs,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMining",
		"Get",
		"/dev/characters/{character_id}/mining/",
		GetCharactersCharacterIdMining,
	)

	mockesi.NewRoute(
		"GetCorporationCorporationIdMiningExtractions",
		"Get",
		"/dev/corporation/{corporation_id}/mining/extractions/",
		GetCorporationCorporationIdMiningExtractions,
	)

	mockesi.NewRoute(
		"GetCorporationCorporationIdMiningObservers",
		"Get",
		"/dev/corporation/{corporation_id}/mining/observers/",
		GetCorporationCorporationIdMiningObservers,
	)

	mockesi.NewRoute(
		"GetCorporationCorporationIdMiningObserversObserverId",
		"Get",
		"/dev/corporation/{corporation_id}/mining/observers/{observer_id}/",
		GetCorporationCorporationIdMiningObserversObserverId,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdIndustryJobs",
		"Get",
		"/dev/corporations/{corporation_id}/industry/jobs/",
		GetCorporationsCorporationIdIndustryJobs,
	)

	mockesi.NewRoute(
		"GetIndustryFacilities",
		"Get",
		"/dev/industry/facilities/",
		GetIndustryFacilities,
	)

	mockesi.NewRoute(
		"GetIndustrySystems",
		"Get",
		"/dev/industry/systems/",
		GetIndustrySystems,
	)

	mockesi.NewRoute(
		"GetInsurancePrices",
		"Get",
		"/dev/insurance/prices/",
		GetInsurancePrices,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdKillmailsRecent",
		"Get",
		"/dev/characters/{character_id}/killmails/recent/",
		GetCharactersCharacterIdKillmailsRecent,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdKillmailsRecent",
		"Get",
		"/dev/corporations/{corporation_id}/killmails/recent/",
		GetCorporationsCorporationIdKillmailsRecent,
	)

	mockesi.NewRoute(
		"GetKillmailsKillmailIdKillmailHash",
		"Get",
		"/dev/killmails/{killmail_id}/{killmail_hash}/",
		GetKillmailsKillmailIdKillmailHash,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdLocation",
		"Get",
		"/dev/characters/{character_id}/location/",
		GetCharactersCharacterIdLocation,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOnline",
		"Get",
		"/dev/characters/{character_id}/online/",
		GetCharactersCharacterIdOnline,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdShip",
		"Get",
		"/dev/characters/{character_id}/ship/",
		GetCharactersCharacterIdShip,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdLoyaltyPoints",
		"Get",
		"/dev/characters/{character_id}/loyalty/points/",
		GetCharactersCharacterIdLoyaltyPoints,
	)

	mockesi.NewRoute(
		"GetLoyaltyStoresCorporationIdOffers",
		"Get",
		"/dev/loyalty/stores/{corporation_id}/offers/",
		GetLoyaltyStoresCorporationIdOffers,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdMailLabelsLabelId",
		"Delete",
		"/dev/characters/{character_id}/mail/labels/{label_id}/",
		DeleteCharactersCharacterIdMailLabelsLabelId,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdMailMailId",
		"Delete",
		"/dev/characters/{character_id}/mail/{mail_id}/",
		DeleteCharactersCharacterIdMailMailId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMail",
		"Get",
		"/dev/characters/{character_id}/mail/",
		GetCharactersCharacterIdMail,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailLabels",
		"Get",
		"/dev/characters/{character_id}/mail/labels/",
		GetCharactersCharacterIdMailLabels,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailLists",
		"Get",
		"/dev/characters/{character_id}/mail/lists/",
		GetCharactersCharacterIdMailLists,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailMailId",
		"Get",
		"/dev/characters/{character_id}/mail/{mail_id}/",
		GetCharactersCharacterIdMailMailId,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdMail",
		"Post",
		"/dev/characters/{character_id}/mail/",
		PostCharactersCharacterIdMail,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdMailLabels",
		"Post",
		"/dev/characters/{character_id}/mail/labels/",
		PostCharactersCharacterIdMailLabels,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdMailMailId",
		"Put",
		"/dev/characters/{character_id}/mail/{mail_id}/",
		PutCharactersCharacterIdMailMailId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOrders",
		"Get",
		"/dev/characters/{character_id}/orders/",
		GetCharactersCharacterIdOrders,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOrdersHistory",
		"Get",
		"/dev/characters/{character_id}/orders/history/",
		GetCharactersCharacterIdOrdersHistory,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdOrders",
		"Get",
		"/dev/corporations/{corporation_id}/orders/",
		GetCorporationsCorporationIdOrders,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdOrdersHistory",
		"Get",
		"/dev/corporations/{corporation_id}/orders/history/",
		GetCorporationsCorporationIdOrdersHistory,
	)

	mockesi.NewRoute(
		"GetMarketsGroups",
		"Get",
		"/dev/markets/groups/",
		GetMarketsGroups,
	)

	mockesi.NewRoute(
		"GetMarketsGroupsMarketGroupId",
		"Get",
		"/dev/markets/groups/{market_group_id}/",
		GetMarketsGroupsMarketGroupId,
	)

	mockesi.NewRoute(
		"GetMarketsPrices",
		"Get",
		"/dev/markets/prices/",
		GetMarketsPrices,
	)

	mockesi.NewRoute(
		"GetMarketsRegionIdHistory",
		"Get",
		"/dev/markets/{region_id}/history/",
		GetMarketsRegionIdHistory,
	)

	mockesi.NewRoute(
		"GetMarketsRegionIdOrders",
		"Get",
		"/dev/markets/{region_id}/orders/",
		GetMarketsRegionIdOrders,
	)

	mockesi.NewRoute(
		"GetMarketsRegionIdTypes",
		"Get",
		"/dev/markets/{region_id}/types/",
		GetMarketsRegionIdTypes,
	)

	mockesi.NewRoute(
		"GetMarketsStructuresStructureId",
		"Get",
		"/dev/markets/structures/{structure_id}/",
		GetMarketsStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOpportunities",
		"Get",
		"/dev/characters/{character_id}/opportunities/",
		GetCharactersCharacterIdOpportunities,
	)

	mockesi.NewRoute(
		"GetOpportunitiesGroups",
		"Get",
		"/dev/opportunities/groups/",
		GetOpportunitiesGroups,
	)

	mockesi.NewRoute(
		"GetOpportunitiesGroupsGroupId",
		"Get",
		"/dev/opportunities/groups/{group_id}/",
		GetOpportunitiesGroupsGroupId,
	)

	mockesi.NewRoute(
		"GetOpportunitiesTasks",
		"Get",
		"/dev/opportunities/tasks/",
		GetOpportunitiesTasks,
	)

	mockesi.NewRoute(
		"GetOpportunitiesTasksTaskId",
		"Get",
		"/dev/opportunities/tasks/{task_id}/",
		GetOpportunitiesTasksTaskId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPlanets",
		"Get",
		"/dev/characters/{character_id}/planets/",
		GetCharactersCharacterIdPlanets,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPlanetsPlanetId",
		"Get",
		"/dev/characters/{character_id}/planets/{planet_id}/",
		GetCharactersCharacterIdPlanetsPlanetId,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdCustomsOffices",
		"Get",
		"/dev/corporations/{corporation_id}/customs_offices/",
		GetCorporationsCorporationIdCustomsOffices,
	)

	mockesi.NewRoute(
		"GetUniverseSchematicsSchematicId",
		"Get",
		"/dev/universe/schematics/{schematic_id}/",
		GetUniverseSchematicsSchematicId,
	)

	mockesi.NewRoute(
		"GetRouteOriginDestination",
		"Get",
		"/dev/route/{origin}/{destination}/",
		GetRouteOriginDestination,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdSearch",
		"Get",
		"/dev/characters/{character_id}/search/",
		GetCharactersCharacterIdSearch,
	)

	mockesi.NewRoute(
		"GetSearch",
		"Get",
		"/dev/search/",
		GetSearch,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAttributes",
		"Get",
		"/dev/characters/{character_id}/attributes/",
		GetCharactersCharacterIdAttributes,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdSkillqueue",
		"Get",
		"/dev/characters/{character_id}/skillqueue/",
		GetCharactersCharacterIdSkillqueue,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdSkills",
		"Get",
		"/dev/characters/{character_id}/skills/",
		GetCharactersCharacterIdSkills,
	)

	mockesi.NewRoute(
		"GetSovereigntyCampaigns",
		"Get",
		"/dev/sovereignty/campaigns/",
		GetSovereigntyCampaigns,
	)

	mockesi.NewRoute(
		"GetSovereigntyMap",
		"Get",
		"/dev/sovereignty/map/",
		GetSovereigntyMap,
	)

	mockesi.NewRoute(
		"GetSovereigntyStructures",
		"Get",
		"/dev/sovereignty/structures/",
		GetSovereigntyStructures,
	)

	mockesi.NewRoute(
		"GetStatus",
		"Get",
		"/dev/status/",
		GetStatus,
	)

	mockesi.NewRoute(
		"GetUniverseAncestries",
		"Get",
		"/dev/universe/ancestries/",
		GetUniverseAncestries,
	)

	mockesi.NewRoute(
		"GetUniverseAsteroidBeltsAsteroidBeltId",
		"Get",
		"/dev/universe/asteroid_belts/{asteroid_belt_id}/",
		GetUniverseAsteroidBeltsAsteroidBeltId,
	)

	mockesi.NewRoute(
		"GetUniverseBloodlines",
		"Get",
		"/dev/universe/bloodlines/",
		GetUniverseBloodlines,
	)

	mockesi.NewRoute(
		"GetUniverseCategories",
		"Get",
		"/dev/universe/categories/",
		GetUniverseCategories,
	)

	mockesi.NewRoute(
		"GetUniverseCategoriesCategoryId",
		"Get",
		"/dev/universe/categories/{category_id}/",
		GetUniverseCategoriesCategoryId,
	)

	mockesi.NewRoute(
		"GetUniverseConstellations",
		"Get",
		"/dev/universe/constellations/",
		GetUniverseConstellations,
	)

	mockesi.NewRoute(
		"GetUniverseConstellationsConstellationId",
		"Get",
		"/dev/universe/constellations/{constellation_id}/",
		GetUniverseConstellationsConstellationId,
	)

	mockesi.NewRoute(
		"GetUniverseFactions",
		"Get",
		"/dev/universe/factions/",
		GetUniverseFactions,
	)

	mockesi.NewRoute(
		"GetUniverseGraphics",
		"Get",
		"/dev/universe/graphics/",
		GetUniverseGraphics,
	)

	mockesi.NewRoute(
		"GetUniverseGraphicsGraphicId",
		"Get",
		"/dev/universe/graphics/{graphic_id}/",
		GetUniverseGraphicsGraphicId,
	)

	mockesi.NewRoute(
		"GetUniverseGroups",
		"Get",
		"/dev/universe/groups/",
		GetUniverseGroups,
	)

	mockesi.NewRoute(
		"GetUniverseGroupsGroupId",
		"Get",
		"/dev/universe/groups/{group_id}/",
		GetUniverseGroupsGroupId,
	)

	mockesi.NewRoute(
		"GetUniverseMoonsMoonId",
		"Get",
		"/dev/universe/moons/{moon_id}/",
		GetUniverseMoonsMoonId,
	)

	mockesi.NewRoute(
		"GetUniversePlanetsPlanetId",
		"Get",
		"/dev/universe/planets/{planet_id}/",
		GetUniversePlanetsPlanetId,
	)

	mockesi.NewRoute(
		"GetUniverseRaces",
		"Get",
		"/dev/universe/races/",
		GetUniverseRaces,
	)

	mockesi.NewRoute(
		"GetUniverseRegions",
		"Get",
		"/dev/universe/regions/",
		GetUniverseRegions,
	)

	mockesi.NewRoute(
		"GetUniverseRegionsRegionId",
		"Get",
		"/dev/universe/regions/{region_id}/",
		GetUniverseRegionsRegionId,
	)

	mockesi.NewRoute(
		"GetUniverseStargatesStargateId",
		"Get",
		"/dev/universe/stargates/{stargate_id}/",
		GetUniverseStargatesStargateId,
	)

	mockesi.NewRoute(
		"GetUniverseStarsStarId",
		"Get",
		"/dev/universe/stars/{star_id}/",
		GetUniverseStarsStarId,
	)

	mockesi.NewRoute(
		"GetUniverseStationsStationId",
		"Get",
		"/dev/universe/stations/{station_id}/",
		GetUniverseStationsStationId,
	)

	mockesi.NewRoute(
		"GetUniverseStructures",
		"Get",
		"/dev/universe/structures/",
		GetUniverseStructures,
	)

	mockesi.NewRoute(
		"GetUniverseStructuresStructureId",
		"Get",
		"/dev/universe/structures/{structure_id}/",
		GetUniverseStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetUniverseSystemJumps",
		"Get",
		"/dev/universe/system_jumps/",
		GetUniverseSystemJumps,
	)

	mockesi.NewRoute(
		"GetUniverseSystemKills",
		"Get",
		"/dev/universe/system_kills/",
		GetUniverseSystemKills,
	)

	mockesi.NewRoute(
		"GetUniverseSystems",
		"Get",
		"/dev/universe/systems/",
		GetUniverseSystems,
	)

	mockesi.NewRoute(
		"GetUniverseSystemsSystemId",
		"Get",
		"/dev/universe/systems/{system_id}/",
		GetUniverseSystemsSystemId,
	)

	mockesi.NewRoute(
		"GetUniverseTypes",
		"Get",
		"/dev/universe/types/",
		GetUniverseTypes,
	)

	mockesi.NewRoute(
		"GetUniverseTypesTypeId",
		"Get",
		"/dev/universe/types/{type_id}/",
		GetUniverseTypesTypeId,
	)

	mockesi.NewRoute(
		"PostUniverseIds",
		"Post",
		"/dev/universe/ids/",
		PostUniverseIds,
	)

	mockesi.NewRoute(
		"PostUniverseNames",
		"Post",
		"/dev/universe/names/",
		PostUniverseNames,
	)

	mockesi.NewRoute(
		"PostUiAutopilotWaypoint",
		"Post",
		"/dev/ui/autopilot/waypoint/",
		PostUiAutopilotWaypoint,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowContract",
		"Post",
		"/dev/ui/openwindow/contract/",
		PostUiOpenwindowContract,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowInformation",
		"Post",
		"/dev/ui/openwindow/information/",
		PostUiOpenwindowInformation,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowMarketdetails",
		"Post",
		"/dev/ui/openwindow/marketdetails/",
		PostUiOpenwindowMarketdetails,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowNewmail",
		"Post",
		"/dev/ui/openwindow/newmail/",
		PostUiOpenwindowNewmail,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdWallet",
		"Get",
		"/dev/characters/{character_id}/wallet/",
		GetCharactersCharacterIdWallet,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdWalletJournal",
		"Get",
		"/dev/characters/{character_id}/wallet/journal/",
		GetCharactersCharacterIdWalletJournal,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdWalletTransactions",
		"Get",
		"/dev/characters/{character_id}/wallet/transactions/",
		GetCharactersCharacterIdWalletTransactions,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdWallets",
		"Get",
		"/dev/corporations/{corporation_id}/wallets/",
		GetCorporationsCorporationIdWallets,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdWalletsDivisionJournal",
		"Get",
		"/dev/corporations/{corporation_id}/wallets/{division}/journal/",
		GetCorporationsCorporationIdWalletsDivisionJournal,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdWalletsDivisionTransactions",
		"Get",
		"/dev/corporations/{corporation_id}/wallets/{division}/transactions/",
		GetCorporationsCorporationIdWalletsDivisionTransactions,
	)

	mockesi.NewRoute(
		"GetWars",
		"Get",
		"/dev/wars/",
		GetWars,
	)

	mockesi.NewRoute(
		"GetWarsWarId",
		"Get",
		"/dev/wars/{war_id}/",
		GetWarsWarId,
	)

	mockesi.NewRoute(
		"GetWarsWarIdKillmails",
		"Get",
		"/dev/wars/{war_id}/killmails/",
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
