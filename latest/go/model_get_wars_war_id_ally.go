package esilatest

/*
ally object */
type GetWarsWarIdAlly struct {
	/*
	 Alliance ID if and only if this ally is an alliance */
	AllianceId int32 `json:"alliance_id,omitempty"`
	/*
	 Corporation ID if and only if this ally is a corporation */
	CorporationId int32 `json:"corporation_id,omitempty"`
}
