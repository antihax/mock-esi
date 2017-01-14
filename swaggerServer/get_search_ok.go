package swaggerServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetSearchOk struct {
/*
	 agent array */
	agent []int32 `json:"agent,omitempty"`
/*
	 alliance array */
	alliance []int32 `json:"alliance,omitempty"`
/*
	 character array */
	character []int32 `json:"character,omitempty"`
/*
	 constellation array */
	constellation []int32 `json:"constellation,omitempty"`
/*
	 corporation array */
	corporation []int32 `json:"corporation,omitempty"`
/*
	 faction array */
	faction []int32 `json:"faction,omitempty"`
/*
	 inventorytype array */
	inventorytype []int32 `json:"inventorytype,omitempty"`
/*
	 region array */
	region []int32 `json:"region,omitempty"`
/*
	 solarsystem array */
	solarsystem []int32 `json:"solarsystem,omitempty"`
/*
	 station array */
	station []int32 `json:"station,omitempty"`
/*
	 wormhole array */
	wormhole []int32 `json:"wormhole,omitempty"`
}
