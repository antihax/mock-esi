package esidev

import "time"

/*
200 ok object */
type GetCharactersCharacterIdAgentsResearch200Ok struct {
	/*
	 agent_id integer */
	AgentId int32 `json:"agent_id,omitempty"`
	/*
	 points_per_day number */
	PointsPerDay float32 `json:"points_per_day,omitempty"`
	/*
	 remainder_points number */
	RemainderPoints float32 `json:"remainder_points,omitempty"`
	/*
	 skill_type_id integer */
	SkillTypeId int32 `json:"skill_type_id,omitempty"`
	/*
	 started_at string */
	StartedAt time.Time `json:"started_at,omitempty"`
}
