package swaggerServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetAlliancesAllianceIdIconsOk struct {
/*
	 px128x128 string */
	px128x128 string `json:"px128x128,omitempty"`
/*
	 px64x64 string */
	px64x64 string `json:"px64x64,omitempty"`
}
