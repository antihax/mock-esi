package esiv3

/*
200 ok object
*/
type GetCharactersCharacterIdMailLabelsOk struct {
	/*
	 labels array */
	Labels []GetCharactersCharacterIdMailLabelsLabel `json:"labels,omitempty"`
	/*
	 total_unread_count integer */
	TotalUnreadCount int32 `json:"total_unread_count,omitempty"`
}
