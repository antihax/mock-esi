package esilegacy



/* 
200 ok object */
type GetCharactersCharacterIdChatChannels200Ok struct {
/*
	 allowed array */
	Allowed []GetCharactersCharacterIdChatChannelsAllowed `json:"allowed,omitempty"`
/*
	 blocked array */
	Blocked []GetCharactersCharacterIdChatChannelsBlocked `json:"blocked,omitempty"`
/*
	 Unique channel ID. Always negative for player-created channels. Permanent (CCP created) channels have a positive ID, but don't appear in the API */
	ChannelId int32 `json:"channel_id,omitempty"`
/*
	 Normalized, unique string used to compare channel names */
	ComparisonKey string `json:"comparison_key,omitempty"`
/*
	 If this is a password protected channel */
	HasPassword bool `json:"has_password,omitempty"`
/*
	 Message of the day for this channel */
	Motd string `json:"motd,omitempty"`
/*
	 muted array */
	Muted []GetCharactersCharacterIdChatChannelsMuted `json:"muted,omitempty"`
/*
	 Displayed name of channel */
	Name string `json:"name,omitempty"`
/*
	 operators array */
	Operators []GetCharactersCharacterIdChatChannelsOperator `json:"operators,omitempty"`
/*
	 owner_id integer */
	OwnerId int32 `json:"owner_id,omitempty"`
}
