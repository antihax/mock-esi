package esiv1

/*
200 ok object
*/
type GetCorporationsCorporationIdStandings200Ok struct {
	/*
	 from_id integer */
	FromId int32 `json:"from_id,omitempty"`
	/*
	 from_type string */
	FromType string `json:"from_type,omitempty"`
	/*
	 standing number */
	Standing float32 `json:"standing,omitempty"`
}
