package esilatest

/*
new_settings object */
type PutFleetsFleetIdNewSettings struct {
	/*
	 New fleet MOTD in CCP flavoured HTML */
	Motd string `json:"motd,omitempty"`
	/*
	 Should free-move be enabled in the fleet */
	IsFreeMove bool `json:"is_free_move,omitempty"`
}
