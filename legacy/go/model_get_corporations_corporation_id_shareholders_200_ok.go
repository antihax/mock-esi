package esilegacy

/*
200 ok object */
type GetCorporationsCorporationIdShareholders200Ok struct {
	/*
	 share_count integer */
	ShareCount int64 `json:"share_count,omitempty"`
	/*
	 shareholder_id integer */
	ShareholderId int32 `json:"shareholder_id,omitempty"`
	/*
	 shareholder_type string */
	ShareholderType string `json:"shareholder_type,omitempty"`
}
