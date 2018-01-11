package esilegacy

/*
200 ok object */
type GetOpportunitiesTasksTaskIdOk struct {
	/*
	 task_id integer */
	TaskId int32 `json:"task_id,omitempty"`
	/*
	 name string */
	Name string `json:"name,omitempty"`
	/*
	 description string */
	Description string `json:"description,omitempty"`
	/*
	 notification string */
	Notification string `json:"notification,omitempty"`
}
