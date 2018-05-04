package esilegacy



/* 
200 ok object */
type GetSovereigntyCampaigns200Ok struct {
/*
	 Score for all attacking parties, only present in Defense Events.  */
	AttackersScore float32 `json:"attackers_score,omitempty"`
/*
	 Unique ID for this campaign. */
	CampaignId int32 `json:"campaign_id,omitempty"`
/*
	 The constellation in which the campaign will take place.  */
	ConstellationId int32 `json:"constellation_id,omitempty"`
/*
	 Defending alliance, only present in Defense Events  */
	DefenderId int32 `json:"defender_id,omitempty"`
/*
	 Score for the defending alliance, only present in Defense Events.  */
	DefenderScore float32 `json:"defender_score,omitempty"`
/*
	 Type of event this campaign is for. tcu_defense, ihub_defense and station_defense are referred to as \"Defense Events\", station_freeport as \"Freeport Events\".  */
	EventType string `json:"event_type,omitempty"`
/*
	 Alliance participating and their respective scores, only present in Freeport Events.  */
	Participants []GetSovereigntyCampaignsParticipant `json:"participants,omitempty"`
/*
	 The solar system the structure is located in.  */
	SolarSystemId int32 `json:"solar_system_id,omitempty"`
/*
	 Time the event is scheduled to start.  */
	StartTime time.Time `json:"start_time,omitempty"`
/*
	 The structure item ID that is related to this campaign.  */
	StructureId int64 `json:"structure_id,omitempty"`
}
