package swaggerServer

import (
	"net/http"

)

func DeleteCharactersCharacterIdMailLabelsLabelId(w http.ResponseWriter, r *http.Request) {

	j := (``)
		w.Header().Set("Content-Type", "")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func DeleteCharactersCharacterIdMailMailId(w http.ResponseWriter, r *http.Request) {

	j := (``)
		w.Header().Set("Content-Type", "")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetCharactersCharacterIdMail(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "from" : 90000001,
  "is_read" : true,
  "labels" : [ 3 ],
  "mail_id" : 7,
  "recipients" : [ {
    "recipient_id" : 90000002,
    "recipient_type" : "character"
  } ],
  "subject" : "Title for EVE Mail",
  "timestamp" : "2015-09-30T16:07:00Z"
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetCharactersCharacterIdMailLabels(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "labels" : [ {
    "color_hex" : "#660066",
    "label_id" : 16,
    "name" : "PINK",
    "unread_count" : 4
  }, {
    "color_hex" : "#ffffff",
    "label_id" : 17,
    "name" : "WHITE",
    "unread_count" : 1
  } ],
  "total_unread_count" : 5
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetCharactersCharacterIdMailLists(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "mailing_list_id" : 1,
  "name" : "test_mailing_list"
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetCharactersCharacterIdMailMailId(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "body" : "blah blah blah",
  "from" : 90000001,
  "labels" : [ 2, 32 ],
  "read" : false,
  "subject" : "test",
  "timestamp" : "2015-09-30T16:07:00Z"
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func PostCharactersCharacterIdMail(w http.ResponseWriter, r *http.Request) {

	j := (`13`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func PostCharactersCharacterIdMailLabels(w http.ResponseWriter, r *http.Request) {

	j := (`128`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func PutCharactersCharacterIdMailMailId(w http.ResponseWriter, r *http.Request) {

	j := (``)
		w.Header().Set("Content-Type", "")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


