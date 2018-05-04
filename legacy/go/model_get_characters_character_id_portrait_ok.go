package esilegacy

/*
200 ok object */
type GetCharactersCharacterIdPortraitOk struct {
	/*
	 128x128 string */
	Var128x128 string `json:"128x128,omitempty"`
	/*
	 256x256 string */
	Var256x256 string `json:"256x256,omitempty"`
	/*
	 512x512 string */
	Var512x512 string `json:"512x512,omitempty"`
	/*
	 64x64 string */
	Var64x64 string `json:"64x64,omitempty"`
}
