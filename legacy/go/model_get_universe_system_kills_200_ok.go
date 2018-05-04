package esilegacy



/* 
200 ok object */
type GetUniverseSystemKills200Ok struct {
/*
	 Number of NPC ships killed in this system */
	NpcKills int32 `json:"npc_kills,omitempty"`
/*
	 Number of pods killed in this system */
	PodKills int32 `json:"pod_kills,omitempty"`
/*
	 Number of player and NPC ships killed in this system */
	ShipKills int32 `json:"ship_kills,omitempty"`
/*
	 system_id integer */
	SystemId int32 `json:"system_id,omitempty"`
}
