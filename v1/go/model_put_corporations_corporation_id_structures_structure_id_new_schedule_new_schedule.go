package esiv1

/*
new_schedule object */
type PutCorporationsCorporationIdStructuresStructureIdNewScheduleNewSchedule struct {
	/*
	 Day of the week, zero-indexed to Monday */
	Day int32 `json:"day,omitempty"`
	/*
	 Hour of the day evetime, zero-indexed to midnight */
	Hour int32 `json:"hour,omitempty"`
}
