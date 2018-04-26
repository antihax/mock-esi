package esidev

/*
200 ok object */
type GetFleetsFleetIdWings200Ok struct {
	/*
	 id integer */
	Id int64 `json:"id,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 squads array */
	Squads []GetFleetsFleetIdWingsSquad `json:"squads,omitempty"`
}
