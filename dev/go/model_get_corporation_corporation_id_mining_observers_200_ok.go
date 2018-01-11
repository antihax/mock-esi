package esidev

/*
200 ok object */
type GetCorporationCorporationIdMiningObservers200Ok struct {
	/*
	 last_updated string */
	LastUpdated string `json:"last_updated,omitempty"`
	/*
	 The entity that was observing the asteroid field when it was mined.  */
	ObserverId int64 `json:"observer_id,omitempty"`
	/*
	 The category of the observing entity */
	ObserverType string `json:"observer_type,omitempty"`
}
