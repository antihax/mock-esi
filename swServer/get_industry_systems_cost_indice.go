package swServer

import "time"
var _ time.Time

/* 
cost_indice object */
type GetIndustrySystemsCostIndice struct {
/*
	 activity string */
	activity string `json:"activity,omitempty"`
/*
	 cost_index number */
	cost_index float32 `json:"cost_index,omitempty"`
}
