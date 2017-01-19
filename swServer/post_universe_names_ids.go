package swServer

import "time"
var _ time.Time

/* 
ids schema */
type PostUniverseNamesIds struct {
/*
	 ids array */
	ids []int32 `json:"ids,omitempty"`
}
