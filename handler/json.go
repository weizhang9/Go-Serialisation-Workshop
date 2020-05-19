package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/353solutions/weather/weather"
	"github.com/gorilla/mux"
)

const (
	jsonCtype  = "application/json"
	protoCtype = "application/protobuf" // not standardised yet, comsumer will need to know to send this type
)

// curl localhost:8080/2020-03-01
func JSONGet(w http.ResponseWriter, r *http.Request) {
	// Step 1: De-serialise
	vars := mux.Vars(r)
	s := vars["date"] // "2020-03-01" "2020/03/01" RFC3339

	date, err := time.Parse("2006-01-02", s) // normalise date of different formats
	if err != nil {
		msg := fmt.Sprintf("bad date: %q", s)
		log.Printf(msg)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	// Step 2: Business Logic
	rec, err := weather.GetRecord(date)
	if err != nil {
		log.Printf(err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// handle protobuf request
	ctype := r.Header.Get("Accept")
	// case insensitive
	if strings.EqualFold(ctype, protoCtype) {
		protoSend(w, rec)
		return
	}

	// Step 3: Serialise
	data, err := json.Marshal(rec)
	if err != nil {
		log.Printf("marshal - %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("json data size: %d", len(data))

	w.Header().Set("Content-Type", jsonCtype)
	_, err = w.Write(data)
	if err != nil {
		log.Printf("can't write %x - %s", data, err)
	}
}

func JSONAdd(w http.ResponseWriter, r *http.Request) {
	// Step 1: De-serialise
	var rec weather.Record
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&rec); err != nil {
		log.Printf("Unmarshal: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Step 2: Business Logic
	log.Printf("Adding %#v", rec)
	n := weather.AddRecord(rec)

	resp := map[string]interface{}{
		"ok":          true,
		"num_records": n,
	}

	// Step 3: Serialise
	w.Header().Set("Content-Type", jsonCtype)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Printf("can't encode %#v - %s", resp, err)
	}
}
