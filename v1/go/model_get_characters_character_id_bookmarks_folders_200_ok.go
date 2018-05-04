package esiv1

/*
200 ok object */
type GetCharactersCharacterIdBookmarksFolders200Ok struct {
	/*
	 folder_id integer */
	FolderId int32 `json:"folder_id,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 owner_id integer */
	OwnerId int32 `json:"owner_id,omitempty"`
}
