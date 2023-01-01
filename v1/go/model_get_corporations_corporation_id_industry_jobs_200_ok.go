package esiv1

import "time"

/*
200 ok object
*/
type GetCorporationsCorporationIdIndustryJobs200Ok struct {
	/*
	 Job activity ID */
	ActivityId int32 `json:"activity_id,omitempty"`
	/*
	 blueprint_id integer */
	BlueprintId int64 `json:"blueprint_id,omitempty"`
	/*
	 Location ID of the location from which the blueprint was installed. Normally a station ID, but can also be an asset (e.g. container) or corporation facility */
	BlueprintLocationId int64 `json:"blueprint_location_id,omitempty"`
	/*
	 blueprint_type_id integer */
	BlueprintTypeId int32 `json:"blueprint_type_id,omitempty"`
	/*
	 ID of the character which completed this job */
	CompletedCharacterId int32 `json:"completed_character_id,omitempty"`
	/*
	 Date and time when this job was completed */
	CompletedDate time.Time `json:"completed_date,omitempty"`
	/*
	 The sume of job installation fee and industry facility tax */
	Cost float64 `json:"cost,omitempty"`
	/*
	 Job duration in seconds */
	Duration int32 `json:"duration,omitempty"`
	/*
	 Date and time when this job finished */
	EndDate time.Time `json:"end_date,omitempty"`
	/*
	 ID of the facility where this job is running */
	FacilityId int64 `json:"facility_id,omitempty"`
	/*
	 ID of the character which installed this job */
	InstallerId int32 `json:"installer_id,omitempty"`
	/*
	 Unique job ID */
	JobId int32 `json:"job_id,omitempty"`
	/*
	 Number of runs blueprint is licensed for */
	LicensedRuns int32 `json:"licensed_runs,omitempty"`
	/*
	 ID of the location for the industry facility */
	LocationId int64 `json:"location_id,omitempty"`
	/*
	 Location ID of the location to which the output of the job will be delivered. Normally a station ID, but can also be a corporation facility */
	OutputLocationId int64 `json:"output_location_id,omitempty"`
	/*
	 Date and time when this job was paused (i.e. time when the facility where this job was installed went offline) */
	PauseDate time.Time `json:"pause_date,omitempty"`
	/*
	 Chance of success for invention */
	Probability float32 `json:"probability,omitempty"`
	/*
	 Type ID of product (manufactured, copied or invented) */
	ProductTypeId int32 `json:"product_type_id,omitempty"`
	/*
	 Number of runs for a manufacturing job, or number of copies to make for a blueprint copy */
	Runs int32 `json:"runs,omitempty"`
	/*
	 Date and time when this job started */
	StartDate time.Time `json:"start_date,omitempty"`
	/*
	 status string */
	Status string `json:"status,omitempty"`
	/*
	 Number of successful runs for this job. Equal to runs unless this is an invention job */
	SuccessfulRuns int32 `json:"successful_runs,omitempty"`
}
