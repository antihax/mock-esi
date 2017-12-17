package esidev

/*
ally object */
type GetWarsWarIdAlly struct {
	/*
	 Corporation ID if and only if this ally is a corporation */
	CorporationId int32 `json:"corporation_id,omitempty"`
	/*
	 Alliance ID if and only if this ally is an alliance */
	AllianceId int32 `json:"alliance_id,omitempty"`
}
