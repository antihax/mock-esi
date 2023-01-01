package esilatest

/*
200 ok object
*/
type GetCorporationsCorporationIdDivisionsOk struct {
	/*
	 hangar array */
	Hangar []GetCorporationsCorporationIdDivisionsHangarHangar `json:"hangar,omitempty"`
	/*
	 wallet array */
	Wallet []GetCorporationsCorporationIdDivisionsWalletWallet `json:"wallet,omitempty"`
}
