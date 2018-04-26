package esiv1

import "time"

/*
200 ok object */
type GetCorporationsCorporationIdMedals200Ok struct {
	/*
	 created_at string */
	CreatedAt time.Time `json:"created_at,omitempty"`
	/*
	 ID of the character who created this medal */
	CreatorId int32 `json:"creator_id,omitempty"`
	/*
	 description string */
	Description string `json:"description,omitempty"`
	/*
	 medal_id integer */
	MedalId int32 `json:"medal_id,omitempty"`
	/*
	 title string */
	Title string `json:"title,omitempty"`
}
