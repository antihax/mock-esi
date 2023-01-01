package esidev

import "time"

/*
200 ok object
*/
type GetCharactersCharacterIdOnlineOk struct {
	/*
	 Timestamp of the last login */
	LastLogin time.Time `json:"last_login,omitempty"`
	/*
	 Timestamp of the last logout */
	LastLogout time.Time `json:"last_logout,omitempty"`
	/*
	 Total number of times the character has logged in */
	Logins int32 `json:"logins,omitempty"`
	/*
	 If the character is online */
	Online bool `json:"online,omitempty"`
}
