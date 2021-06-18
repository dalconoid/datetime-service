package server

import (
	"encoding/json"
	"github.com/dalconoid/datetime-service/model"
	"github.com/dalconoid/datetime-service/time_operations"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// handleAlive - handle for [/alive]
func handleAlive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

// handleGetTimeNow - handle for [/time/now]
func handleGetTimeNow() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := time.Now().Add(delta)
		timeF64 := time_operations.TimeToF64(t)
		response := &model.NowResponse{Time: timeF64}
		enc := json.NewEncoder(w)
		if err := enc.Encode(&response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// handleGetTimeString - handle for [time/string]
func handleGetTimeString() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		req := &model.TimeToStrRequest{}
		if err = json.Unmarshal(data, req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		t, err := time_operations.F64ToTime(req.Time)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		str := time_operations.TimeToString(t)
		response := &model.TimeToStrResponse{Str: str}
		enc := json.NewEncoder(w)
		if err = enc.Encode(&response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// handleAddTime - handle for [time/add]
func handleAddTime() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		req := &model.TimeAddRequest{}
		if err = json.Unmarshal(data, req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		t, err := time_operations.F64ToTime(req.Time)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		d, err := time_operations.F64ToDuration(req.Delta)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tPlusDelta := t.Add(d)
		tF64 := time_operations.TimeToF64(tPlusDelta)

		response := &model.TimeAddResponse{Time: tF64}
		enc := json.NewEncoder(w)
		if err = enc.Encode(&response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// handleCorrectServerTime - handle for [/time/correct]
func handleCorrectServerTime() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		defer r.Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		req := &model.TimeCorrectRequest{}
		if err = json.Unmarshal(data, req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		t, err := time_operations.F64ToTime(req.Time)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		now := time.Now()
		delta = t.Sub(now)
		//log.Printf("Correct time: %s\nNow time: %s\n", time_operations.TimeToString(t), time_operations.TimeToString(now))
		log.Printf("Delta changed to [%s]", delta.String())
		w.WriteHeader(http.StatusOK)
	}
}