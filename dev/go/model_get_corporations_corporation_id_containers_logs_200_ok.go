package esidev

import "time"

/*
200 ok object
*/
type GetCorporationsCorporationIdContainersLogs200Ok struct {
	/*
	 action string */
	Action string `json:"action,omitempty"`
	/*
	 ID of the character who performed the action. */
	CharacterId int32 `json:"character_id,omitempty"`
	/*
	 ID of the container */
	ContainerId int64 `json:"container_id,omitempty"`
	/*
	 Type ID of the container */
	ContainerTypeId int32 `json:"container_type_id,omitempty"`
	/*
	 location_flag string */
	LocationFlag string `json:"location_flag,omitempty"`
	/*
	 location_id integer */
	LocationId int64 `json:"location_id,omitempty"`
	/*
	 Timestamp when this log was created */
	LoggedAt time.Time `json:"logged_at,omitempty"`
	/*
	 new_config_bitmask integer */
	NewConfigBitmask int32 `json:"new_config_bitmask,omitempty"`
	/*
	 old_config_bitmask integer */
	OldConfigBitmask int32 `json:"old_config_bitmask,omitempty"`
	/*
	 Type of password set if action is of type SetPassword or EnterPassword */
	PasswordType string `json:"password_type,omitempty"`
	/*
	 Quantity of the item being acted upon */
	Quantity int32 `json:"quantity,omitempty"`
	/*
	 Type ID of the item being acted upon */
	TypeId int32 `json:"type_id,omitempty"`
}
