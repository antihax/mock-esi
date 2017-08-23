package esilegacy

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var _ time.Time
var _ = mux.NewRouter

func GetFwLeaderboards(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		datasource string
		userAgent  string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "kills" : {
    "active_total" : [ {
      "amount" : 832273,
      "faction_id" : 500004
    }, {
      "amount" : 687915,
      "faction_id" : 500001
    } ],
    "last_week" : [ {
      "amount" : 730,
      "faction_id" : 500001
    }, {
      "amount" : 671,
      "faction_id" : 500004
    } ],
    "yesterday" : [ {
      "amount" : 100,
      "faction_id" : 500001
    }, {
      "amount" : 50,
      "faction_id" : 500004
    } ]
  },
  "victory_points" : {
    "active_total" : [ {
      "amount" : 53130500,
      "faction_id" : 500001
    }, {
      "amount" : 50964263,
      "faction_id" : 500004
    } ],
    "last_week" : [ {
      "amount" : 97360,
      "faction_id" : 500001
    }, {
      "amount" : 84980,
      "faction_id" : 500004
    } ],
    "yesterday" : [ {
      "amount" : 5000,
      "faction_id" : 500002
    }, {
      "amount" : 3500,
      "faction_id" : 500003
    } ]
  }
}`
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
	if r.Form.Get("userAgent") != "" {
		localV, err = processParameters(userAgent, r.Form.Get("user_agent"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		userAgent = localV.(string)
	}

	if r.Form.Get("page") != "" {
		var (
			localPage    int32
			localIntPage interface{}
		)
		localIntPage, err := processParameters(localPage, r.Form.Get("page"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		localPage = localIntPage.(int32)
		if localPage > 1 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}

func GetFwLeaderboardsCharacters(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		datasource string
		userAgent  string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "kills" : {
    "active_total" : [ {
      "amount" : 10000,
      "character_id" : 2112625428
    }, {
      "amount" : 8500,
      "character_id" : 95465499
    } ],
    "last_week" : [ {
      "amount" : 100,
      "character_id" : 2112625428
    }, {
      "amount" : 70,
      "character_id" : 95465499
    } ],
    "yesterday" : [ {
      "amount" : 34,
      "character_id" : 2112625428
    }, {
      "amount" : 20,
      "character_id" : 95465499
    } ]
  },
  "victory_points" : {
    "active_total" : [ {
      "amount" : 1239158,
      "character_id" : 2112625428
    }, {
      "amount" : 1139029,
      "character_id" : 95465499
    } ],
    "last_week" : [ {
      "amount" : 2660,
      "character_id" : 2112625428
    }, {
      "amount" : 2000,
      "character_id" : 95465499
    } ],
    "yesterday" : [ {
      "amount" : 620,
      "character_id" : 2112625428
    }, {
      "amount" : 550,
      "character_id" : 95465499
    } ]
  }
}`
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
	if r.Form.Get("userAgent") != "" {
		localV, err = processParameters(userAgent, r.Form.Get("user_agent"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		userAgent = localV.(string)
	}

	if r.Form.Get("page") != "" {
		var (
			localPage    int32
			localIntPage interface{}
		)
		localIntPage, err := processParameters(localPage, r.Form.Get("page"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		localPage = localIntPage.(int32)
		if localPage > 1 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}

func GetFwLeaderboardsCorporations(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		datasource string
		userAgent  string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `{
  "kills" : {
    "active_total" : [ {
      "amount" : 81692,
      "corporation_id" : 1000180
    }, {
      "amount" : 76793,
      "corporation_id" : 1000182
    } ],
    "last_week" : [ {
      "amount" : 290,
      "corporation_id" : 1000180
    }, {
      "amount" : 169,
      "corporation_id" : 1000182
    } ],
    "yesterday" : [ {
      "amount" : 51,
      "corporation_id" : 1000180
    }, {
      "amount" : 39,
      "corporation_id" : 1000182
    } ]
  },
  "victory_points" : {
    "active_total" : [ {
      "amount" : 18640927,
      "corporation_id" : 1000180
    }, {
      "amount" : 18078265,
      "corporation_id" : 1000181
    } ],
    "last_week" : [ {
      "amount" : 91980,
      "corporation_id" : 1000180
    }, {
      "amount" : 58920,
      "corporation_id" : 1000181
    } ],
    "yesterday" : [ {
      "amount" : 12600,
      "corporation_id" : 1000180
    }, {
      "amount" : 8240,
      "corporation_id" : 1000181
    } ]
  }
}`
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
	if r.Form.Get("userAgent") != "" {
		localV, err = processParameters(userAgent, r.Form.Get("user_agent"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		userAgent = localV.(string)
	}

	if r.Form.Get("page") != "" {
		var (
			localPage    int32
			localIntPage interface{}
		)
		localIntPage, err := processParameters(localPage, r.Form.Get("page"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		localPage = localIntPage.(int32)
		if localPage > 1 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}

func GetFwStats(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		datasource string
		userAgent  string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "faction_id" : 500001,
  "kills" : {
    "last_week" : 893,
    "total" : 684350,
    "yesterday" : 136
  },
  "pilots" : 28863,
  "systems_controlled" : 20,
  "victory_points" : {
    "last_week" : 102640,
    "total" : 52658260,
    "yesterday" : 15980
  }
} ]`
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
	if r.Form.Get("userAgent") != "" {
		localV, err = processParameters(userAgent, r.Form.Get("user_agent"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		userAgent = localV.(string)
	}

	if r.Form.Get("page") != "" {
		var (
			localPage    int32
			localIntPage interface{}
		)
		localIntPage, err := processParameters(localPage, r.Form.Get("page"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		localPage = localIntPage.(int32)
		if localPage > 1 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}

func GetFwSystems(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		datasource string
		userAgent  string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "contested" : true,
  "occupier_faction_id" : 500001,
  "owner_faction_id" : 500001,
  "solar_system_id" : 30002096,
  "victory_points" : 60,
  "victory_points_threshold" : 3000
} ]`
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
	if r.Form.Get("userAgent") != "" {
		localV, err = processParameters(userAgent, r.Form.Get("user_agent"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		userAgent = localV.(string)
	}

	if r.Form.Get("page") != "" {
		var (
			localPage    int32
			localIntPage interface{}
		)
		localIntPage, err := processParameters(localPage, r.Form.Get("page"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		localPage = localIntPage.(int32)
		if localPage > 1 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}

func GetFwWars(w http.ResponseWriter, r *http.Request) {

	var (
		localV     interface{}
		err        error
		datasource string
		userAgent  string
	)
	// shut up warnings
	localV = localV
	err = err

	j := `[ {
  "against_id" : 500002,
  "faction_id" : 500001
} ]`
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
	if r.Form.Get("userAgent") != "" {
		localV, err = processParameters(userAgent, r.Form.Get("user_agent"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		userAgent = localV.(string)
	}

	if r.Form.Get("page") != "" {
		var (
			localPage    int32
			localIntPage interface{}
		)
		localIntPage, err := processParameters(localPage, r.Form.Get("page"))
		if err != nil {
			errorOut(w, r, err)
			return
		}
		localPage = localIntPage.(int32)
		if localPage > 1 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("[]"))
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write([]byte(j))
}
