package swS

import "time"
var _ time.Time

/* 
new_mail object */
type PostUiOpenwindowNewmailNewMail struct {
/*
	 body string */
	body string `json:"body,omitempty"`
/*
	 recipients array */
	recipients []int32 `json:"recipients,omitempty"`
/*
	 subject string */
	subject string `json:"subject,omitempty"`
/*
	 to_corp_or_alliance_id integer */
	to_corp_or_alliance_id int32 `json:"to_corp_or_alliance_id,omitempty"`
/*
	 Corporations, alliances and mailing lists are all types of mailing groups. You may only send to one mailing group, at a time, so you may fill out either this field or the to_corp_or_alliance_ids field */
	to_mailing_list_id int32 `json:"to_mailing_list_id,omitempty"`
}
