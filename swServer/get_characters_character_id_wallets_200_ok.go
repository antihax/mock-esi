package swServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetCharactersCharacterIdWallets200Ok struct {
/*
	 Wallet's balance in ISK hundredths. */
	balance int64 `json:"balance,omitempty"`
/*
	 wallet_id integer */
	wallet_id int32 `json:"wallet_id,omitempty"`
}
