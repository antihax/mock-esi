package esiv1

/*
participant object
*/
type GetSovereigntyCampaignsParticipant struct {
	/*
	 alliance_id integer */
	AllianceId int32 `json:"alliance_id,omitempty"`
	/*
	 score number */
	Score float32 `json:"score,omitempty"`
}
