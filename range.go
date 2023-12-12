package main

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"net/url"
	"strconv"
)

func parseLatlon(param url.Values, s string) (float32, bool) {
	if len(param[s]) != 1 {
		return 0, false
	}
	ll, e := strconv.ParseFloat(param[s][0], 32)
	if e != nil {
		return 0, false
	}
	if math.IsNaN(ll) || math.IsInf(ll, 0) || ll < (-370) || 370 < ll {
		return 0, false
	}
	return float32(ll), true
}

var rangeHandler = func(resp http.ResponseWriter, req *http.Request) {
	param := req.URL.Query()
	lat1, ok1 := parseLatlon(param, "latitude1")
	lat2, ok2 := parseLatlon(param, "latitude2")
	lng1, ok3 := parseLatlon(param, "longitude1")
	lng2, ok4 := parseLatlon(param, "longitude2")
	if !(ok1 && ok2 && ok3 && ok4) {
		resp.WriteHeader(http.StatusForbidden)
	}
	if lat1 > 91 || lat1 < (-91) || lat2 > 91 || lat2 < (-91) ||
		lng1 > 181 || lng2 > 181 || lng1 < (-181) || lng2 < (-181) {
		resp.WriteHeader(http.StatusOK)
		if _, err := resp.Write([]byte("hey yo motherfucker")); err != nil {
			log.Fatalf("taunt failed %+v", err)
		}
		return
	}

	notes, err := rangeFetch(lat1, lng1, lat2, lng2)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("range query db error %+v \n", err)
	}
	notesString, err := json.Marshal(notes)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("range query marshal error %+v \n", err)
	}
	if _, err = resp.Write(notesString); err != nil {
		log.Fatalf("range query write response error %+v \n", err)
	}
}
