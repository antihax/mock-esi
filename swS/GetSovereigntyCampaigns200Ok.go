package swS

import "time"
var _ time.Time

/* 
200 ok object */
type GetSovereigntyCampaigns200Ok struct {
/*
	 Score for all attacking parties, only present in Defense Events.
 */
	attackers_score float32 `json:"attackers_score,omitempty"`
/*
	 Unique ID for this campaign. */
	campaign_id int32 `json:"campaign_id,omitempty"`
/*
	 The constellation in which the campaign will take place.
 */
	constellation_id int32 `json:"constellation_id,omitempty"`
/*
	 Defending alliance, only present in Defense Events
 */
	defender_id int32 `json:"defender_id,omitempty"`
/*
	 Score for the defending alliance, only present in Defense Events.
 */
	defender_score float32 `json:"defender_score,omitempty"`
/*
	 Type of event this campaign is for. tcu_defense, ihub_defense and station_defense are referred to as "Defense Events", station_freeport as "Freeport Events".
 */
	event_type string `json:"event_type,omitempty"`
/*
	 Alliance participating and their respective scores, only present in Freeport Events.
 */
	participants []GetSovereigntyCampaignsParticipant `json:"participants,omitempty"`
/*
	 The solar system the structure is located in.
 */
	solar_system_id int32 `json:"solar_system_id,omitempty"`
/*
	 Time the event is scheduled to start.
 */
	start_time time.Time `json:"start_time,omitempty"`
/*
	 The structure item ID that is related to this campaign.
 */
	structure_id int64 `json:"structure_id,omitempty"`
}
