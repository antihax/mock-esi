package esilegacy

/*
alliance object */
type GetCorporationsCorporationIdAlliancehistoryAlliance struct {
	/*
	 alliance_id integer */
	AllianceId int32 `json:"alliance_id,omitempty"`
	/*
	 True if the alliance has been deleted */
	IsDeleted bool `json:"is_deleted,omitempty"`
}
