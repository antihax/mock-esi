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
		"Get",
		"/legacy/alliances/",
		GetAlliances,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceId",
		"Get",
		"/legacy/alliances/{alliance_id}/",
		GetAlliancesAllianceId,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdCorporations",
		"Get",
		"/legacy/alliances/{alliance_id}/corporations/",
		GetAlliancesAllianceIdCorporations,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdIcons",
		"Get",
		"/legacy/alliances/{alliance_id}/icons/",
		GetAlliancesAllianceIdIcons,
	)

	mockesi.NewRoute(
		"GetAlliancesNames",
		"Get",
		"/legacy/alliances/names/",
		GetAlliancesNames,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAssets",
		"Get",
		"/legacy/characters/{character_id}/assets/",
		GetCharactersCharacterIdAssets,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdAssets",
		"Get",
		"/legacy/corporations/{corporation_id}/assets/",
		GetCorporationsCorporationIdAssets,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdAssetsLocations",
		"Post",
		"/legacy/characters/{character_id}/assets/locations/",
		PostCharactersCharacterIdAssetsLocations,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdAssetsNames",
		"Post",
		"/legacy/characters/{character_id}/assets/names/",
		PostCharactersCharacterIdAssetsNames,
	)

	mockesi.NewRoute(
		"PostCorporationsCorporationIdAssetsLocations",
		"Post",
		"/legacy/corporations/{corporation_id}/assets/locations/",
		PostCorporationsCorporationIdAssetsLocations,
	)

	mockesi.NewRoute(
		"PostCorporationsCorporationIdAssetsNames",
		"Post",
		"/legacy/corporations/{corporation_id}/assets/names/",
		PostCorporationsCorporationIdAssetsNames,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBookmarks",
		"Get",
		"/legacy/characters/{character_id}/bookmarks/",
		GetCharactersCharacterIdBookmarks,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBookmarksFolders",
		"Get",
		"/legacy/characters/{character_id}/bookmarks/folders/",
		GetCharactersCharacterIdBookmarksFolders,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdBookmarks",
		"Get",
		"/legacy/corporations/{corporation_id}/bookmarks/",
		GetCorporationsCorporationIdBookmarks,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdBookmarksFolders",
		"Get",
		"/legacy/corporations/{corporation_id}/bookmarks/folders/",
		GetCorporationsCorporationIdBookmarksFolders,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCalendar",
		"Get",
		"/legacy/characters/{character_id}/calendar/",
		GetCharactersCharacterIdCalendar,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCalendarEventId",
		"Get",
		"/legacy/characters/{character_id}/calendar/{event_id}/",
		GetCharactersCharacterIdCalendarEventId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCalendarEventIdAttendees",
		"Get",
		"/legacy/characters/{character_id}/calendar/{event_id}/attendees/",
		GetCharactersCharacterIdCalendarEventIdAttendees,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdCalendarEventId",
		"Put",
		"/legacy/characters/{character_id}/calendar/{event_id}/",
		PutCharactersCharacterIdCalendarEventId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterId",
		"Get",
		"/legacy/characters/{character_id}/",
		GetCharactersCharacterId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAgentsResearch",
		"Get",
		"/legacy/characters/{character_id}/agents_research/",
		GetCharactersCharacterIdAgentsResearch,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBlueprints",
		"Get",
		"/legacy/characters/{character_id}/blueprints/",
		GetCharactersCharacterIdBlueprints,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdChatChannels",
		"Get",
		"/legacy/characters/{character_id}/chat_channels/",
		GetCharactersCharacterIdChatChannels,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCorporationhistory",
		"Get",
		"/legacy/characters/{character_id}/corporationhistory/",
		GetCharactersCharacterIdCorporationhistory,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFatigue",
		"Get",
		"/legacy/characters/{character_id}/fatigue/",
		GetCharactersCharacterIdFatigue,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMedals",
		"Get",
		"/legacy/characters/{character_id}/medals/",
		GetCharactersCharacterIdMedals,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdNotifications",
		"Get",
		"/legacy/characters/{character_id}/notifications/",
		GetCharactersCharacterIdNotifications,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdNotificationsContacts",
		"Get",
		"/legacy/characters/{character_id}/notifications/contacts/",
		GetCharactersCharacterIdNotificationsContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPortrait",
		"Get",
		"/legacy/characters/{character_id}/portrait/",
		GetCharactersCharacterIdPortrait,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdRoles",
		"Get",
		"/legacy/characters/{character_id}/roles/",
		GetCharactersCharacterIdRoles,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdStandings",
		"Get",
		"/legacy/characters/{character_id}/standings/",
		GetCharactersCharacterIdStandings,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdStats",
		"Get",
		"/legacy/characters/{character_id}/stats/",
		GetCharactersCharacterIdStats,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdTitles",
		"Get",
		"/legacy/characters/{character_id}/titles/",
		GetCharactersCharacterIdTitles,
	)

	mockesi.NewRoute(
		"GetCharactersNames",
		"Get",
		"/legacy/characters/names/",
		GetCharactersNames,
	)

	mockesi.NewRoute(
		"PostCharactersAffiliation",
		"Post",
		"/legacy/characters/affiliation/",
		PostCharactersAffiliation,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdCspa",
		"Post",
		"/legacy/characters/{character_id}/cspa/",
		PostCharactersCharacterIdCspa,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdClones",
		"Get",
		"/legacy/characters/{character_id}/clones/",
		GetCharactersCharacterIdClones,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdImplants",
		"Get",
		"/legacy/characters/{character_id}/implants/",
		GetCharactersCharacterIdImplants,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdContacts",
		"Delete",
		"/legacy/characters/{character_id}/contacts/",
		DeleteCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdContacts",
		"Get",
		"/legacy/alliances/{alliance_id}/contacts/",
		GetAlliancesAllianceIdContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContacts",
		"Get",
		"/legacy/characters/{character_id}/contacts/",
		GetCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContactsLabels",
		"Get",
		"/legacy/characters/{character_id}/contacts/labels/",
		GetCharactersCharacterIdContactsLabels,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContacts",
		"Get",
		"/legacy/corporations/{corporation_id}/contacts/",
		GetCorporationsCorporationIdContacts,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdContacts",
		"Post",
		"/legacy/characters/{character_id}/contacts/",
		PostCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdContacts",
		"Put",
		"/legacy/characters/{character_id}/contacts/",
		PutCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContracts",
		"Get",
		"/legacy/characters/{character_id}/contracts/",
		GetCharactersCharacterIdContracts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContractsContractIdBids",
		"Get",
		"/legacy/characters/{character_id}/contracts/{contract_id}/bids/",
		GetCharactersCharacterIdContractsContractIdBids,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContractsContractIdItems",
		"Get",
		"/legacy/characters/{character_id}/contracts/{contract_id}/items/",
		GetCharactersCharacterIdContractsContractIdItems,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContracts",
		"Get",
		"/legacy/corporations/{corporation_id}/contracts/",
		GetCorporationsCorporationIdContracts,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContractsContractIdBids",
		"Get",
		"/legacy/corporations/{corporation_id}/contracts/{contract_id}/bids/",
		GetCorporationsCorporationIdContractsContractIdBids,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContractsContractIdItems",
		"Get",
		"/legacy/corporations/{corporation_id}/contracts/{contract_id}/items/",
		GetCorporationsCorporationIdContractsContractIdItems,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationId",
		"Get",
		"/legacy/corporations/{corporation_id}/",
		GetCorporationsCorporationId,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdAlliancehistory",
		"Get",
		"/legacy/corporations/{corporation_id}/alliancehistory/",
		GetCorporationsCorporationIdAlliancehistory,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdBlueprints",
		"Get",
		"/legacy/corporations/{corporation_id}/blueprints/",
		GetCorporationsCorporationIdBlueprints,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContainersLogs",
		"Get",
		"/legacy/corporations/{corporation_id}/containers/logs/",
		GetCorporationsCorporationIdContainersLogs,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdDivisions",
		"Get",
		"/legacy/corporations/{corporation_id}/divisions/",
		GetCorporationsCorporationIdDivisions,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdFacilities",
		"Get",
		"/legacy/corporations/{corporation_id}/facilities/",
		GetCorporationsCorporationIdFacilities,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdIcons",
		"Get",
		"/legacy/corporations/{corporation_id}/icons/",
		GetCorporationsCorporationIdIcons,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMedals",
		"Get",
		"/legacy/corporations/{corporation_id}/medals/",
		GetCorporationsCorporationIdMedals,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMedalsIssued",
		"Get",
		"/legacy/corporations/{corporation_id}/medals/issued/",
		GetCorporationsCorporationIdMedalsIssued,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMembers",
		"Get",
		"/legacy/corporations/{corporation_id}/members/",
		GetCorporationsCorporationIdMembers,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMembersLimit",
		"Get",
		"/legacy/corporations/{corporation_id}/members/limit/",
		GetCorporationsCorporationIdMembersLimit,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMembersTitles",
		"Get",
		"/legacy/corporations/{corporation_id}/members/titles/",
		GetCorporationsCorporationIdMembersTitles,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMembertracking",
		"Get",
		"/legacy/corporations/{corporation_id}/membertracking/",
		GetCorporationsCorporationIdMembertracking,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdOutposts",
		"Get",
		"/legacy/corporations/{corporation_id}/outposts/",
		GetCorporationsCorporationIdOutposts,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdOutpostsOutpostId",
		"Get",
		"/legacy/corporations/{corporation_id}/outposts/{outpost_id}/",
		GetCorporationsCorporationIdOutpostsOutpostId,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdRoles",
		"Get",
		"/legacy/corporations/{corporation_id}/roles/",
		GetCorporationsCorporationIdRoles,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdRolesHistory",
		"Get",
		"/legacy/corporations/{corporation_id}/roles/history/",
		GetCorporationsCorporationIdRolesHistory,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdShareholders",
		"Get",
		"/legacy/corporations/{corporation_id}/shareholders/",
		GetCorporationsCorporationIdShareholders,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdStandings",
		"Get",
		"/legacy/corporations/{corporation_id}/standings/",
		GetCorporationsCorporationIdStandings,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdStarbases",
		"Get",
		"/legacy/corporations/{corporation_id}/starbases/",
		GetCorporationsCorporationIdStarbases,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdStarbasesStarbaseId",
		"Get",
		"/legacy/corporations/{corporation_id}/starbases/{starbase_id}/",
		GetCorporationsCorporationIdStarbasesStarbaseId,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdStructures",
		"Get",
		"/legacy/corporations/{corporation_id}/structures/",
		GetCorporationsCorporationIdStructures,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdTitles",
		"Get",
		"/legacy/corporations/{corporation_id}/titles/",
		GetCorporationsCorporationIdTitles,
	)

	mockesi.NewRoute(
		"GetCorporationsNames",
		"Get",
		"/legacy/corporations/names/",
		GetCorporationsNames,
	)

	mockesi.NewRoute(
		"GetCorporationsNpccorps",
		"Get",
		"/legacy/corporations/npccorps/",
		GetCorporationsNpccorps,
	)

	mockesi.NewRoute(
		"PutCorporationsCorporationIdStructuresStructureId",
		"Put",
		"/legacy/corporations/{corporation_id}/structures/{structure_id}/",
		PutCorporationsCorporationIdStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetDogmaAttributes",
		"Get",
		"/legacy/dogma/attributes/",
		GetDogmaAttributes,
	)

	mockesi.NewRoute(
		"GetDogmaAttributesAttributeId",
		"Get",
		"/legacy/dogma/attributes/{attribute_id}/",
		GetDogmaAttributesAttributeId,
	)

	mockesi.NewRoute(
		"GetDogmaEffects",
		"Get",
		"/legacy/dogma/effects/",
		GetDogmaEffects,
	)

	mockesi.NewRoute(
		"GetDogmaEffectsEffectId",
		"Get",
		"/legacy/dogma/effects/{effect_id}/",
		GetDogmaEffectsEffectId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFwStats",
		"Get",
		"/legacy/characters/{character_id}/fw/stats/",
		GetCharactersCharacterIdFwStats,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdFwStats",
		"Get",
		"/legacy/corporations/{corporation_id}/fw/stats/",
		GetCorporationsCorporationIdFwStats,
	)

	mockesi.NewRoute(
		"GetFwLeaderboards",
		"Get",
		"/legacy/fw/leaderboards/",
		GetFwLeaderboards,
	)

	mockesi.NewRoute(
		"GetFwLeaderboardsCharacters",
		"Get",
		"/legacy/fw/leaderboards/characters/",
		GetFwLeaderboardsCharacters,
	)

	mockesi.NewRoute(
		"GetFwLeaderboardsCorporations",
		"Get",
		"/legacy/fw/leaderboards/corporations/",
		GetFwLeaderboardsCorporations,
	)

	mockesi.NewRoute(
		"GetFwStats",
		"Get",
		"/legacy/fw/stats/",
		GetFwStats,
	)

	mockesi.NewRoute(
		"GetFwSystems",
		"Get",
		"/legacy/fw/systems/",
		GetFwSystems,
	)

	mockesi.NewRoute(
		"GetFwWars",
		"Get",
		"/legacy/fw/wars/",
		GetFwWars,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdFittingsFittingId",
		"Delete",
		"/legacy/characters/{character_id}/fittings/{fitting_id}/",
		DeleteCharactersCharacterIdFittingsFittingId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFittings",
		"Get",
		"/legacy/characters/{character_id}/fittings/",
		GetCharactersCharacterIdFittings,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdFittings",
		"Post",
		"/legacy/characters/{character_id}/fittings/",
		PostCharactersCharacterIdFittings,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdMembersMemberId",
		"Delete",
		"/legacy/fleets/{fleet_id}/members/{member_id}/",
		DeleteFleetsFleetIdMembersMemberId,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdSquadsSquadId",
		"Delete",
		"/legacy/fleets/{fleet_id}/squads/{squad_id}/",
		DeleteFleetsFleetIdSquadsSquadId,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdWingsWingId",
		"Delete",
		"/legacy/fleets/{fleet_id}/wings/{wing_id}/",
		DeleteFleetsFleetIdWingsWingId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFleet",
		"Get",
		"/legacy/characters/{character_id}/fleet/",
		GetCharactersCharacterIdFleet,
	)

	mockesi.NewRoute(
		"GetFleetsFleetId",
		"Get",
		"/legacy/fleets/{fleet_id}/",
		GetFleetsFleetId,
	)

	mockesi.NewRoute(
		"GetFleetsFleetIdMembers",
		"Get",
		"/legacy/fleets/{fleet_id}/members/",
		GetFleetsFleetIdMembers,
	)

	mockesi.NewRoute(
		"GetFleetsFleetIdWings",
		"Get",
		"/legacy/fleets/{fleet_id}/wings/",
		GetFleetsFleetIdWings,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdMembers",
		"Post",
		"/legacy/fleets/{fleet_id}/members/",
		PostFleetsFleetIdMembers,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdWings",
		"Post",
		"/legacy/fleets/{fleet_id}/wings/",
		PostFleetsFleetIdWings,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdWingsWingIdSquads",
		"Post",
		"/legacy/fleets/{fleet_id}/wings/{wing_id}/squads/",
		PostFleetsFleetIdWingsWingIdSquads,
	)

	mockesi.NewRoute(
		"PutFleetsFleetId",
		"Put",
		"/legacy/fleets/{fleet_id}/",
		PutFleetsFleetId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdMembersMemberId",
		"Put",
		"/legacy/fleets/{fleet_id}/members/{member_id}/",
		PutFleetsFleetIdMembersMemberId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdSquadsSquadId",
		"Put",
		"/legacy/fleets/{fleet_id}/squads/{squad_id}/",
		PutFleetsFleetIdSquadsSquadId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdWingsWingId",
		"Put",
		"/legacy/fleets/{fleet_id}/wings/{wing_id}/",
		PutFleetsFleetIdWingsWingId,
	)

	mockesi.NewRoute(
		"GetIncursions",
		"Get",
		"/legacy/incursions/",
		GetIncursions,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdIndustryJobs",
		"Get",
		"/legacy/characters/{character_id}/industry/jobs/",
		GetCharactersCharacterIdIndustryJobs,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMining",
		"Get",
		"/legacy/characters/{character_id}/mining/",
		GetCharactersCharacterIdMining,
	)

	mockesi.NewRoute(
		"GetCorporationCorporationIdMiningExtractions",
		"Get",
		"/legacy/corporation/{corporation_id}/mining/extractions/",
		GetCorporationCorporationIdMiningExtractions,
	)

	mockesi.NewRoute(
		"GetCorporationCorporationIdMiningObservers",
		"Get",
		"/legacy/corporation/{corporation_id}/mining/observers/",
		GetCorporationCorporationIdMiningObservers,
	)

	mockesi.NewRoute(
		"GetCorporationCorporationIdMiningObserversObserverId",
		"Get",
		"/legacy/corporation/{corporation_id}/mining/observers/{observer_id}/",
		GetCorporationCorporationIdMiningObserversObserverId,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdIndustryJobs",
		"Get",
		"/legacy/corporations/{corporation_id}/industry/jobs/",
		GetCorporationsCorporationIdIndustryJobs,
	)

	mockesi.NewRoute(
		"GetIndustryFacilities",
		"Get",
		"/legacy/industry/facilities/",
		GetIndustryFacilities,
	)

	mockesi.NewRoute(
		"GetIndustrySystems",
		"Get",
		"/legacy/industry/systems/",
		GetIndustrySystems,
	)

	mockesi.NewRoute(
		"GetInsurancePrices",
		"Get",
		"/legacy/insurance/prices/",
		GetInsurancePrices,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdKillmailsRecent",
		"Get",
		"/legacy/characters/{character_id}/killmails/recent/",
		GetCharactersCharacterIdKillmailsRecent,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdKillmailsRecent",
		"Get",
		"/legacy/corporations/{corporation_id}/killmails/recent/",
		GetCorporationsCorporationIdKillmailsRecent,
	)

	mockesi.NewRoute(
		"GetKillmailsKillmailIdKillmailHash",
		"Get",
		"/legacy/killmails/{killmail_id}/{killmail_hash}/",
		GetKillmailsKillmailIdKillmailHash,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdLocation",
		"Get",
		"/legacy/characters/{character_id}/location/",
		GetCharactersCharacterIdLocation,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOnline",
		"Get",
		"/legacy/characters/{character_id}/online/",
		GetCharactersCharacterIdOnline,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdShip",
		"Get",
		"/legacy/characters/{character_id}/ship/",
		GetCharactersCharacterIdShip,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdLoyaltyPoints",
		"Get",
		"/legacy/characters/{character_id}/loyalty/points/",
		GetCharactersCharacterIdLoyaltyPoints,
	)

	mockesi.NewRoute(
		"GetLoyaltyStoresCorporationIdOffers",
		"Get",
		"/legacy/loyalty/stores/{corporation_id}/offers/",
		GetLoyaltyStoresCorporationIdOffers,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdMailLabelsLabelId",
		"Delete",
		"/legacy/characters/{character_id}/mail/labels/{label_id}/",
		DeleteCharactersCharacterIdMailLabelsLabelId,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdMailMailId",
		"Delete",
		"/legacy/characters/{character_id}/mail/{mail_id}/",
		DeleteCharactersCharacterIdMailMailId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMail",
		"Get",
		"/legacy/characters/{character_id}/mail/",
		GetCharactersCharacterIdMail,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailLabels",
		"Get",
		"/legacy/characters/{character_id}/mail/labels/",
		GetCharactersCharacterIdMailLabels,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailLists",
		"Get",
		"/legacy/characters/{character_id}/mail/lists/",
		GetCharactersCharacterIdMailLists,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailMailId",
		"Get",
		"/legacy/characters/{character_id}/mail/{mail_id}/",
		GetCharactersCharacterIdMailMailId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailUnread",
		"Get",
		"/legacy/characters/{character_id}/mail/unread/",
		GetCharactersCharacterIdMailUnread,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdMail",
		"Post",
		"/legacy/characters/{character_id}/mail/",
		PostCharactersCharacterIdMail,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdMailLabels",
		"Post",
		"/legacy/characters/{character_id}/mail/labels/",
		PostCharactersCharacterIdMailLabels,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdMailMailId",
		"Put",
		"/legacy/characters/{character_id}/mail/{mail_id}/",
		PutCharactersCharacterIdMailMailId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOrders",
		"Get",
		"/legacy/characters/{character_id}/orders/",
		GetCharactersCharacterIdOrders,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOrdersHistory",
		"Get",
		"/legacy/characters/{character_id}/orders/history/",
		GetCharactersCharacterIdOrdersHistory,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdOrders",
		"Get",
		"/legacy/corporations/{corporation_id}/orders/",
		GetCorporationsCorporationIdOrders,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdOrdersHistory",
		"Get",
		"/legacy/corporations/{corporation_id}/orders/history/",
		GetCorporationsCorporationIdOrdersHistory,
	)

	mockesi.NewRoute(
		"GetMarketsGroups",
		"Get",
		"/legacy/markets/groups/",
		GetMarketsGroups,
	)

	mockesi.NewRoute(
		"GetMarketsGroupsMarketGroupId",
		"Get",
		"/legacy/markets/groups/{market_group_id}/",
		GetMarketsGroupsMarketGroupId,
	)

	mockesi.NewRoute(
		"GetMarketsPrices",
		"Get",
		"/legacy/markets/prices/",
		GetMarketsPrices,
	)

	mockesi.NewRoute(
		"GetMarketsRegionIdHistory",
		"Get",
		"/legacy/markets/{region_id}/history/",
		GetMarketsRegionIdHistory,
	)

	mockesi.NewRoute(
		"GetMarketsRegionIdOrders",
		"Get",
		"/legacy/markets/{region_id}/orders/",
		GetMarketsRegionIdOrders,
	)

	mockesi.NewRoute(
		"GetMarketsRegionIdTypes",
		"Get",
		"/legacy/markets/{region_id}/types/",
		GetMarketsRegionIdTypes,
	)

	mockesi.NewRoute(
		"GetMarketsStructuresStructureId",
		"Get",
		"/legacy/markets/structures/{structure_id}/",
		GetMarketsStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOpportunities",
		"Get",
		"/legacy/characters/{character_id}/opportunities/",
		GetCharactersCharacterIdOpportunities,
	)

	mockesi.NewRoute(
		"GetOpportunitiesGroups",
		"Get",
		"/legacy/opportunities/groups/",
		GetOpportunitiesGroups,
	)

	mockesi.NewRoute(
		"GetOpportunitiesGroupsGroupId",
		"Get",
		"/legacy/opportunities/groups/{group_id}/",
		GetOpportunitiesGroupsGroupId,
	)

	mockesi.NewRoute(
		"GetOpportunitiesTasks",
		"Get",
		"/legacy/opportunities/tasks/",
		GetOpportunitiesTasks,
	)

	mockesi.NewRoute(
		"GetOpportunitiesTasksTaskId",
		"Get",
		"/legacy/opportunities/tasks/{task_id}/",
		GetOpportunitiesTasksTaskId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPlanets",
		"Get",
		"/legacy/characters/{character_id}/planets/",
		GetCharactersCharacterIdPlanets,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPlanetsPlanetId",
		"Get",
		"/legacy/characters/{character_id}/planets/{planet_id}/",
		GetCharactersCharacterIdPlanetsPlanetId,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdCustomsOffices",
		"Get",
		"/legacy/corporations/{corporation_id}/customs_offices/",
		GetCorporationsCorporationIdCustomsOffices,
	)

	mockesi.NewRoute(
		"GetUniverseSchematicsSchematicId",
		"Get",
		"/legacy/universe/schematics/{schematic_id}/",
		GetUniverseSchematicsSchematicId,
	)

	mockesi.NewRoute(
		"GetRouteOriginDestination",
		"Get",
		"/legacy/route/{origin}/{destination}/",
		GetRouteOriginDestination,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdSearch",
		"Get",
		"/legacy/characters/{character_id}/search/",
		GetCharactersCharacterIdSearch,
	)

	mockesi.NewRoute(
		"GetSearch",
		"Get",
		"/legacy/search/",
		GetSearch,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAttributes",
		"Get",
		"/legacy/characters/{character_id}/attributes/",
		GetCharactersCharacterIdAttributes,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdSkillqueue",
		"Get",
		"/legacy/characters/{character_id}/skillqueue/",
		GetCharactersCharacterIdSkillqueue,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdSkills",
		"Get",
		"/legacy/characters/{character_id}/skills/",
		GetCharactersCharacterIdSkills,
	)

	mockesi.NewRoute(
		"GetSovereigntyCampaigns",
		"Get",
		"/legacy/sovereignty/campaigns/",
		GetSovereigntyCampaigns,
	)

	mockesi.NewRoute(
		"GetSovereigntyMap",
		"Get",
		"/legacy/sovereignty/map/",
		GetSovereigntyMap,
	)

	mockesi.NewRoute(
		"GetSovereigntyStructures",
		"Get",
		"/legacy/sovereignty/structures/",
		GetSovereigntyStructures,
	)

	mockesi.NewRoute(
		"GetStatus",
		"Get",
		"/legacy/status/",
		GetStatus,
	)

	mockesi.NewRoute(
		"GetUniverseBloodlines",
		"Get",
		"/legacy/universe/bloodlines/",
		GetUniverseBloodlines,
	)

	mockesi.NewRoute(
		"GetUniverseCategories",
		"Get",
		"/legacy/universe/categories/",
		GetUniverseCategories,
	)

	mockesi.NewRoute(
		"GetUniverseCategoriesCategoryId",
		"Get",
		"/legacy/universe/categories/{category_id}/",
		GetUniverseCategoriesCategoryId,
	)

	mockesi.NewRoute(
		"GetUniverseConstellations",
		"Get",
		"/legacy/universe/constellations/",
		GetUniverseConstellations,
	)

	mockesi.NewRoute(
		"GetUniverseConstellationsConstellationId",
		"Get",
		"/legacy/universe/constellations/{constellation_id}/",
		GetUniverseConstellationsConstellationId,
	)

	mockesi.NewRoute(
		"GetUniverseFactions",
		"Get",
		"/legacy/universe/factions/",
		GetUniverseFactions,
	)

	mockesi.NewRoute(
		"GetUniverseGraphics",
		"Get",
		"/legacy/universe/graphics/",
		GetUniverseGraphics,
	)

	mockesi.NewRoute(
		"GetUniverseGraphicsGraphicId",
		"Get",
		"/legacy/universe/graphics/{graphic_id}/",
		GetUniverseGraphicsGraphicId,
	)

	mockesi.NewRoute(
		"GetUniverseGroups",
		"Get",
		"/legacy/universe/groups/",
		GetUniverseGroups,
	)

	mockesi.NewRoute(
		"GetUniverseGroupsGroupId",
		"Get",
		"/legacy/universe/groups/{group_id}/",
		GetUniverseGroupsGroupId,
	)

	mockesi.NewRoute(
		"GetUniverseMoonsMoonId",
		"Get",
		"/legacy/universe/moons/{moon_id}/",
		GetUniverseMoonsMoonId,
	)

	mockesi.NewRoute(
		"GetUniversePlanetsPlanetId",
		"Get",
		"/legacy/universe/planets/{planet_id}/",
		GetUniversePlanetsPlanetId,
	)

	mockesi.NewRoute(
		"GetUniverseRaces",
		"Get",
		"/legacy/universe/races/",
		GetUniverseRaces,
	)

	mockesi.NewRoute(
		"GetUniverseRegions",
		"Get",
		"/legacy/universe/regions/",
		GetUniverseRegions,
	)

	mockesi.NewRoute(
		"GetUniverseRegionsRegionId",
		"Get",
		"/legacy/universe/regions/{region_id}/",
		GetUniverseRegionsRegionId,
	)

	mockesi.NewRoute(
		"GetUniverseStargatesStargateId",
		"Get",
		"/legacy/universe/stargates/{stargate_id}/",
		GetUniverseStargatesStargateId,
	)

	mockesi.NewRoute(
		"GetUniverseStarsStarId",
		"Get",
		"/legacy/universe/stars/{star_id}/",
		GetUniverseStarsStarId,
	)

	mockesi.NewRoute(
		"GetUniverseStationsStationId",
		"Get",
		"/legacy/universe/stations/{station_id}/",
		GetUniverseStationsStationId,
	)

	mockesi.NewRoute(
		"GetUniverseStructures",
		"Get",
		"/legacy/universe/structures/",
		GetUniverseStructures,
	)

	mockesi.NewRoute(
		"GetUniverseStructuresStructureId",
		"Get",
		"/legacy/universe/structures/{structure_id}/",
		GetUniverseStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetUniverseSystemJumps",
		"Get",
		"/legacy/universe/system_jumps/",
		GetUniverseSystemJumps,
	)

	mockesi.NewRoute(
		"GetUniverseSystemKills",
		"Get",
		"/legacy/universe/system_kills/",
		GetUniverseSystemKills,
	)

	mockesi.NewRoute(
		"GetUniverseSystems",
		"Get",
		"/legacy/universe/systems/",
		GetUniverseSystems,
	)

	mockesi.NewRoute(
		"GetUniverseSystemsSystemId",
		"Get",
		"/legacy/universe/systems/{system_id}/",
		GetUniverseSystemsSystemId,
	)

	mockesi.NewRoute(
		"GetUniverseTypes",
		"Get",
		"/legacy/universe/types/",
		GetUniverseTypes,
	)

	mockesi.NewRoute(
		"GetUniverseTypesTypeId",
		"Get",
		"/legacy/universe/types/{type_id}/",
		GetUniverseTypesTypeId,
	)

	mockesi.NewRoute(
		"PostUniverseIds",
		"Post",
		"/legacy/universe/ids/",
		PostUniverseIds,
	)

	mockesi.NewRoute(
		"PostUniverseNames",
		"Post",
		"/legacy/universe/names/",
		PostUniverseNames,
	)

	mockesi.NewRoute(
		"PostUiAutopilotWaypoint",
		"Post",
		"/legacy/ui/autopilot/waypoint/",
		PostUiAutopilotWaypoint,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowContract",
		"Post",
		"/legacy/ui/openwindow/contract/",
		PostUiOpenwindowContract,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowInformation",
		"Post",
		"/legacy/ui/openwindow/information/",
		PostUiOpenwindowInformation,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowMarketdetails",
		"Post",
		"/legacy/ui/openwindow/marketdetails/",
		PostUiOpenwindowMarketdetails,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowNewmail",
		"Post",
		"/legacy/ui/openwindow/newmail/",
		PostUiOpenwindowNewmail,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdWallet",
		"Get",
		"/legacy/characters/{character_id}/wallet/",
		GetCharactersCharacterIdWallet,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdWalletJournal",
		"Get",
		"/legacy/characters/{character_id}/wallet/journal/",
		GetCharactersCharacterIdWalletJournal,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdWalletTransactions",
		"Get",
		"/legacy/characters/{character_id}/wallet/transactions/",
		GetCharactersCharacterIdWalletTransactions,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdWallets",
		"Get",
		"/legacy/corporations/{corporation_id}/wallets/",
		GetCorporationsCorporationIdWallets,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdWalletsDivisionJournal",
		"Get",
		"/legacy/corporations/{corporation_id}/wallets/{division}/journal/",
		GetCorporationsCorporationIdWalletsDivisionJournal,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdWalletsDivisionTransactions",
		"Get",
		"/legacy/corporations/{corporation_id}/wallets/{division}/transactions/",
		GetCorporationsCorporationIdWalletsDivisionTransactions,
	)

	mockesi.NewRoute(
		"GetWars",
		"Get",
		"/legacy/wars/",
		GetWars,
	)

	mockesi.NewRoute(
		"GetWarsWarId",
		"Get",
		"/legacy/wars/{war_id}/",
		GetWarsWarId,
	)

	mockesi.NewRoute(
		"GetWarsWarIdKillmails",
		"Get",
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
