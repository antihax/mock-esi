package esiv1

/*
new_settings object */
type PutFleetsFleetIdNewSettings struct {
	/*
	 Should free-move be enabled in the fleet */
	IsFreeMove bool `json:"is_free_move,omitempty"`
	/*
	 New fleet MOTD in CCP flavoured HTML */
	Motd string `json:"motd,omitempty"`
}
