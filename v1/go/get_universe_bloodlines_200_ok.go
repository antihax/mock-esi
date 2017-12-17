package esiv1

/*
200 ok object */
type GetUniverseBloodlines200Ok struct {
	/*
	 bloodline_id integer */
	BloodlineId int32 `json:"bloodline_id,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 description string */
	Description string `json:"description,omitempty"`
	/*
	 race_id integer */
	RaceId int32 `json:"race_id,omitempty"`
	/*
	 ship_type_id integer */
	ShipTypeId int32 `json:"ship_type_id,omitempty"`
	/*
	 corporation_id integer */
	CorporationId int32 `json:"corporation_id,omitempty"`
	/*
	 perception integer */
	Perception int32 `json:"perception,omitempty"`
	/*
	 willpower integer */
	Willpower int32 `json:"willpower,omitempty"`
	/*
	 charisma integer */
	Charisma int32 `json:"charisma,omitempty"`
	/*
	 memory integer */
	Memory int32 `json:"memory,omitempty"`
	/*
	 intelligence integer */
	Intelligence int32 `json:"intelligence,omitempty"`
}
