package esiv2

/*
pve object */
type GetCharactersCharacterIdStatsPve struct {
	/*
	 dungeons_completed_agent integer */
	DungeonsCompletedAgent int64 `json:"dungeons_completed_agent,omitempty"`
	/*
	 dungeons_completed_distribution integer */
	DungeonsCompletedDistribution int64 `json:"dungeons_completed_distribution,omitempty"`
	/*
	 missions_succeeded integer */
	MissionsSucceeded int64 `json:"missions_succeeded,omitempty"`
	/*
	 missions_succeeded_epic_arc integer */
	MissionsSucceededEpicArc int64 `json:"missions_succeeded_epic_arc,omitempty"`
}
