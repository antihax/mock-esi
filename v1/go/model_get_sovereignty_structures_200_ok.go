package esiv1



/* 
200 ok object */
type GetSovereigntyStructures200Ok struct {
/*
	 The alliance that owns the structure.  */
	AllianceId int32 `json:"alliance_id,omitempty"`
/*
	 Solar system in which the structure is located.  */
	SolarSystemId int32 `json:"solar_system_id,omitempty"`
/*
	 Unique item ID for this structure. */
	StructureId int64 `json:"structure_id,omitempty"`
/*
	 A reference to the type of structure this is.  */
	StructureTypeId int32 `json:"structure_type_id,omitempty"`
/*
	 The occupancy level for the next or current vulnerability window. This takes into account all development indexes and capital system bonuses. Also known as Activity Defense Multiplier from in the client. It increases the time that attackers must spend using their entosis links on the structure.  */
	VulnerabilityOccupancyLevel float32 `json:"vulnerability_occupancy_level,omitempty"`
/*
	 The time at which the next or current vulnerability window ends. At the end of a vulnerability window the next window is recalculated and locked in along with the vulnerabilityOccupancyLevel. If the structure is not in 100% entosis control of the defender, it will go in to 'overtime' and stay vulnerable for as long as that situation persists. Only once the defenders have 100% entosis control and has the vulnerableEndTime passed does the vulnerability interval expire and a new one is calculated.  */
	VulnerableEndTime time.Time `json:"vulnerable_end_time,omitempty"`
/*
	 The next time at which the structure will become vulnerable. Or the start time of the current window if current time is between this and vulnerableEndTime.  */
	VulnerableStartTime time.Time `json:"vulnerable_start_time,omitempty"`
}
