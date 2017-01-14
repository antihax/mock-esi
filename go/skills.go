package swaggerServer

import (
	"net/http"

)

func GetCharactersCharacterIdSkillqueue(w http.ResponseWriter, r *http.Request) {

	j := (`[ {
  "finish_date" : "2016-06-29T10:47:00Z",
  "finished_level" : 3,
  "queue_position" : 0,
  "skill_id" : 1,
  "start_date" : "2016-06-29T10:46:00Z"
}, {
  "finish_date" : "2016-07-15T10:47:00Z",
  "finished_level" : 4,
  "queue_position" : 1,
  "skill_id" : 1,
  "start_date" : "2016-06-29T10:47:00Z"
}, {
  "finish_date" : "2016-08-30T10:47:00Z",
  "finished_level" : 2,
  "queue_position" : 2,
  "skill_id" : 2,
  "start_date" : "2016-07-15T10:47:00Z"
} ]`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}

func GetCharactersCharacterIdSkills(w http.ResponseWriter, r *http.Request) {

	j := (`{
  "skills" : [ {
    "current_skill_level" : 1,
    "skill_id" : 1,
    "skillpoints_in_skill" : 10000
  }, {
    "current_skill_level" : 1,
    "skill_id" : 2,
    "skillpoints_in_skill" : 10000
  } ],
  "total_sp" : 20000
}`)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		w.Write([]byte(j))
}


