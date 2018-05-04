package esilatest



/* 
200 ok object */
type GetOpportunitiesGroupsGroupIdOk struct {
/*
	 The groups that are connected to this group on the opportunities map */
	ConnectedGroups []int32 `json:"connected_groups,omitempty"`
/*
	 description string */
	Description string `json:"description,omitempty"`
/*
	 group_id integer */
	GroupId int32 `json:"group_id,omitempty"`
/*
	 name string */
	Name string `json:"name,omitempty"`
/*
	 notification string */
	Notification string `json:"notification,omitempty"`
/*
	 Tasks need to complete for this group */
	RequiredTasks []int32 `json:"required_tasks,omitempty"`
}
