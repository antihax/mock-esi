package esilegacy



/* 
Extra information for different type of transaction */
type GetCorporationsCorporationIdWalletsDivisionJournalExtraInfo struct {
/*
	 alliance_id integer */
	AllianceId int32 `json:"alliance_id,omitempty"`
/*
	 character_id integer */
	CharacterId int32 `json:"character_id,omitempty"`
/*
	 contract_id integer */
	ContractId int32 `json:"contract_id,omitempty"`
/*
	 corporation_id integer */
	CorporationId int32 `json:"corporation_id,omitempty"`
/*
	 destroyed_ship_type_id integer */
	DestroyedShipTypeId int32 `json:"destroyed_ship_type_id,omitempty"`
/*
	 job_id integer */
	JobId int32 `json:"job_id,omitempty"`
/*
	 location_id integer */
	LocationId int64 `json:"location_id,omitempty"`
/*
	 npc_id integer */
	NpcId int32 `json:"npc_id,omitempty"`
/*
	 npc_name string */
	NpcName string `json:"npc_name,omitempty"`
/*
	 planet_id integer */
	PlanetId int32 `json:"planet_id,omitempty"`
/*
	 system_id integer */
	SystemId int32 `json:"system_id,omitempty"`
/*
	 transaction_id integer */
	TransactionId int64 `json:"transaction_id,omitempty"`
}
