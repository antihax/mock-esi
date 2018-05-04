package esilatest



/* 
200 ok object */
type GetCharactersCharacterIdNotificationsContacts200Ok struct {
/*
	 message string */
	Message string `json:"message,omitempty"`
/*
	 notification_id integer */
	NotificationId int32 `json:"notification_id,omitempty"`
/*
	 send_date string */
	SendDate time.Time `json:"send_date,omitempty"`
/*
	 sender_character_id integer */
	SenderCharacterId int32 `json:"sender_character_id,omitempty"`
/*
	 A number representing the standing level the receiver has been added at by the sender. The standing levels are as follows: -10 -> Terrible | -5 -> Bad |  0 -> Neutral |  5 -> Good |  10 -> Excellent */
	StandingLevel float32 `json:"standing_level,omitempty"`
}
