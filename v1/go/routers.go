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
		"Get",
		"/v1/alliances/",
		GetAlliances,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdCorporations",
		"Get",
		"/v1/alliances/{alliance_id}/corporations/",
		GetAlliancesAllianceIdCorporations,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdIcons",
		"Get",
		"/v1/alliances/{alliance_id}/icons/",
		GetAlliancesAllianceIdIcons,
	)

	mockesi.NewRoute(
		"GetAlliancesNames",
		"Get",
		"/v1/alliances/names/",
		GetAlliancesNames,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdAssetsLocations",
		"Post",
		"/v1/characters/{character_id}/assets/locations/",
		PostCharactersCharacterIdAssetsLocations,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdAssetsNames",
		"Post",
		"/v1/characters/{character_id}/assets/names/",
		PostCharactersCharacterIdAssetsNames,
	)

	mockesi.NewRoute(
		"PostCorporationsCorporationIdAssetsLocations",
		"Post",
		"/v1/corporations/{corporation_id}/assets/locations/",
		PostCorporationsCorporationIdAssetsLocations,
	)

	mockesi.NewRoute(
		"PostCorporationsCorporationIdAssetsNames",
		"Post",
		"/v1/corporations/{corporation_id}/assets/names/",
		PostCorporationsCorporationIdAssetsNames,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBookmarks",
		"Get",
		"/v1/characters/{character_id}/bookmarks/",
		GetCharactersCharacterIdBookmarks,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBookmarksFolders",
		"Get",
		"/v1/characters/{character_id}/bookmarks/folders/",
		GetCharactersCharacterIdBookmarksFolders,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdBookmarks",
		"Get",
		"/v1/corporations/{corporation_id}/bookmarks/",
		GetCorporationsCorporationIdBookmarks,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdBookmarksFolders",
		"Get",
		"/v1/corporations/{corporation_id}/bookmarks/folders/",
		GetCorporationsCorporationIdBookmarksFolders,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCalendar",
		"Get",
		"/v1/characters/{character_id}/calendar/",
		GetCharactersCharacterIdCalendar,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCalendarEventIdAttendees",
		"Get",
		"/v1/characters/{character_id}/calendar/{event_id}/attendees/",
		GetCharactersCharacterIdCalendarEventIdAttendees,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAgentsResearch",
		"Get",
		"/v1/characters/{character_id}/agents_research/",
		GetCharactersCharacterIdAgentsResearch,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdBlueprints",
		"Get",
		"/v1/characters/{character_id}/blueprints/",
		GetCharactersCharacterIdBlueprints,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdCorporationhistory",
		"Get",
		"/v1/characters/{character_id}/corporationhistory/",
		GetCharactersCharacterIdCorporationhistory,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFatigue",
		"Get",
		"/v1/characters/{character_id}/fatigue/",
		GetCharactersCharacterIdFatigue,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMedals",
		"Get",
		"/v1/characters/{character_id}/medals/",
		GetCharactersCharacterIdMedals,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdNotifications",
		"Get",
		"/v1/characters/{character_id}/notifications/",
		GetCharactersCharacterIdNotifications,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdNotificationsContacts",
		"Get",
		"/v1/characters/{character_id}/notifications/contacts/",
		GetCharactersCharacterIdNotificationsContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPortrait",
		"Get",
		"/v1/characters/{character_id}/portrait/",
		GetCharactersCharacterIdPortrait,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdRoles",
		"Get",
		"/v1/characters/{character_id}/roles/",
		GetCharactersCharacterIdRoles,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdStandings",
		"Get",
		"/v1/characters/{character_id}/standings/",
		GetCharactersCharacterIdStandings,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdStats",
		"Get",
		"/v1/characters/{character_id}/stats/",
		GetCharactersCharacterIdStats,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdTitles",
		"Get",
		"/v1/characters/{character_id}/titles/",
		GetCharactersCharacterIdTitles,
	)

	mockesi.NewRoute(
		"PostCharactersAffiliation",
		"Post",
		"/v1/characters/affiliation/",
		PostCharactersAffiliation,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdImplants",
		"Get",
		"/v1/characters/{character_id}/implants/",
		GetCharactersCharacterIdImplants,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdContacts",
		"Delete",
		"/v1/characters/{character_id}/contacts/",
		DeleteCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdContacts",
		"Get",
		"/v1/alliances/{alliance_id}/contacts/",
		GetAlliancesAllianceIdContacts,
	)

	mockesi.NewRoute(
		"GetAlliancesAllianceIdContactsLabels",
		"Get",
		"/v1/alliances/{alliance_id}/contacts/labels/",
		GetAlliancesAllianceIdContactsLabels,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContacts",
		"Get",
		"/v1/characters/{character_id}/contacts/",
		GetCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContactsLabels",
		"Get",
		"/v1/characters/{character_id}/contacts/labels/",
		GetCharactersCharacterIdContactsLabels,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContacts",
		"Get",
		"/v1/corporations/{corporation_id}/contacts/",
		GetCorporationsCorporationIdContacts,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContactsLabels",
		"Get",
		"/v1/corporations/{corporation_id}/contacts/labels/",
		GetCorporationsCorporationIdContactsLabels,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdContacts",
		"Post",
		"/v1/characters/{character_id}/contacts/",
		PostCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdContacts",
		"Put",
		"/v1/characters/{character_id}/contacts/",
		PutCharactersCharacterIdContacts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContracts",
		"Get",
		"/v1/characters/{character_id}/contracts/",
		GetCharactersCharacterIdContracts,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContractsContractIdBids",
		"Get",
		"/v1/characters/{character_id}/contracts/{contract_id}/bids/",
		GetCharactersCharacterIdContractsContractIdBids,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdContractsContractIdItems",
		"Get",
		"/v1/characters/{character_id}/contracts/{contract_id}/items/",
		GetCharactersCharacterIdContractsContractIdItems,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContracts",
		"Get",
		"/v1/corporations/{corporation_id}/contracts/",
		GetCorporationsCorporationIdContracts,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContractsContractIdBids",
		"Get",
		"/v1/corporations/{corporation_id}/contracts/{contract_id}/bids/",
		GetCorporationsCorporationIdContractsContractIdBids,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContractsContractIdItems",
		"Get",
		"/v1/corporations/{corporation_id}/contracts/{contract_id}/items/",
		GetCorporationsCorporationIdContractsContractIdItems,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdAlliancehistory",
		"Get",
		"/v1/corporations/{corporation_id}/alliancehistory/",
		GetCorporationsCorporationIdAlliancehistory,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdBlueprints",
		"Get",
		"/v1/corporations/{corporation_id}/blueprints/",
		GetCorporationsCorporationIdBlueprints,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdContainersLogs",
		"Get",
		"/v1/corporations/{corporation_id}/containers/logs/",
		GetCorporationsCorporationIdContainersLogs,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdDivisions",
		"Get",
		"/v1/corporations/{corporation_id}/divisions/",
		GetCorporationsCorporationIdDivisions,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdFacilities",
		"Get",
		"/v1/corporations/{corporation_id}/facilities/",
		GetCorporationsCorporationIdFacilities,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdIcons",
		"Get",
		"/v1/corporations/{corporation_id}/icons/",
		GetCorporationsCorporationIdIcons,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMedals",
		"Get",
		"/v1/corporations/{corporation_id}/medals/",
		GetCorporationsCorporationIdMedals,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMedalsIssued",
		"Get",
		"/v1/corporations/{corporation_id}/medals/issued/",
		GetCorporationsCorporationIdMedalsIssued,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMembersLimit",
		"Get",
		"/v1/corporations/{corporation_id}/members/limit/",
		GetCorporationsCorporationIdMembersLimit,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMembersTitles",
		"Get",
		"/v1/corporations/{corporation_id}/members/titles/",
		GetCorporationsCorporationIdMembersTitles,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdMembertracking",
		"Get",
		"/v1/corporations/{corporation_id}/membertracking/",
		GetCorporationsCorporationIdMembertracking,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdOutposts",
		"Get",
		"/v1/corporations/{corporation_id}/outposts/",
		GetCorporationsCorporationIdOutposts,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdOutpostsOutpostId",
		"Get",
		"/v1/corporations/{corporation_id}/outposts/{outpost_id}/",
		GetCorporationsCorporationIdOutpostsOutpostId,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdRoles",
		"Get",
		"/v1/corporations/{corporation_id}/roles/",
		GetCorporationsCorporationIdRoles,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdRolesHistory",
		"Get",
		"/v1/corporations/{corporation_id}/roles/history/",
		GetCorporationsCorporationIdRolesHistory,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdShareholders",
		"Get",
		"/v1/corporations/{corporation_id}/shareholders/",
		GetCorporationsCorporationIdShareholders,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdStandings",
		"Get",
		"/v1/corporations/{corporation_id}/standings/",
		GetCorporationsCorporationIdStandings,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdStarbases",
		"Get",
		"/v1/corporations/{corporation_id}/starbases/",
		GetCorporationsCorporationIdStarbases,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdStarbasesStarbaseId",
		"Get",
		"/v1/corporations/{corporation_id}/starbases/{starbase_id}/",
		GetCorporationsCorporationIdStarbasesStarbaseId,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdStructures",
		"Get",
		"/v1/corporations/{corporation_id}/structures/",
		GetCorporationsCorporationIdStructures,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdTitles",
		"Get",
		"/v1/corporations/{corporation_id}/titles/",
		GetCorporationsCorporationIdTitles,
	)

	mockesi.NewRoute(
		"GetCorporationsNpccorps",
		"Get",
		"/v1/corporations/npccorps/",
		GetCorporationsNpccorps,
	)

	mockesi.NewRoute(
		"PutCorporationsCorporationIdStructuresStructureId",
		"Put",
		"/v1/corporations/{corporation_id}/structures/{structure_id}/",
		PutCorporationsCorporationIdStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetDogmaAttributes",
		"Get",
		"/v1/dogma/attributes/",
		GetDogmaAttributes,
	)

	mockesi.NewRoute(
		"GetDogmaAttributesAttributeId",
		"Get",
		"/v1/dogma/attributes/{attribute_id}/",
		GetDogmaAttributesAttributeId,
	)

	mockesi.NewRoute(
		"GetDogmaDynamicItemsTypeIdItemId",
		"Get",
		"/v1/dogma/dynamic/items/{type_id}/{item_id}/",
		GetDogmaDynamicItemsTypeIdItemId,
	)

	mockesi.NewRoute(
		"GetDogmaEffects",
		"Get",
		"/v1/dogma/effects/",
		GetDogmaEffects,
	)

	mockesi.NewRoute(
		"GetDogmaEffectsEffectId",
		"Get",
		"/v1/dogma/effects/{effect_id}/",
		GetDogmaEffectsEffectId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFwStats",
		"Get",
		"/v1/characters/{character_id}/fw/stats/",
		GetCharactersCharacterIdFwStats,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdFwStats",
		"Get",
		"/v1/corporations/{corporation_id}/fw/stats/",
		GetCorporationsCorporationIdFwStats,
	)

	mockesi.NewRoute(
		"GetFwLeaderboards",
		"Get",
		"/v1/fw/leaderboards/",
		GetFwLeaderboards,
	)

	mockesi.NewRoute(
		"GetFwLeaderboardsCharacters",
		"Get",
		"/v1/fw/leaderboards/characters/",
		GetFwLeaderboardsCharacters,
	)

	mockesi.NewRoute(
		"GetFwLeaderboardsCorporations",
		"Get",
		"/v1/fw/leaderboards/corporations/",
		GetFwLeaderboardsCorporations,
	)

	mockesi.NewRoute(
		"GetFwStats",
		"Get",
		"/v1/fw/stats/",
		GetFwStats,
	)

	mockesi.NewRoute(
		"GetFwSystems",
		"Get",
		"/v1/fw/systems/",
		GetFwSystems,
	)

	mockesi.NewRoute(
		"GetFwWars",
		"Get",
		"/v1/fw/wars/",
		GetFwWars,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdFittingsFittingId",
		"Delete",
		"/v1/characters/{character_id}/fittings/{fitting_id}/",
		DeleteCharactersCharacterIdFittingsFittingId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFittings",
		"Get",
		"/v1/characters/{character_id}/fittings/",
		GetCharactersCharacterIdFittings,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdFittings",
		"Post",
		"/v1/characters/{character_id}/fittings/",
		PostCharactersCharacterIdFittings,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdMembersMemberId",
		"Delete",
		"/v1/fleets/{fleet_id}/members/{member_id}/",
		DeleteFleetsFleetIdMembersMemberId,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdSquadsSquadId",
		"Delete",
		"/v1/fleets/{fleet_id}/squads/{squad_id}/",
		DeleteFleetsFleetIdSquadsSquadId,
	)

	mockesi.NewRoute(
		"DeleteFleetsFleetIdWingsWingId",
		"Delete",
		"/v1/fleets/{fleet_id}/wings/{wing_id}/",
		DeleteFleetsFleetIdWingsWingId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdFleet",
		"Get",
		"/v1/characters/{character_id}/fleet/",
		GetCharactersCharacterIdFleet,
	)

	mockesi.NewRoute(
		"GetFleetsFleetId",
		"Get",
		"/v1/fleets/{fleet_id}/",
		GetFleetsFleetId,
	)

	mockesi.NewRoute(
		"GetFleetsFleetIdMembers",
		"Get",
		"/v1/fleets/{fleet_id}/members/",
		GetFleetsFleetIdMembers,
	)

	mockesi.NewRoute(
		"GetFleetsFleetIdWings",
		"Get",
		"/v1/fleets/{fleet_id}/wings/",
		GetFleetsFleetIdWings,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdMembers",
		"Post",
		"/v1/fleets/{fleet_id}/members/",
		PostFleetsFleetIdMembers,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdWings",
		"Post",
		"/v1/fleets/{fleet_id}/wings/",
		PostFleetsFleetIdWings,
	)

	mockesi.NewRoute(
		"PostFleetsFleetIdWingsWingIdSquads",
		"Post",
		"/v1/fleets/{fleet_id}/wings/{wing_id}/squads/",
		PostFleetsFleetIdWingsWingIdSquads,
	)

	mockesi.NewRoute(
		"PutFleetsFleetId",
		"Put",
		"/v1/fleets/{fleet_id}/",
		PutFleetsFleetId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdMembersMemberId",
		"Put",
		"/v1/fleets/{fleet_id}/members/{member_id}/",
		PutFleetsFleetIdMembersMemberId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdSquadsSquadId",
		"Put",
		"/v1/fleets/{fleet_id}/squads/{squad_id}/",
		PutFleetsFleetIdSquadsSquadId,
	)

	mockesi.NewRoute(
		"PutFleetsFleetIdWingsWingId",
		"Put",
		"/v1/fleets/{fleet_id}/wings/{wing_id}/",
		PutFleetsFleetIdWingsWingId,
	)

	mockesi.NewRoute(
		"GetIncursions",
		"Get",
		"/v1/incursions/",
		GetIncursions,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdIndustryJobs",
		"Get",
		"/v1/characters/{character_id}/industry/jobs/",
		GetCharactersCharacterIdIndustryJobs,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMining",
		"Get",
		"/v1/characters/{character_id}/mining/",
		GetCharactersCharacterIdMining,
	)

	mockesi.NewRoute(
		"GetCorporationCorporationIdMiningExtractions",
		"Get",
		"/v1/corporation/{corporation_id}/mining/extractions/",
		GetCorporationCorporationIdMiningExtractions,
	)

	mockesi.NewRoute(
		"GetCorporationCorporationIdMiningObservers",
		"Get",
		"/v1/corporation/{corporation_id}/mining/observers/",
		GetCorporationCorporationIdMiningObservers,
	)

	mockesi.NewRoute(
		"GetCorporationCorporationIdMiningObserversObserverId",
		"Get",
		"/v1/corporation/{corporation_id}/mining/observers/{observer_id}/",
		GetCorporationCorporationIdMiningObserversObserverId,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdIndustryJobs",
		"Get",
		"/v1/corporations/{corporation_id}/industry/jobs/",
		GetCorporationsCorporationIdIndustryJobs,
	)

	mockesi.NewRoute(
		"GetIndustryFacilities",
		"Get",
		"/v1/industry/facilities/",
		GetIndustryFacilities,
	)

	mockesi.NewRoute(
		"GetIndustrySystems",
		"Get",
		"/v1/industry/systems/",
		GetIndustrySystems,
	)

	mockesi.NewRoute(
		"GetInsurancePrices",
		"Get",
		"/v1/insurance/prices/",
		GetInsurancePrices,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdKillmailsRecent",
		"Get",
		"/v1/characters/{character_id}/killmails/recent/",
		GetCharactersCharacterIdKillmailsRecent,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdKillmailsRecent",
		"Get",
		"/v1/corporations/{corporation_id}/killmails/recent/",
		GetCorporationsCorporationIdKillmailsRecent,
	)

	mockesi.NewRoute(
		"GetKillmailsKillmailIdKillmailHash",
		"Get",
		"/v1/killmails/{killmail_id}/{killmail_hash}/",
		GetKillmailsKillmailIdKillmailHash,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdLocation",
		"Get",
		"/v1/characters/{character_id}/location/",
		GetCharactersCharacterIdLocation,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOnline",
		"Get",
		"/v1/characters/{character_id}/online/",
		GetCharactersCharacterIdOnline,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdShip",
		"Get",
		"/v1/characters/{character_id}/ship/",
		GetCharactersCharacterIdShip,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdLoyaltyPoints",
		"Get",
		"/v1/characters/{character_id}/loyalty/points/",
		GetCharactersCharacterIdLoyaltyPoints,
	)

	mockesi.NewRoute(
		"GetLoyaltyStoresCorporationIdOffers",
		"Get",
		"/v1/loyalty/stores/{corporation_id}/offers/",
		GetLoyaltyStoresCorporationIdOffers,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdMailLabelsLabelId",
		"Delete",
		"/v1/characters/{character_id}/mail/labels/{label_id}/",
		DeleteCharactersCharacterIdMailLabelsLabelId,
	)

	mockesi.NewRoute(
		"DeleteCharactersCharacterIdMailMailId",
		"Delete",
		"/v1/characters/{character_id}/mail/{mail_id}/",
		DeleteCharactersCharacterIdMailMailId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMail",
		"Get",
		"/v1/characters/{character_id}/mail/",
		GetCharactersCharacterIdMail,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailLists",
		"Get",
		"/v1/characters/{character_id}/mail/lists/",
		GetCharactersCharacterIdMailLists,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailMailId",
		"Get",
		"/v1/characters/{character_id}/mail/{mail_id}/",
		GetCharactersCharacterIdMailMailId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdMailUnread",
		"Get",
		"/v1/characters/{character_id}/mail/unread/",
		GetCharactersCharacterIdMailUnread,
	)

	mockesi.NewRoute(
		"PostCharactersCharacterIdMail",
		"Post",
		"/v1/characters/{character_id}/mail/",
		PostCharactersCharacterIdMail,
	)

	mockesi.NewRoute(
		"PutCharactersCharacterIdMailMailId",
		"Put",
		"/v1/characters/{character_id}/mail/{mail_id}/",
		PutCharactersCharacterIdMailMailId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOrders",
		"Get",
		"/v1/characters/{character_id}/orders/",
		GetCharactersCharacterIdOrders,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOrdersHistory",
		"Get",
		"/v1/characters/{character_id}/orders/history/",
		GetCharactersCharacterIdOrdersHistory,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdOrders",
		"Get",
		"/v1/corporations/{corporation_id}/orders/",
		GetCorporationsCorporationIdOrders,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdOrdersHistory",
		"Get",
		"/v1/corporations/{corporation_id}/orders/history/",
		GetCorporationsCorporationIdOrdersHistory,
	)

	mockesi.NewRoute(
		"GetMarketsGroups",
		"Get",
		"/v1/markets/groups/",
		GetMarketsGroups,
	)

	mockesi.NewRoute(
		"GetMarketsGroupsMarketGroupId",
		"Get",
		"/v1/markets/groups/{market_group_id}/",
		GetMarketsGroupsMarketGroupId,
	)

	mockesi.NewRoute(
		"GetMarketsPrices",
		"Get",
		"/v1/markets/prices/",
		GetMarketsPrices,
	)

	mockesi.NewRoute(
		"GetMarketsRegionIdHistory",
		"Get",
		"/v1/markets/{region_id}/history/",
		GetMarketsRegionIdHistory,
	)

	mockesi.NewRoute(
		"GetMarketsRegionIdOrders",
		"Get",
		"/v1/markets/{region_id}/orders/",
		GetMarketsRegionIdOrders,
	)

	mockesi.NewRoute(
		"GetMarketsRegionIdTypes",
		"Get",
		"/v1/markets/{region_id}/types/",
		GetMarketsRegionIdTypes,
	)

	mockesi.NewRoute(
		"GetMarketsStructuresStructureId",
		"Get",
		"/v1/markets/structures/{structure_id}/",
		GetMarketsStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdOpportunities",
		"Get",
		"/v1/characters/{character_id}/opportunities/",
		GetCharactersCharacterIdOpportunities,
	)

	mockesi.NewRoute(
		"GetOpportunitiesGroups",
		"Get",
		"/v1/opportunities/groups/",
		GetOpportunitiesGroups,
	)

	mockesi.NewRoute(
		"GetOpportunitiesGroupsGroupId",
		"Get",
		"/v1/opportunities/groups/{group_id}/",
		GetOpportunitiesGroupsGroupId,
	)

	mockesi.NewRoute(
		"GetOpportunitiesTasks",
		"Get",
		"/v1/opportunities/tasks/",
		GetOpportunitiesTasks,
	)

	mockesi.NewRoute(
		"GetOpportunitiesTasksTaskId",
		"Get",
		"/v1/opportunities/tasks/{task_id}/",
		GetOpportunitiesTasksTaskId,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdPlanets",
		"Get",
		"/v1/characters/{character_id}/planets/",
		GetCharactersCharacterIdPlanets,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdCustomsOffices",
		"Get",
		"/v1/corporations/{corporation_id}/customs_offices/",
		GetCorporationsCorporationIdCustomsOffices,
	)

	mockesi.NewRoute(
		"GetUniverseSchematicsSchematicId",
		"Get",
		"/v1/universe/schematics/{schematic_id}/",
		GetUniverseSchematicsSchematicId,
	)

	mockesi.NewRoute(
		"GetRouteOriginDestination",
		"Get",
		"/v1/route/{origin}/{destination}/",
		GetRouteOriginDestination,
	)

	mockesi.NewRoute(
		"GetSearch",
		"Get",
		"/v1/search/",
		GetSearch,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdAttributes",
		"Get",
		"/v1/characters/{character_id}/attributes/",
		GetCharactersCharacterIdAttributes,
	)

	mockesi.NewRoute(
		"GetSovereigntyCampaigns",
		"Get",
		"/v1/sovereignty/campaigns/",
		GetSovereigntyCampaigns,
	)

	mockesi.NewRoute(
		"GetSovereigntyMap",
		"Get",
		"/v1/sovereignty/map/",
		GetSovereigntyMap,
	)

	mockesi.NewRoute(
		"GetSovereigntyStructures",
		"Get",
		"/v1/sovereignty/structures/",
		GetSovereigntyStructures,
	)

	mockesi.NewRoute(
		"GetStatus",
		"Get",
		"/v1/status/",
		GetStatus,
	)

	mockesi.NewRoute(
		"GetUniverseAncestries",
		"Get",
		"/v1/universe/ancestries/",
		GetUniverseAncestries,
	)

	mockesi.NewRoute(
		"GetUniverseAsteroidBeltsAsteroidBeltId",
		"Get",
		"/v1/universe/asteroid_belts/{asteroid_belt_id}/",
		GetUniverseAsteroidBeltsAsteroidBeltId,
	)

	mockesi.NewRoute(
		"GetUniverseBloodlines",
		"Get",
		"/v1/universe/bloodlines/",
		GetUniverseBloodlines,
	)

	mockesi.NewRoute(
		"GetUniverseCategories",
		"Get",
		"/v1/universe/categories/",
		GetUniverseCategories,
	)

	mockesi.NewRoute(
		"GetUniverseCategoriesCategoryId",
		"Get",
		"/v1/universe/categories/{category_id}/",
		GetUniverseCategoriesCategoryId,
	)

	mockesi.NewRoute(
		"GetUniverseConstellations",
		"Get",
		"/v1/universe/constellations/",
		GetUniverseConstellations,
	)

	mockesi.NewRoute(
		"GetUniverseConstellationsConstellationId",
		"Get",
		"/v1/universe/constellations/{constellation_id}/",
		GetUniverseConstellationsConstellationId,
	)

	mockesi.NewRoute(
		"GetUniverseFactions",
		"Get",
		"/v1/universe/factions/",
		GetUniverseFactions,
	)

	mockesi.NewRoute(
		"GetUniverseGraphics",
		"Get",
		"/v1/universe/graphics/",
		GetUniverseGraphics,
	)

	mockesi.NewRoute(
		"GetUniverseGraphicsGraphicId",
		"Get",
		"/v1/universe/graphics/{graphic_id}/",
		GetUniverseGraphicsGraphicId,
	)

	mockesi.NewRoute(
		"GetUniverseGroups",
		"Get",
		"/v1/universe/groups/",
		GetUniverseGroups,
	)

	mockesi.NewRoute(
		"GetUniverseGroupsGroupId",
		"Get",
		"/v1/universe/groups/{group_id}/",
		GetUniverseGroupsGroupId,
	)

	mockesi.NewRoute(
		"GetUniverseMoonsMoonId",
		"Get",
		"/v1/universe/moons/{moon_id}/",
		GetUniverseMoonsMoonId,
	)

	mockesi.NewRoute(
		"GetUniversePlanetsPlanetId",
		"Get",
		"/v1/universe/planets/{planet_id}/",
		GetUniversePlanetsPlanetId,
	)

	mockesi.NewRoute(
		"GetUniverseRaces",
		"Get",
		"/v1/universe/races/",
		GetUniverseRaces,
	)

	mockesi.NewRoute(
		"GetUniverseRegions",
		"Get",
		"/v1/universe/regions/",
		GetUniverseRegions,
	)

	mockesi.NewRoute(
		"GetUniverseRegionsRegionId",
		"Get",
		"/v1/universe/regions/{region_id}/",
		GetUniverseRegionsRegionId,
	)

	mockesi.NewRoute(
		"GetUniverseStargatesStargateId",
		"Get",
		"/v1/universe/stargates/{stargate_id}/",
		GetUniverseStargatesStargateId,
	)

	mockesi.NewRoute(
		"GetUniverseStarsStarId",
		"Get",
		"/v1/universe/stars/{star_id}/",
		GetUniverseStarsStarId,
	)

	mockesi.NewRoute(
		"GetUniverseStationsStationId",
		"Get",
		"/v1/universe/stations/{station_id}/",
		GetUniverseStationsStationId,
	)

	mockesi.NewRoute(
		"GetUniverseStructures",
		"Get",
		"/v1/universe/structures/",
		GetUniverseStructures,
	)

	mockesi.NewRoute(
		"GetUniverseStructuresStructureId",
		"Get",
		"/v1/universe/structures/{structure_id}/",
		GetUniverseStructuresStructureId,
	)

	mockesi.NewRoute(
		"GetUniverseSystemJumps",
		"Get",
		"/v1/universe/system_jumps/",
		GetUniverseSystemJumps,
	)

	mockesi.NewRoute(
		"GetUniverseSystemKills",
		"Get",
		"/v1/universe/system_kills/",
		GetUniverseSystemKills,
	)

	mockesi.NewRoute(
		"GetUniverseSystems",
		"Get",
		"/v1/universe/systems/",
		GetUniverseSystems,
	)

	mockesi.NewRoute(
		"GetUniverseTypes",
		"Get",
		"/v1/universe/types/",
		GetUniverseTypes,
	)

	mockesi.NewRoute(
		"PostUniverseIds",
		"Post",
		"/v1/universe/ids/",
		PostUniverseIds,
	)

	mockesi.NewRoute(
		"PostUniverseNames",
		"Post",
		"/v1/universe/names/",
		PostUniverseNames,
	)

	mockesi.NewRoute(
		"PostUiAutopilotWaypoint",
		"Post",
		"/v1/ui/autopilot/waypoint/",
		PostUiAutopilotWaypoint,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowContract",
		"Post",
		"/v1/ui/openwindow/contract/",
		PostUiOpenwindowContract,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowInformation",
		"Post",
		"/v1/ui/openwindow/information/",
		PostUiOpenwindowInformation,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowMarketdetails",
		"Post",
		"/v1/ui/openwindow/marketdetails/",
		PostUiOpenwindowMarketdetails,
	)

	mockesi.NewRoute(
		"PostUiOpenwindowNewmail",
		"Post",
		"/v1/ui/openwindow/newmail/",
		PostUiOpenwindowNewmail,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdWallet",
		"Get",
		"/v1/characters/{character_id}/wallet/",
		GetCharactersCharacterIdWallet,
	)

	mockesi.NewRoute(
		"GetCharactersCharacterIdWalletTransactions",
		"Get",
		"/v1/characters/{character_id}/wallet/transactions/",
		GetCharactersCharacterIdWalletTransactions,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdWallets",
		"Get",
		"/v1/corporations/{corporation_id}/wallets/",
		GetCorporationsCorporationIdWallets,
	)

	mockesi.NewRoute(
		"GetCorporationsCorporationIdWalletsDivisionTransactions",
		"Get",
		"/v1/corporations/{corporation_id}/wallets/{division}/transactions/",
		GetCorporationsCorporationIdWalletsDivisionTransactions,
	)

	mockesi.NewRoute(
		"GetWars",
		"Get",
		"/v1/wars/",
		GetWars,
	)

	mockesi.NewRoute(
		"GetWarsWarId",
		"Get",
		"/v1/wars/{war_id}/",
		GetWarsWarId,
	)

	mockesi.NewRoute(
		"GetWarsWarIdKillmails",
		"Get",
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
