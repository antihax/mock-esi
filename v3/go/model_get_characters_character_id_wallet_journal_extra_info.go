package esiv3

/*
Extra information for different type of transaction */
type GetCharactersCharacterIdWalletJournalExtraInfo struct {
	/*
	 location_id integer */
	LocationId int64 `json:"location_id,omitempty"`
	/*
	 transaction_id integer */
	TransactionId int64 `json:"transaction_id,omitempty"`
	/*
	 npc_name string */
	NpcName string `json:"npc_name,omitempty"`
	/*
	 npc_id integer */
	NpcId int32 `json:"npc_id,omitempty"`
	/*
	 destroyed_ship_type_id integer */
	DestroyedShipTypeId int32 `json:"destroyed_ship_type_id,omitempty"`
	/*
	 character_id integer */
	CharacterId int32 `json:"character_id,omitempty"`
	/*
	 corporation_id integer */
	CorporationId int32 `json:"corporation_id,omitempty"`
	/*
	 alliance_id integer */
	AllianceId int32 `json:"alliance_id,omitempty"`
	/*
	 job_id integer */
	JobId int32 `json:"job_id,omitempty"`
	/*
	 contract_id integer */
	ContractId int32 `json:"contract_id,omitempty"`
	/*
	 system_id integer */
	SystemId int32 `json:"system_id,omitempty"`
	/*
	 planet_id integer */
	PlanetId int32 `json:"planet_id,omitempty"`
}
