package esilatest

/*
200 ok object
*/
type GetFleetsFleetIdOk struct {
	/*
	 Is free-move enabled */
	IsFreeMove bool `json:"is_free_move,omitempty"`
	/*
	 Does the fleet have an active fleet advertisement */
	IsRegistered bool `json:"is_registered,omitempty"`
	/*
	 Is EVE Voice enabled */
	IsVoiceEnabled bool `json:"is_voice_enabled,omitempty"`
	/*
	 Fleet MOTD in CCP flavoured HTML */
	Motd string `json:"motd,omitempty"`
}
