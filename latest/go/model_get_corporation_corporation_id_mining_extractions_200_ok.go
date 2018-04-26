package esilatest

import "time"

/*
200 ok object */
type GetCorporationCorporationIdMiningExtractions200Ok struct {
	/*
	 The time at which the chunk being extracted will arrive and can be fractured by the moon mining drill.  */
	ChunkArrivalTime time.Time `json:"chunk_arrival_time,omitempty"`
	/*
	 The time at which the current extraction was initiated.  */
	ExtractionStartTime time.Time `json:"extraction_start_time,omitempty"`
	/*
	 moon_id integer */
	MoonId int32 `json:"moon_id,omitempty"`
	/*
	 The time at which the chunk being extracted will naturally fracture if it is not first fractured by the moon mining drill.  */
	NaturalDecayTime time.Time `json:"natural_decay_time,omitempty"`
	/*
	 structure_id integer */
	StructureId int64 `json:"structure_id,omitempty"`
}
