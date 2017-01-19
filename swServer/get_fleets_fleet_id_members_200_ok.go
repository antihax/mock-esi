package swServer

import "time"
var _ time.Time

/* 
200 ok object */
type GetFleetsFleetIdMembers200Ok struct {
/*
	 character_id integer */
	character_id int32 `json:"character_id,omitempty"`
/*
	 join_time string */
	join_time time.Time `json:"join_time,omitempty"`
/*
	 Member’s role in fleet */
	role string `json:"role,omitempty"`
/*
	 Localized role names */
	role_name string `json:"role_name,omitempty"`
/*
	 ship_type_id integer */
	ship_type_id int32 `json:"ship_type_id,omitempty"`
/*
	 Solar system the member is located in */
	solar_system_id int32 `json:"solar_system_id,omitempty"`
/*
	 ID of the squad the member is in. If not applicable, will be set to -1 */
	squad_id int64 `json:"squad_id,omitempty"`
/*
	 Station in which the member is docked in, if applicable */
	station_id int64 `json:"station_id,omitempty"`
/*
	 Whether the member take fleet warps */
	takes_fleet_warp bool `json:"takes_fleet_warp,omitempty"`
/*
	 ID of the wing the member is in. If not applicable, will be set to -1 */
	wing_id int64 `json:"wing_id,omitempty"`
}
