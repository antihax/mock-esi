package esidev

/*
200 ok object */
type GetFleetsFleetIdWings200Ok struct {
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 id integer */
	Id int64 `json:"id,omitempty"`
	/*
	 squads array */
	Squads []GetFleetsFleetIdWingsSquad `json:"squads,omitempty"`
}
