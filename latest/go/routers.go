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
		"Get",
		"/latest/alliances/",
		GetAlliances,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceId",
		"Get",
		"/latest/alliances/{alliance_id}/",
		GetAlliancesAllianceId,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdCorporations",
		"Get",
		"/latest/alliances/{alliance_id}/corporations/",
		GetAlliancesAllianceIdCorporations,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdIcons",
		"Get",
		"/latest/alliances/{alliance_id}/icons/",
		GetAlliancesAllianceIdIcons,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAssets",
		"Get",
		"/latest/characters/{character_id}/assets/",
		GetCharactersCharacterIdAssets,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdAssets",
		"Get",
		"/latest/corporations/{corporation_id}/assets/",
		GetCorporationsCorporationIdAssets,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdAssetsLocations",
		"Post",
		"/latest/characters/{character_id}/assets/locations/",
		PostCharactersCharacterIdAssetsLocations,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdAssetsNames",
		"Post",
		"/latest/characters/{character_id}/assets/names/",
		PostCharactersCharacterIdAssetsNames,
	)

	mockesi.NewRoute(
		"PostCorporationsCorporationIdAssetsLocations",
		"Post",
		"/latest/corporations/{corporation_id}/assets/locations/",
		PostCorporationsCorporationIdAssetsLocations,
	)

	mockesi.NewRoute(
		"PostCorporationsCorporationIdAssetsNames",
		"Post",
		"/latest/corporations/{corporation_id}/assets/names/",
		PostCorporationsCorporationIdAssetsNames,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBookmarks",
		"Get",
		"/latest/characters/{character_id}/bookmarks/",
		GetCharactersCharacterIdBookmarks,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBookmarksFolders",
		"Get",
		"/latest/characters/{character_id}/bookmarks/folders/",
		GetCharactersCharacterIdBookmarksFolders,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdBookmarks",
		"Get",
		"/latest/corporations/{corporation_id}/bookmarks/",
		GetCorporationsCorporationIdBookmarks,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdBookmarksFolders",
		"Get",
		"/latest/corporations/{corporation_id}/bookmarks/folders/",
		GetCorporationsCorporationIdBookmarksFolders,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCalendar",
		"Get",
		"/latest/characters/{character_id}/calendar/",
		GetCharactersCharacterIdCalendar,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCalendarEventId",
		"Get",
		"/latest/characters/{character_id}/calendar/{event_id}/",
		GetCharactersCharacterIdCalendarEventId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCalendarEventIdAttendees",
		"Get",
		"/latest/characters/{character_id}/calendar/{event_id}/attendees/",
		GetCharactersCharacterIdCalendarEventIdAttendees,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdCalendarEventId",
		"Put",
		"/latest/characters/{character_id}/calendar/{event_id}/",
		PutCharactersCharacterIdCalendarEventId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterId",
		"Get",
		"/latest/characters/{character_id}/",
		GetCharactersCharacterId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAgentsResearch",
		"Get",
		"/latest/characters/{character_id}/agents_research/",
		GetCharactersCharacterIdAgentsResearch,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBlueprints",
		"Get",
		"/latest/characters/{character_id}/blueprints/",
		GetCharactersCharacterIdBlueprints,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCorporationhistory",
		"Get",
		"/latest/characters/{character_id}/corporationhistory/",
		GetCharactersCharacterIdCorporationhistory,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFatigue",
		"Get",
		"/latest/characters/{character_id}/fatigue/",
		GetCharactersCharacterIdFatigue,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMedals",
		"Get",
		"/latest/characters/{character_id}/medals/",
		GetCharactersCharacterIdMedals,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdNotifications",
		"Get",
		"/latest/characters/{character_id}/notifications/",
		GetCharactersCharacterIdNotifications,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdNotificationsContacts",
		"Get",
		"/latest/characters/{character_id}/notifications/contacts/",
		GetCharactersCharacterIdNotificationsContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPortrait",
		"Get",
		"/latest/characters/{character_id}/portrait/",
		GetCharactersCharacterIdPortrait,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdRoles",
		"Get",
		"/latest/characters/{character_id}/roles/",
		GetCharactersCharacterIdRoles,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdStandings",
		"Get",
		"/latest/characters/{character_id}/standings/",
		GetCharactersCharacterIdStandings,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdStats",
		"Get",
		"/latest/characters/{character_id}/stats/",
		GetCharactersCharacterIdStats,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdTitles",
		"Get",
		"/latest/characters/{character_id}/titles/",
		GetCharactersCharacterIdTitles,
	)

	mockesi.NewRoute(
		"PostCharactersAffiliation",
		"Post",
		"/latest/characters/affiliation/",
		PostCharactersAffiliation,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdCspa",
		"Post",
		"/latest/characters/{character_id}/cspa/",
		PostCharactersCharacterIdCspa,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdClones",
		"Get",
		"/latest/characters/{character_id}/clones/",
		GetCharactersCharacterIdClones,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdImplants",
		"Get",
		"/latest/characters/{character_id}/implants/",
		GetCharactersCharacterIdImplants,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdContacts",
		"Delete",
		"/latest/characters/{character_id}/contacts/",
		DeleteCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdContacts",
		"Get",
		"/latest/alliances/{alliance_id}/contacts/",
		GetAlliancesAllianceIdContacts,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdContactsLabels",
		"Get",
		"/latest/alliances/{alliance_id}/contacts/labels/",
		GetAlliancesAllianceIdContactsLabels,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContacts",
		"Get",
		"/latest/characters/{character_id}/contacts/",
		GetCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContactsLabels",
		"Get",
		"/latest/characters/{character_id}/contacts/labels/",
		GetCharactersCharacterIdContactsLabels,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContacts",
		"Get",
		"/latest/corporations/{corporation_id}/contacts/",
		GetCorporationsCorporationIdContacts,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContactsLabels",
		"Get",
		"/latest/corporations/{corporation_id}/contacts/labels/",
		GetCorporationsCorporationIdContactsLabels,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdContacts",
		"Post",
		"/latest/characters/{character_id}/contacts/",
		PostCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdContacts",
		"Put",
		"/latest/characters/{character_id}/contacts/",
		PutCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContracts",
		"Get",
		"/latest/characters/{character_id}/contracts/",
		GetCharactersCharacterIdContracts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContractsContractIdBids",
		"Get",
		"/latest/characters/{character_id}/contracts/{contract_id}/bids/",
		GetCharactersCharacterIdContractsContractIdBids,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContractsContractIdItems",
		"Get",
		"/latest/characters/{character_id}/contracts/{contract_id}/items/",
		GetCharactersCharacterIdContractsContractIdItems,
	)

	mockesi.NewRoute(
		"GetContractsPublicBidsContractId",
		"Get",
		"/latest/contracts/public/bids/{contract_id}/",
		GetContractsPublicBidsContractId,
	)

	mockesi.NewRoute(
		"GetContractsPublicItemsContractId",
		"Get",
		"/latest/contracts/public/items/{contract_id}/",
		GetContractsPublicItemsContractId,
	)

	mockesi.NewRoute(
		"GetContractsPublicRegionId",
		"Get",
		"/latest/contracts/public/{region_id}/",
		GetContractsPublicRegionId,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContracts",
		"Get",
		"/latest/corporations/{corporation_id}/contracts/",
		GetCorporationsCorporationIdContracts,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContractsContractIdBids",
		"Get",
		"/latest/corporations/{corporation_id}/contracts/{contract_id}/bids/",
		GetCorporationsCorporationIdContractsContractIdBids,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContractsContractIdItems",
		"Get",
		"/latest/corporations/{corporation_id}/contracts/{contract_id}/items/",
		GetCorporationsCorporationIdContractsContractIdItems,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationId",
		"Get",
		"/latest/corporations/{corporation_id}/",
		GetCorporationsCorporationId,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdAlliancehistory",
		"Get",
		"/latest/corporations/{corporation_id}/alliancehistory/",
		GetCorporationsCorporationIdAlliancehistory,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdBlueprints",
		"Get",
		"/latest/corporations/{corporation_id}/blueprints/",
		GetCorporationsCorporationIdBlueprints,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContainersLogs",
		"Get",
		"/latest/corporations/{corporation_id}/containers/logs/",
		GetCorporationsCorporationIdContainersLogs,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdDivisions",
		"Get",
		"/latest/corporations/{corporation_id}/divisions/",
		GetCorporationsCorporationIdDivisions,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdFacilities",
		"Get",
		"/latest/corporations/{corporation_id}/facilities/",
		GetCorporationsCorporationIdFacilities,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdIcons",
		"Get",
		"/latest/corporations/{corporation_id}/icons/",
		GetCorporationsCorporationIdIcons,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMedals",
		"Get",
		"/latest/corporations/{corporation_id}/medals/",
		GetCorporationsCorporationIdMedals,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMedalsIssued",
		"Get",
		"/latest/corporations/{corporation_id}/medals/issued/",
		GetCorporationsCorporationIdMedalsIssued,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMembers",
		"Get",
		"/latest/corporations/{corporation_id}/members/",
		GetCorporationsCorporationIdMembers,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMembersLimit",
		"Get",
		"/latest/corporations/{corporation_id}/members/limit/",
		GetCorporationsCorporationIdMembersLimit,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMembersTitles",
		"Get",
		"/latest/corporations/{corporation_id}/members/titles/",
		GetCorporationsCorporationIdMembersTitles,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMembertracking",
		"Get",
		"/latest/corporations/{corporation_id}/membertracking/",
		GetCorporationsCorporationIdMembertracking,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdRoles",
		"Get",
		"/latest/corporations/{corporation_id}/roles/",
		GetCorporationsCorporationIdRoles,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdRolesHistory",
		"Get",
		"/latest/corporations/{corporation_id}/roles/history/",
		GetCorporationsCorporationIdRolesHistory,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdShareholders",
		"Get",
		"/latest/corporations/{corporation_id}/shareholders/",
		GetCorporationsCorporationIdShareholders,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdStandings",
		"Get",
		"/latest/corporations/{corporation_id}/standings/",
		GetCorporationsCorporationIdStandings,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdStarbases",
		"Get",
		"/latest/corporations/{corporation_id}/starbases/",
		GetCorporationsCorporationIdStarbases,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdStarbasesStarbaseId",
		"Get",
		"/latest/corporations/{corporation_id}/starbases/{starbase_id}/",
		GetCorporationsCorporationIdStarbasesStarbaseId,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdStructures",
		"Get",
		"/latest/corporations/{corporation_id}/structures/",
		GetCorporationsCorporationIdStructures,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdTitles",
		"Get",
		"/latest/corporations/{corporation_id}/titles/",
		GetCorporationsCorporationIdTitles,
	)

	mockesi.NewRoute(
		"GetCorporationsNpccorps",
		"Get",
		"/latest/corporations/npccorps/",
		GetCorporationsNpccorps,
	)

	mockesi.NewRoute(
		"GetDogmaAttributes",
		"Get",
		"/latest/dogma/attributes/",
		GetDogmaAttributes,
	)

	mockesi.NewRoute(
		"GetDogmaAttributesAttributeId",
		"Get",
		"/latest/dogma/attributes/{attribute_id}/",
		GetDogmaAttributesAttributeId,
	)

	mockesi.NewRoute(
		"GetDogmaDynamicItemsTypeIdItemId",
		"Get",
		"/latest/dogma/dynamic/items/{type_id}/{item_id}/",
		GetDogmaDynamicItemsTypeIdItemId,
	)

	mockesi.NewRoute(
		"GetDogmaEffects",
		"Get",
		"/latest/dogma/effects/",
		GetDogmaEffects,
	)

	mockesi.NewRoute(
		"GetDogmaEffectsEffectId",
		"Get",
		"/latest/dogma/effects/{effect_id}/",
		GetDogmaEffectsEffectId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFwStats",
		"Get",
		"/latest/characters/{character_id}/fw/stats/",
		GetCharactersCharacterIdFwStats,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdFwStats",
		"Get",
		"/latest/corporations/{corporation_id}/fw/stats/",
		GetCorporationsCorporationIdFwStats,
	)

	mockesi.NewRoute(
		"GetFwLeaderboards",
		"Get",
		"/latest/fw/leaderboards/",
		GetFwLeaderboards,
	)

	mockesi.NewRoute(
		"GetFwLeaderboardsCharacters",
		"Get",
		"/latest/fw/leaderboards/characters/",
		GetFwLeaderboardsCharacters,
	)

	mockesi.NewRoute(
		"GetFwLeaderboardsCorporations",
		"Get",
		"/latest/fw/leaderboards/corporations/",
		GetFwLeaderboardsCorporations,
	)

	mockesi.NewRoute(
		"GetFwStats",
		"Get",
		"/latest/fw/stats/",
		GetFwStats,
	)

	mockesi.NewRoute(
		"GetFwSystems",
		"Get",
		"/latest/fw/systems/",
		GetFwSystems,
	)

	mockesi.NewRoute(
		"GetFwWars",
		"Get",
		"/latest/fw/wars/",
		GetFwWars,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdFittingsFittingId",
		"Delete",
		"/latest/characters/{character_id}/fittings/{fitting_id}/",
		DeleteCharactersCharacterIdFittingsFittingId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFittings",
		"Get",
		"/latest/characters/{character_id}/fittings/",
		GetCharactersCharacterIdFittings,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdFittings",
		"Post",
		"/latest/characters/{character_id}/fittings/",
		PostCharactersCharacterIdFittings,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdMembersMemberId",
		"Delete",
		"/latest/fleets/{fleet_id}/members/{member_id}/",
		DeleteFleetsFleetIdMembersMemberId,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdSquadsSquadId",
		"Delete",
		"/latest/fleets/{fleet_id}/squads/{squad_id}/",
		DeleteFleetsFleetIdSquadsSquadId,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdWingsWingId",
		"Delete",
		"/latest/fleets/{fleet_id}/wings/{wing_id}/",
		DeleteFleetsFleetIdWingsWingId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFleet",
		"Get",
		"/latest/characters/{character_id}/fleet/",
		GetCharactersCharacterIdFleet,
	)

	mockesi.NewRoute(
		"GetFleetsFleetId",
		"Get",
		"/latest/fleets/{fleet_id}/",
		GetFleetsFleetId,
	)

	mockesi.NewRoute(
		"GetFleetsFleetIdMembers",
		"Get",
		"/latest/fleets/{fleet_id}/members/",
		GetFleetsFleetIdMembers,
	)

	mockesi.NewRoute(
		"GetFleetsFleetIdWings",
		"Get",
		"/latest/fleets/{fleet_id}/wings/",
		GetFleetsFleetIdWings,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdMembers",
		"Post",
		"/latest/fleets/{fleet_id}/members/",
		PostFleetsFleetIdMembers,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdWings",
		"Post",
		"/latest/fleets/{fleet_id}/wings/",
		PostFleetsFleetIdWings,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdWingsWingIdSquads",
		"Post",
		"/latest/fleets/{fleet_id}/wings/{wing_id}/squads/",
		PostFleetsFleetIdWingsWingIdSquads,
	)

	mockesi.NewRoute(
		"PutFleetsFleetId",
		"Put",
		"/latest/fleets/{fleet_id}/",
		PutFleetsFleetId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdMembersMemberId",
		"Put",
		"/latest/fleets/{fleet_id}/members/{member_id}/",
		PutFleetsFleetIdMembersMemberId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdSquadsSquadId",
		"Put",
		"/latest/fleets/{fleet_id}/squads/{squad_id}/",
		PutFleetsFleetIdSquadsSquadId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdWingsWingId",
		"Put",
		"/latest/fleets/{fleet_id}/wings/{wing_id}/",
		PutFleetsFleetIdWingsWingId,
	)

	mockesi.NewRoute(
		"GetIncursions",
		"Get",
		"/latest/incursions/",
		GetIncursions,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdIndustryJobs",
		"Get",
		"/latest/characters/{character_id}/industry/jobs/",
		GetCharactersCharacterIdIndustryJobs,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMining",
		"Get",
		"/latest/characters/{character_id}/mining/",
		GetCharactersCharacterIdMining,
	)

	mockesi.NewRoute(
		"GetCorporationCorporationIdMiningExtractions",
		"Get",
		"/latest/corporation/{corporation_id}/mining/extractions/",
		GetCorporationCorporationIdMiningExtractions,
	)

	mockesi.NewRoute(
		"GetCorporationCorporationIdMiningObservers",
		"Get",
		"/latest/corporation/{corporation_id}/mining/observers/",
		GetCorporationCorporationIdMiningObservers,
	)

	mockesi.NewRoute(
		"GetCorporationCorporationIdMiningObserversObserverId",
		"Get",
		"/latest/corporation/{corporation_id}/mining/observers/{observer_id}/",
		GetCorporationCorporationIdMiningObserversObserverId,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdIndustryJobs",
		"Get",
		"/latest/corporations/{corporation_id}/industry/jobs/",
		GetCorporationsCorporationIdIndustryJobs,
	)

	mockesi.NewRoute(
		"GetIndustryFacilities",
		"Get",
		"/latest/industry/facilities/",
		GetIndustryFacilities,
	)

	mockesi.NewRoute(
		"GetIndustrySystems",
		"Get",
		"/latest/industry/systems/",
		GetIndustrySystems,
	)

	mockesi.NewRoute(
		"GetInsurancePrices",
		"Get",
		"/latest/insurance/prices/",
		GetInsurancePrices,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdKillmailsRecent",
		"Get",
		"/latest/characters/{character_id}/killmails/recent/",
		GetCharactersCharacterIdKillmailsRecent,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdKillmailsRecent",
		"Get",
		"/latest/corporations/{corporation_id}/killmails/recent/",
		GetCorporationsCorporationIdKillmailsRecent,
	)

	mockesi.NewRoute(
		"GetKillmailsKillmailIdKillmailHash",
		"Get",
		"/latest/killmails/{killmail_id}/{killmail_hash}/",
		GetKillmailsKillmailIdKillmailHash,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdLocation",
		"Get",
		"/latest/characters/{character_id}/location/",
		GetCharactersCharacterIdLocation,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOnline",
		"Get",
		"/latest/characters/{character_id}/online/",
		GetCharactersCharacterIdOnline,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdShip",
		"Get",
		"/latest/characters/{character_id}/ship/",
		GetCharactersCharacterIdShip,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdLoyaltyPoints",
		"Get",
		"/latest/characters/{character_id}/loyalty/points/",
		GetCharactersCharacterIdLoyaltyPoints,
	)

	mockesi.NewRoute(
		"GetLoyaltyStoresCorporationIdOffers",
		"Get",
		"/latest/loyalty/stores/{corporation_id}/offers/",
		GetLoyaltyStoresCorporationIdOffers,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdMailLabelsLabelId",
		"Delete",
		"/latest/characters/{character_id}/mail/labels/{label_id}/",
		DeleteCharactersCharacterIdMailLabelsLabelId,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdMailMailId",
		"Delete",
		"/latest/characters/{character_id}/mail/{mail_id}/",
		DeleteCharactersCharacterIdMailMailId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMail",
		"Get",
		"/latest/characters/{character_id}/mail/",
		GetCharactersCharacterIdMail,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailLabels",
		"Get",
		"/latest/characters/{character_id}/mail/labels/",
		GetCharactersCharacterIdMailLabels,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailLists",
		"Get",
		"/latest/characters/{character_id}/mail/lists/",
		GetCharactersCharacterIdMailLists,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailMailId",
		"Get",
		"/latest/characters/{character_id}/mail/{mail_id}/",
		GetCharactersCharacterIdMailMailId,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdMail",
		"Post",
		"/latest/characters/{character_id}/mail/",
		PostCharactersCharacterIdMail,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdMailLabels",
		"Post",
		"/latest/characters/{character_id}/mail/labels/",
		PostCharactersCharacterIdMailLabels,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdMailMailId",
		"Put",
		"/latest/characters/{character_id}/mail/{mail_id}/",
		PutCharactersCharacterIdMailMailId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOrders",
		"Get",
		"/latest/characters/{character_id}/orders/",
		GetCharactersCharacterIdOrders,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOrdersHistory",
		"Get",
		"/latest/characters/{character_id}/orders/history/",
		GetCharactersCharacterIdOrdersHistory,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdOrders",
		"Get",
		"/latest/corporations/{corporation_id}/orders/",
		GetCorporationsCorporationIdOrders,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdOrdersHistory",
		"Get",
		"/latest/corporations/{corporation_id}/orders/history/",
		GetCorporationsCorporationIdOrdersHistory,
	)

	mockesi.NewRoute(
		"GetMarketsGroups",
		"Get",
		"/latest/markets/groups/",
		GetMarketsGroups,
	)

	mockesi.NewRoute(
		"GetMarketsGroupsMarketGroupId",
		"Get",
		"/latest/markets/groups/{market_group_id}/",
		GetMarketsGroupsMarketGroupId,
	)

	mockesi.NewRoute(
		"GetMarketsPrices",
		"Get",
		"/latest/markets/prices/",
		GetMarketsPrices,
	)

	mockesi.NewRoute(
		"GetMarketsRegionIdHistory",
		"Get",
		"/latest/markets/{region_id}/history/",
		GetMarketsRegionIdHistory,
	)

	mockesi.NewRoute(
		"GetMarketsRegionIdOrders",
		"Get",
		"/latest/markets/{region_id}/orders/",
		GetMarketsRegionIdOrders,
	)

	mockesi.NewRoute(
		"GetMarketsRegionIdTypes",
		"Get",
		"/latest/markets/{region_id}/types/",
		GetMarketsRegionIdTypes,
	)

	mockesi.NewRoute(
		"GetMarketsStructuresStructureId",
		"Get",
		"/latest/markets/structures/{structure_id}/",
		GetMarketsStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOpportunities",
		"Get",
		"/latest/characters/{character_id}/opportunities/",
		GetCharactersCharacterIdOpportunities,
	)

	mockesi.NewRoute(
		"GetOpportunitiesGroups",
		"Get",
		"/latest/opportunities/groups/",
		GetOpportunitiesGroups,
	)

	mockesi.NewRoute(
		"GetOpportunitiesGroupsGroupId",
		"Get",
		"/latest/opportunities/groups/{group_id}/",
		GetOpportunitiesGroupsGroupId,
	)

	mockesi.NewRoute(
		"GetOpportunitiesTasks",
		"Get",
		"/latest/opportunities/tasks/",
		GetOpportunitiesTasks,
	)

	mockesi.NewRoute(
		"GetOpportunitiesTasksTaskId",
		"Get",
		"/latest/opportunities/tasks/{task_id}/",
		GetOpportunitiesTasksTaskId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPlanets",
		"Get",
		"/latest/characters/{character_id}/planets/",
		GetCharactersCharacterIdPlanets,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPlanetsPlanetId",
		"Get",
		"/latest/characters/{character_id}/planets/{planet_id}/",
		GetCharactersCharacterIdPlanetsPlanetId,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdCustomsOffices",
		"Get",
		"/latest/corporations/{corporation_id}/customs_offices/",
		GetCorporationsCorporationIdCustomsOffices,
	)

	mockesi.NewRoute(
		"GetUniverseSchematicsSchematicId",
		"Get",
		"/latest/universe/schematics/{schematic_id}/",
		GetUniverseSchematicsSchematicId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdSearch",
		"Get",
		"/latest/characters/{character_id}/search/",
		GetCharactersCharacterIdSearch,
	)

	mockesi.NewRoute(
		"GetSearch",
		"Get",
		"/latest/search/",
		GetSearch,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAttributes",
		"Get",
		"/latest/characters/{character_id}/attributes/",
		GetCharactersCharacterIdAttributes,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdSkillqueue",
		"Get",
		"/latest/characters/{character_id}/skillqueue/",
		GetCharactersCharacterIdSkillqueue,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdSkills",
		"Get",
		"/latest/characters/{character_id}/skills/",
		GetCharactersCharacterIdSkills,
	)

	mockesi.NewRoute(
		"GetSovereigntyCampaigns",
		"Get",
		"/latest/sovereignty/campaigns/",
		GetSovereigntyCampaigns,
	)

	mockesi.NewRoute(
		"GetSovereigntyMap",
		"Get",
		"/latest/sovereignty/map/",
		GetSovereigntyMap,
	)

	mockesi.NewRoute(
		"GetSovereigntyStructures",
		"Get",
		"/latest/sovereignty/structures/",
		GetSovereigntyStructures,
	)

	mockesi.NewRoute(
		"GetStatus",
		"Get",
		"/latest/status/",
		GetStatus,
	)

	mockesi.NewRoute(
		"GetUniverseAncestries",
		"Get",
		"/latest/universe/ancestries/",
		GetUniverseAncestries,
	)

	mockesi.NewRoute(
		"GetUniverseAsteroidBeltsAsteroidBeltId",
		"Get",
		"/latest/universe/asteroid_belts/{asteroid_belt_id}/",
		GetUniverseAsteroidBeltsAsteroidBeltId,
	)

	mockesi.NewRoute(
		"GetUniverseBloodlines",
		"Get",
		"/latest/universe/bloodlines/",
		GetUniverseBloodlines,
	)

	mockesi.NewRoute(
		"GetUniverseCategories",
		"Get",
		"/latest/universe/categories/",
		GetUniverseCategories,
	)

	mockesi.NewRoute(
		"GetUniverseCategoriesCategoryId",
		"Get",
		"/latest/universe/categories/{category_id}/",
		GetUniverseCategoriesCategoryId,
	)

	mockesi.NewRoute(
		"GetUniverseConstellations",
		"Get",
		"/latest/universe/constellations/",
		GetUniverseConstellations,
	)

	mockesi.NewRoute(
		"GetUniverseConstellationsConstellationId",
		"Get",
		"/latest/universe/constellations/{constellation_id}/",
		GetUniverseConstellationsConstellationId,
	)

	mockesi.NewRoute(
		"GetUniverseFactions",
		"Get",
		"/latest/universe/factions/",
		GetUniverseFactions,
	)

	mockesi.NewRoute(
		"GetUniverseGraphics",
		"Get",
		"/latest/universe/graphics/",
		GetUniverseGraphics,
	)

	mockesi.NewRoute(
		"GetUniverseGraphicsGraphicId",
		"Get",
		"/latest/universe/graphics/{graphic_id}/",
		GetUniverseGraphicsGraphicId,
	)

	mockesi.NewRoute(
		"GetUniverseGroups",
		"Get",
		"/latest/universe/groups/",
		GetUniverseGroups,
	)

	mockesi.NewRoute(
		"GetUniverseGroupsGroupId",
		"Get",
		"/latest/universe/groups/{group_id}/",
		GetUniverseGroupsGroupId,
	)

	mockesi.NewRoute(
		"GetUniverseMoonsMoonId",
		"Get",
		"/latest/universe/moons/{moon_id}/",
		GetUniverseMoonsMoonId,
	)

	mockesi.NewRoute(
		"GetUniversePlanetsPlanetId",
		"Get",
		"/latest/universe/planets/{planet_id}/",
		GetUniversePlanetsPlanetId,
	)

	mockesi.NewRoute(
		"GetUniverseRaces",
		"Get",
		"/latest/universe/races/",
		GetUniverseRaces,
	)

	mockesi.NewRoute(
		"GetUniverseRegions",
		"Get",
		"/latest/universe/regions/",
		GetUniverseRegions,
	)

	mockesi.NewRoute(
		"GetUniverseRegionsRegionId",
		"Get",
		"/latest/universe/regions/{region_id}/",
		GetUniverseRegionsRegionId,
	)

	mockesi.NewRoute(
		"GetUniverseStargatesStargateId",
		"Get",
		"/latest/universe/stargates/{stargate_id}/",
		GetUniverseStargatesStargateId,
	)

	mockesi.NewRoute(
		"GetUniverseStarsStarId",
		"Get",
		"/latest/universe/stars/{star_id}/",
		GetUniverseStarsStarId,
	)

	mockesi.NewRoute(
		"GetUniverseStationsStationId",
		"Get",
		"/latest/universe/stations/{station_id}/",
		GetUniverseStationsStationId,
	)

	mockesi.NewRoute(
		"GetUniverseStructures",
		"Get",
		"/latest/universe/structures/",
		GetUniverseStructures,
	)

	mockesi.NewRoute(
		"GetUniverseStructuresStructureId",
		"Get",
		"/latest/universe/structures/{structure_id}/",
		GetUniverseStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetUniverseSystemJumps",
		"Get",
		"/latest/universe/system_jumps/",
		GetUniverseSystemJumps,
	)

	mockesi.NewRoute(
		"GetUniverseSystemKills",
		"Get",
		"/latest/universe/system_kills/",
		GetUniverseSystemKills,
	)

	mockesi.NewRoute(
		"GetUniverseSystems",
		"Get",
		"/latest/universe/systems/",
		GetUniverseSystems,
	)

	mockesi.NewRoute(
		"GetUniverseSystemsSystemId",
		"Get",
		"/latest/universe/systems/{system_id}/",
		GetUniverseSystemsSystemId,
	)

	mockesi.NewRoute(
		"GetUniverseTypes",
		"Get",
		"/latest/universe/types/",
		GetUniverseTypes,
	)

	mockesi.NewRoute(
		"GetUniverseTypesTypeId",
		"Get",
		"/latest/universe/types/{type_id}/",
		GetUniverseTypesTypeId,
	)

	mockesi.NewRoute(
		"PostUniverseIds",
		"Post",
		"/latest/universe/ids/",
		PostUniverseIds,
	)

	mockesi.NewRoute(
		"PostUniverseNames",
		"Post",
		"/latest/universe/names/",
		PostUniverseNames,
	)

	mockesi.NewRoute(
		"PostUiAutopilotWaypoint",
		"Post",
		"/latest/ui/autopilot/waypoint/",
		PostUiAutopilotWaypoint,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowContract",
		"Post",
		"/latest/ui/openwindow/contract/",
		PostUiOpenwindowContract,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowInformation",
		"Post",
		"/latest/ui/openwindow/information/",
		PostUiOpenwindowInformation,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowMarketdetails",
		"Post",
		"/latest/ui/openwindow/marketdetails/",
		PostUiOpenwindowMarketdetails,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowNewmail",
		"Post",
		"/latest/ui/openwindow/newmail/",
		PostUiOpenwindowNewmail,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdWallet",
		"Get",
		"/latest/characters/{character_id}/wallet/",
		GetCharactersCharacterIdWallet,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdWalletJournal",
		"Get",
		"/latest/characters/{character_id}/wallet/journal/",
		GetCharactersCharacterIdWalletJournal,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdWalletTransactions",
		"Get",
		"/latest/characters/{character_id}/wallet/transactions/",
		GetCharactersCharacterIdWalletTransactions,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdWallets",
		"Get",
		"/latest/corporations/{corporation_id}/wallets/",
		GetCorporationsCorporationIdWallets,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdWalletsDivisionJournal",
		"Get",
		"/latest/corporations/{corporation_id}/wallets/{division}/journal/",
		GetCorporationsCorporationIdWalletsDivisionJournal,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdWalletsDivisionTransactions",
		"Get",
		"/latest/corporations/{corporation_id}/wallets/{division}/transactions/",
		GetCorporationsCorporationIdWalletsDivisionTransactions,
	)

	mockesi.NewRoute(
		"GetWars",
		"Get",
		"/latest/wars/",
		GetWars,
	)

	mockesi.NewRoute(
		"GetWarsWarId",
		"Get",
		"/latest/wars/{war_id}/",
		GetWarsWarId,
	)

	mockesi.NewRoute(
		"GetWarsWarIdKillmails",
		"Get",
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
