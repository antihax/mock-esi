package esilegacy

import "time"

/*
200 ok object */
type GetCharactersCharacterIdNotifications200Ok struct {
	/*
	 is_read boolean */
	IsRead bool `json:"is_read,omitempty"`
	/*
	 notification_id integer */
	NotificationId int64 `json:"notification_id,omitempty"`
	/*
	 sender_id integer */
	SenderId int32 `json:"sender_id,omitempty"`
	/*
	 sender_type string */
	SenderType string `json:"sender_type,omitempty"`
	/*
	 text string */
	Text string `json:"text,omitempty"`
	/*
	 timestamp string */
	Timestamp time.Time `json:"timestamp,omitempty"`
	/*
	 type string */
	Type_ string `json:"type,omitempty"`
}
