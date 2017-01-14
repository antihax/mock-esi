package swaggerServer

import (
	"net/http"
	"github.com/gorilla/mux"
)

var _ = mux.NewRouter

func GetCharactersCharacterIdCalendar(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		characterId int64
		fromEvent int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "event_date" : "2016-06-26T20:00:00Z",
  "event_id" : 1386435,
  "event_response" : "accepted",
  "importance" : 0,
  "title" : "o7 The EVE Online Show"
} ]`
	vars := mux.Vars(r)
	localV, err = processParameters(characterId, vars["characterId"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	characterId = localV.(int64)
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("fromEvent") != "" {
		localV, err = processParameters(fromEvent, r.Form.Get("fromEvent"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		fromEvent = localV.(int32)
	}
	if r.Form.Get("datasource") != "" {
		localV, err = processParameters(datasource, r.Form.Get("datasource"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		datasource = localV.(string)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}

func GetCharactersCharacterIdCalendarEventId(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		characterId int64
		eventId int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
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
}`
	vars := mux.Vars(r)
	localV, err = processParameters(characterId, vars["characterId"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	characterId = localV.(int64)
	localV, err = processParameters(eventId, vars["eventId"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	eventId = localV.(int32)
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("datasource") != "" {
		localV, err = processParameters(datasource, r.Form.Get("datasource"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		datasource = localV.(string)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}

func PutCharactersCharacterIdCalendarEventId(w http.ResponseWriter, r *http.Request) {

	var (
		localV interface{}
		err error
		characterId int32
		eventId int32
		datasource string
	)
	// shut up warnings
	localV = localV
	err = err

	j := ``
	vars := mux.Vars(r)
	localV, err = processParameters(characterId, vars["characterId"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	characterId = localV.(int32)
	localV, err = processParameters(eventId, vars["eventId"])
	if err != nil {
		errorOut(w, r, err)
		return
	}
	eventId = localV.(int32)
	if err := r.ParseForm(); err != nil {
		errorOut(w, r, err)
		return
	}
	if r.Form.Get("datasource") != "" {
		localV, err = processParameters(datasource, r.Form.Get("datasource"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		datasource = localV.(string)
	}

	w.Header().Set("Content-Type", "")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}


