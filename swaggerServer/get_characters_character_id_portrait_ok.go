package swaggerServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersCharacterIdPortraitOk struct {
/*
	 px128x128 string */
	px128x128 string `json:"px128x128,omitempty"`
/*
	 px256x256 string */
	px256x256 string `json:"px256x256,omitempty"`
/*
	 px512x512 string */
	px512x512 string `json:"px512x512,omitempty"`
/*
	 px64x64 string */
	px64x64 string `json:"px64x64,omitempty"`
}
