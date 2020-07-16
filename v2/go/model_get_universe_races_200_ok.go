package esiv2

/*
200 ok object */
type GetUniverseRaces200Ok struct {
	/*
	 The alliance generally associated with this race */
	AllianceId int32 `json:"alliance_id,omitempty"`
	/*
	 description string */
	Description string `json:"description,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 race_id integer */
	RaceId int32 `json:"race_id,omitempty"`
}
