package esidev

/*
200 ok object */
type GetFwSystems200Ok struct {
	/*
	 contested boolean */
	Contested bool `json:"contested,omitempty"`
	/*
	 occupier_faction_id integer */
	OccupierFactionId int32 `json:"occupier_faction_id,omitempty"`
	/*
	 owner_faction_id integer */
	OwnerFactionId int32 `json:"owner_faction_id,omitempty"`
	/*
	 solar_system_id integer */
	SolarSystemId int32 `json:"solar_system_id,omitempty"`
	/*
	 victory_points integer */
	VictoryPoints int32 `json:"victory_points,omitempty"`
	/*
	 victory_points_threshold integer */
	VictoryPointsThreshold int32 `json:"victory_points_threshold,omitempty"`
}
