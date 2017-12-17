package esilatest

import "time"

/*
200 ok object */
type GetStatusOk struct {
	/*
	 Server start timestamp */
	StartTime time.Time `json:"start_time,omitempty"`
	/*
	 Current online player count */
	Players int32 `json:"players,omitempty"`
	/*
	 Running version as string */
	ServerVersion string `json:"server_version,omitempty"`
	/*
	 If the server is in VIP mode */
	Vip bool `json:"vip,omitempty"`
}
