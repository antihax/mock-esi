package swaggerServer

import (
	"net/http"

)

func GetCharactersCharacterIdCalendar(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "event_date" : "2016-06-26T20:00:00Z",
  "event_id" : 1386435,
  "event_response" : "accepted",
  "importance" : 0,
  "title" : "o7 The EVE Online Show"
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetCharactersCharacterIdCalendarEventId(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "date" : "2016-06-26T21:00:00Z",
  "duration" : 60,
  "event_id" : 1386435,
  "importance" : 1,
  "owner_id" : 1,
  "owner_name" : "EVE System",
  "owner_type" : "eve_server",
  "response" : "Undecided",
  "text" : "o7: The EVE Online Show features latest developer news, fast paced action, community overviews and a lot more with CCP Guard and CCP Mimic. Join the thrilling o7 live broadcast at 20:00 EVE time (=UTC) on <a href=\"http://www.twitch.tv/ccp\">EVE TV</a>. Don't miss it!",
  "title" : "o7 The EVE Online Show"
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func PutCharactersCharacterIdCalendarEventId(w http.ResponseWriter, r *http.Request) {

	j := (``)
		w.Header().Set("Content-Type", "")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


