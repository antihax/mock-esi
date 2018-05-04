package esilatest



/* 
inventory object */
type GetCharactersCharacterIdStatsInventory struct {
/*
	 abandon_loot_quantity integer */
	AbandonLootQuantity int64 `json:"abandon_loot_quantity,omitempty"`
/*
	 trash_item_quantity integer */
	TrashItemQuantity int64 `json:"trash_item_quantity,omitempty"`
}
